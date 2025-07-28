package iplicit

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/omniboost/go-iplicit/utils"
)

func (c *Client) NewGetProductRequest() GetProductRequest {
	return GetProductRequest{
		client:      c,
		queryParams: c.NewGetProductQueryParams(),
		pathParams:  c.NewGetProductPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetProductRequestBody(),
	}
}

type GetProductRequest struct {
	client      *Client
	queryParams *GetProductQueryParams
	pathParams  *GetProductPathParams
	method      string
	headers     http.Header
	requestBody GetProductRequestBody
}

func (c *Client) NewGetProductQueryParams() *GetProductQueryParams {
	return &GetProductQueryParams{}
}

type GetProductQueryParams struct {
	DescriptionContains string   `schema:"descriptionContains,omitempty"`
	ProductGroupID      string   `schema:"productGroupId,omitempty"`
	LastModifiedFrom    DateTime `schema:"lastModifiedFrom,omitempty"`
	LastModifiedTo      DateTime `schema:"lastModifiedTo,omitempty"`
	IsPurchase          *bool    `schema:"isPurchase,omitempty"`
	IsSale              *bool    `schema:"isSale,omitempty"`
	// Include can only be used with a specified Product ID or Code.
	Include []IncludeForProduct `schema:"include,omitempty"`
	Take    int                 `schema:"take,omitempty"`
	Skip    int                 `schema:"skip,omitempty"`
}

func (p GetProductQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetProductRequest) QueryParams() *GetProductQueryParams {
	return r.queryParams
}

func (c *Client) NewGetProductPathParams() *GetProductPathParams {
	return &GetProductPathParams{}
}

type GetProductPathParams struct {
	ProductIdOrCode string `schema:"product_id_or_code,omitempty"`
}

func (p *GetProductPathParams) Params() map[string]string {
	return map[string]string{
		"product_id_or_code": p.ProductIdOrCode,
	}
}

func (r *GetProductRequest) PathParams() *GetProductPathParams {
	return r.pathParams
}

func (r *GetProductRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetProductRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetProductRequest) Method() string {
	return r.method
}

func (s *Client) NewGetProductRequestBody() GetProductRequestBody {
	return GetProductRequestBody{}
}

type GetProductRequestBody struct {
}

func (r *GetProductRequest) RequestBody() *GetProductRequestBody {
	return &r.requestBody
}

func (r *GetProductRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetProductRequest) SetRequestBody(body GetProductRequestBody) {
	r.requestBody = body
}

func (r *GetProductRequest) NewResponseBody() *GetProductResponseBody {
	return &GetProductResponseBody{}
}

type GetProductResponseBody []ProductResponseGet

func (v *GetProductResponseBody) UnmarshalJSON(b []byte) (err error) {
	vc, vcs := ProductResponseGet{}, []ProductResponseGet{}
	if err = json.Unmarshal(b, &vc); err == nil {
		*v = make([]ProductResponseGet, 1)
		[]ProductResponseGet(*v)[0] = ProductResponseGet(vc)
		return
	}
	if err = json.Unmarshal(b, &vcs); err == nil {
		*v = GetProductResponseBody(vcs)
		return
	}
	return
}

func (r *GetProductRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("Product/{{.product_id_or_code}}", r.PathParams())
	return &u
}

func (r *GetProductRequest) Do(ctx context.Context) (GetProductResponseBody, error) {
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
