package iplicit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-iplicit/utils"
)

func (c *Client) NewPostSaleInvoiceRequest() PostSaleInvoiceRequest {
	return PostSaleInvoiceRequest{
		client:      c,
		queryParams: c.NewPostSaleInvoiceQueryParams(),
		pathParams:  c.NewPostSaleInvoicePathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostSaleInvoiceRequestBody(),
	}
}

type PostSaleInvoiceRequest struct {
	client      *Client
	queryParams *PostSaleInvoiceQueryParams
	pathParams  *PostSaleInvoicePathParams
	method      string
	headers     http.Header
	requestBody PostSaleInvoiceRequestBody
}

func (c *Client) NewPostSaleInvoiceQueryParams() *PostSaleInvoiceQueryParams {
	return &PostSaleInvoiceQueryParams{}
}

type PostSaleInvoiceQueryParams struct {
}

func (p PostSaleInvoiceQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostSaleInvoiceRequest) QueryParams() *PostSaleInvoiceQueryParams {
	return r.queryParams
}

func (c *Client) NewPostSaleInvoicePathParams() *PostSaleInvoicePathParams {
	return &PostSaleInvoicePathParams{}
}

type PostSaleInvoicePathParams struct {
}

func (p *PostSaleInvoicePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostSaleInvoiceRequest) PathParams() *PostSaleInvoicePathParams {
	return r.pathParams
}

func (r *PostSaleInvoiceRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostSaleInvoiceRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostSaleInvoiceRequest) Method() string {
	return r.method
}

func (s *Client) NewPostSaleInvoiceRequestBody() PostSaleInvoiceRequestBody {
	return PostSaleInvoiceRequestBody{}
}

type PostSaleInvoiceRequestBody SaleInvoicePost

func (r *PostSaleInvoiceRequest) RequestBody() *PostSaleInvoiceRequestBody {
	return &r.requestBody
}

func (r *PostSaleInvoiceRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostSaleInvoiceRequest) SetRequestBody(body PostSaleInvoiceRequestBody) {
	r.requestBody = body
}

func (r *PostSaleInvoiceRequest) NewResponseBody() *PostSaleInvoiceResponseBody {
	return &PostSaleInvoiceResponseBody{}
}

type PostSaleInvoiceResponseBody struct {
	SaleInvoiceID string `json:"sale_invoice_id,omitempty"`
}

func (r *PostSaleInvoiceResponseBody) UnmarshalJSON(data []byte) error {
	var SaleInvoiceID string
	if err := json.Unmarshal(data, &SaleInvoiceID); err == nil {
		r.SaleInvoiceID = SaleInvoiceID
		return nil
	}

	return fmt.Errorf("Unable to unmarshal response")
}

func (r *PostSaleInvoiceRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("SaleInvoice", r.PathParams())
	return &u
}

func (r *PostSaleInvoiceRequest) Do(ctx context.Context) (PostSaleInvoiceResponseBody, error) {
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
