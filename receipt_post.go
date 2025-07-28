package iplicit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-iplicit/utils"
)

func (c *Client) NewPostReceiptRequest() PostReceiptRequest {
	return PostReceiptRequest{
		client:      c,
		queryParams: c.NewPostReceiptQueryParams(),
		pathParams:  c.NewPostReceiptPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostReceiptRequestBody(),
	}
}

type PostReceiptRequest struct {
	client      *Client
	queryParams *PostReceiptQueryParams
	pathParams  *PostReceiptPathParams
	method      string
	headers     http.Header
	requestBody PostReceiptRequestBody
}

func (c *Client) NewPostReceiptQueryParams() *PostReceiptQueryParams {
	return &PostReceiptQueryParams{}
}

type PostReceiptQueryParams struct {
}

func (p PostReceiptQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostReceiptRequest) QueryParams() *PostReceiptQueryParams {
	return r.queryParams
}

func (c *Client) NewPostReceiptPathParams() *PostReceiptPathParams {
	return &PostReceiptPathParams{}
}

type PostReceiptPathParams struct {
}

func (p *PostReceiptPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostReceiptRequest) PathParams() *PostReceiptPathParams {
	return r.pathParams
}

func (r *PostReceiptRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostReceiptRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostReceiptRequest) Method() string {
	return r.method
}

func (s *Client) NewPostReceiptRequestBody() PostReceiptRequestBody {
	return PostReceiptRequestBody{}
}

type PostReceiptRequestBody ReceiptPost

func (r *PostReceiptRequest) RequestBody() *PostReceiptRequestBody {
	return &r.requestBody
}

func (r *PostReceiptRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostReceiptRequest) SetRequestBody(body PostReceiptRequestBody) {
	r.requestBody = body
}

func (r *PostReceiptRequest) NewResponseBody() *PostReceiptResponseBody {
	return &PostReceiptResponseBody{}
}

type PostReceiptResponseBody struct {
	ReceiptID string `json:"receipt_id,omitempty"`
}

func (r *PostReceiptResponseBody) UnmarshalJSON(data []byte) error {
	var ReceiptID string
	if err := json.Unmarshal(data, &ReceiptID); err == nil {
		r.ReceiptID = ReceiptID
		return nil
	}

	return fmt.Errorf("Unable to unmarshal response")
}

func (r *PostReceiptRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("Receipt", r.PathParams())
	return &u
}

func (r *PostReceiptRequest) Do(ctx context.Context) (PostReceiptResponseBody, error) {
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
