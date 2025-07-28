package iplicit

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-iplicit/utils"
)

func (c *Client) NewPostTokenRequest() PostTokenRequest {
	return PostTokenRequest{
		client:      c,
		queryParams: c.NewPostTokenQueryParams(),
		pathParams:  c.NewPostTokenPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostTokenRequestBody(),
	}
}

type PostTokenRequest struct {
	client      *Client
	queryParams *PostTokenQueryParams
	pathParams  *PostTokenPathParams
	method      string
	headers     http.Header
	requestBody PostTokenRequestBody
}

func (c *Client) NewPostTokenQueryParams() *PostTokenQueryParams {
	return &PostTokenQueryParams{}
}

type PostTokenQueryParams struct {
}

func (p PostTokenQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostTokenRequest) QueryParams() *PostTokenQueryParams {
	return r.queryParams
}

func (c *Client) NewPostTokenPathParams() *PostTokenPathParams {
	return &PostTokenPathParams{}
}

type PostTokenPathParams struct {
}

func (p *PostTokenPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostTokenRequest) PathParams() *PostTokenPathParams {
	return r.pathParams
}

func (r *PostTokenRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostTokenRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostTokenRequest) Method() string {
	return r.method
}

func (s *Client) NewPostTokenRequestBody() PostTokenRequestBody {
	return PostTokenRequestBody{}
}

type PostTokenRequestBody struct {
	Username   string `json:"username"`
	UserApiKey string `json:"userApiKey"`
}

func (r *PostTokenRequest) RequestBody() *PostTokenRequestBody {
	return &r.requestBody
}

func (r *PostTokenRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostTokenRequest) SetRequestBody(body PostTokenRequestBody) {
	r.requestBody = body
}

func (r *PostTokenRequest) NewResponseBody() *PostTokenResponseBody {
	return &PostTokenResponseBody{}
}

type PostTokenResponseBody TokenResponsePost

func (r *PostTokenRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("session/create/api", r.PathParams())
	return &u
}

func (r *PostTokenRequest) Do() (PostTokenResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(context.Background(), r)
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
