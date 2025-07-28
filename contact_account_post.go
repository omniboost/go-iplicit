package iplicit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-iplicit/utils"
)

func (c *Client) NewPostContactAccountRequest() PostContactAccountRequest {
	return PostContactAccountRequest{
		client:      c,
		queryParams: c.NewPostContactAccountQueryParams(),
		pathParams:  c.NewPostContactAccountPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostContactAccountRequestBody(),
	}
}

type PostContactAccountRequest struct {
	client      *Client
	queryParams *PostContactAccountQueryParams
	pathParams  *PostContactAccountPathParams
	method      string
	headers     http.Header
	requestBody PostContactAccountRequestBody
}

func (c *Client) NewPostContactAccountQueryParams() *PostContactAccountQueryParams {
	return &PostContactAccountQueryParams{}
}

type PostContactAccountQueryParams struct {
}

func (p PostContactAccountQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostContactAccountRequest) QueryParams() *PostContactAccountQueryParams {
	return r.queryParams
}

func (c *Client) NewPostContactAccountPathParams() *PostContactAccountPathParams {
	return &PostContactAccountPathParams{}
}

type PostContactAccountPathParams struct {
}

func (p *PostContactAccountPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostContactAccountRequest) PathParams() *PostContactAccountPathParams {
	return r.pathParams
}

func (r *PostContactAccountRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostContactAccountRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostContactAccountRequest) Method() string {
	return r.method
}

func (s *Client) NewPostContactAccountRequestBody() PostContactAccountRequestBody {
	return PostContactAccountRequestBody{}
}

type PostContactAccountRequestBody ContactAccountPost

func (r *PostContactAccountRequest) RequestBody() *PostContactAccountRequestBody {
	return &r.requestBody
}

func (r *PostContactAccountRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostContactAccountRequest) SetRequestBody(body PostContactAccountRequestBody) {
	r.requestBody = body
}

func (r *PostContactAccountRequest) NewResponseBody() *PostContactAccountResponseBody {
	return &PostContactAccountResponseBody{}
}

type PostContactAccountResponseBody struct {
	ContactID string `json:"contact_id,omitempty"`
}

func (r *PostContactAccountResponseBody) UnmarshalJSON(data []byte) error {
	var ContactID string
	if err := json.Unmarshal(data, &ContactID); err == nil {
		r.ContactID = ContactID
		return nil
	}

	return fmt.Errorf("Unable to unmarshal response")
}

func (r *PostContactAccountRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("ContactAccount", r.PathParams())
	return &u
}

func (r *PostContactAccountRequest) Do(ctx context.Context) (PostContactAccountResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, true)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody, 3)
	return *responseBody, err
}
