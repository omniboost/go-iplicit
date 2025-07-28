package iplicit

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/hashicorp/go-multierror"
)

var DefaultBaseURL = url.URL{
	Scheme: "https",
	Host:   "api.iplicit.com",
	Path:   "api",
}

const (
	mediaType      = "application/problem+json"
	libraryVersion = "0.0.1"
	userAgent      = "go-iplicit/" + libraryVersion
)

type (
	Client struct {
		HttpClient *http.Client

		Debug      bool
		BaseURL    *url.URL
		Username   string
		UserApiKey string

		UserAgent             string
		DisallowUnknownFields bool

		SessionToken string
		TokenExpiry  time.Time

		Domain string
	}

	ValidationError struct {
		Type    string              `json:"type"`
		Title   string              `json:"title"`
		Status  int                 `json:"status"`
		Message string              `json:"message,omitempty"`
		TraceId string              `json:"traceId,omitempty"`
		Errors  map[string][]string `json:"errors,omitempty"`
	}

	ErrorResponse struct {
		// HTTP response that caused this error
		Response *http.Response `json:"-"`

		ValidationError *ValidationError `json:"-"`
	}
)

func (r *ErrorResponse) Error() string {
	var result *multierror.Error

	if r.Response.StatusCode < 200 || r.Response.StatusCode > 299 {
		result = multierror.Append(result, fmt.Errorf("HTTP %d: %s", r.Response.StatusCode, r.Response.Status))
	}

	if r.ValidationError != nil {
		if r.ValidationError.Title != "" {
			result = multierror.Append(result, fmt.Errorf("API Error: %s", r.ValidationError.Title))
		}

		if r.ValidationError.Message != "" {
			result = multierror.Append(result, fmt.Errorf("Message: %s", r.ValidationError.Message))
		}

		if len(r.ValidationError.Errors) > 0 {
			for field, errors := range r.ValidationError.Errors {
				for _, err := range errors {
					result = multierror.Append(result, fmt.Errorf("Validation error [%s]: %s", field, err))
				}
			}
		}

		if r.ValidationError.TraceId != "" {
			result = multierror.Append(result, fmt.Errorf("Trace ID: %s", r.ValidationError.TraceId))
		}
	}

	if result == nil {
		return ""
	}

	return result.Error()
}

func (r *ErrorResponse) UnmarshalJSON(data []byte) error {
	var validationError ValidationError
	if err := json.Unmarshal(data, &validationError); err == nil {
		if validationError.Type != "" || validationError.Title != "" || len(validationError.Errors) > 0 {
			r.ValidationError = &validationError
			return nil
		}
	}

	type Alias ErrorResponse
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	return nil
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		HttpClient: httpClient,
		BaseURL:    &DefaultBaseURL,
		UserAgent:  userAgent,
	}
}

func (c *Client) isTokenValid() bool {
	valid := c.SessionToken != "" && time.Now().Before(c.TokenExpiry)
	return valid
}

func (c *Client) GetSessionToken() (string, error) {
	if c.isTokenValid() {
		return c.SessionToken, nil
	}

	return c.refreshToken()
}

func (c *Client) refreshToken() (string, error) {
	req := c.NewPostTokenRequest()

	requestBody := PostTokenRequestBody{
		Username:   c.Username,
		UserApiKey: c.UserApiKey,
	}
	req.SetRequestBody(requestBody)

	resp, err := req.Do()
	if err != nil {
		return "", err
	}

	c.SessionToken = resp.SessionToken
	c.TokenExpiry = time.Now().Add(30 * time.Minute)

	return c.SessionToken, nil
}

func (c *Client) GetEndpointURL(p string, pathParams PathParams) url.URL {
	if c.BaseURL == nil {
		log.Fatal("BaseURL cannot be nil")
	}

	clientURL := *c.BaseURL

	parsed, err := url.Parse(p)
	if err != nil {
		log.Fatal(err)
	}
	q := clientURL.Query()
	for k, vv := range parsed.Query() {
		for _, v := range vv {
			q.Add(k, v)
		}
	}
	clientURL.RawQuery = q.Encode()

	clientURL.Path = path.Join(clientURL.Path, parsed.Path)

	tmpl, err := template.New("path").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	if err = tmpl.Execute(buf, params); err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	tokenReq := c.NewPostTokenRequest()

	// convert body struct to json
	var body io.Reader
	if req.RequestBodyInterface() != nil {
		if r, ok := req.RequestBodyInterface().(io.Reader); ok {
			body = r
		} else if bb, ok := req.RequestBodyInterface().([]byte); ok {
			body = bytes.NewReader(bb)
		} else {
			buf := new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(req.RequestBodyInterface())
			if err != nil {
				return nil, err
			}
			body = buf
		}
	}

	// create new http request
	r, err := http.NewRequestWithContext(ctx, req.Method(), req.URL().String(), body)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	if req.URL().String() != tokenReq.URL().String() {
		token, err := c.GetSessionToken()
		if err != nil {
			return nil, err
		}
		r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	// set other headers
	r.Header.Add("Content-Type", mediaType+"; charset=utf-8")
	r.Header.Add("Accept", mediaType)
	r.Header.Add("User-Agent", c.UserAgent)
	r.Header.Add("Domain", c.Domain)

	return r, nil
}

func (c *Client) Do(req *http.Request, responseBody interface{}, maxTries int) (*http.Response, error) {
	if c.Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}
	httpResp, err := c.HttpClient.Do(req)
	if err != nil {
		if c.Debug {
			log.Printf("Request failed: %s", err.Error())
		}
		if maxTries > 0 {
			time.Sleep(100 * time.Millisecond)
			return c.Do(req, responseBody, maxTries-1)
		}
		return nil, err
	}
	defer func() {
		_ = httpResp.Body.Close()
	}()

	if c.Debug {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	// try to decode body into interface parameter
	if responseBody == nil {
		return httpResp, nil
	}

	errorResponse := &ErrorResponse{Response: httpResp}
	err = c.Unmarshal(httpResp.Body, responseBody, errorResponse)
	if err != nil {
		return httpResp, err
	}

	if errorResponse.Error() != "" {
		return httpResp, errorResponse
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	var errs []error
	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)
		if c.DisallowUnknownFields {
			dec.DisallowUnknownFields()
		}

		err := dec.Decode(v)
		if err != nil && !errors.Is(err, io.EOF) {
			errs = append(errs, err)
		}

	}

	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = e.Error()
		}
		return errors.New(strings.Join(msgs, ", "))
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	if c := r.StatusCode; (c >= 200 && c <= 299) || c == 400 {
		return nil
	}

	err := checkContentType(r)
	if err != nil {
		return errors.New(r.Status)
	}

	// read data and copy it back
	data, err := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return nil
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}
