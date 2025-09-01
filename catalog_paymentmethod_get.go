package iplicit

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-iplicit/utils"
)

func (c *Client) NewGetCatalogPaymentMethodRequest() GetCatalogPaymentMethodRequest {
	return GetCatalogPaymentMethodRequest{
		client:      c,
		queryParams: c.NewGetCatalogPaymentMethodQueryParams(),
		pathParams:  c.NewGetCatalogPaymentMethodPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetCatalogPaymentMethodRequestBody(),
	}
}

type GetCatalogPaymentMethodRequest struct {
	client      *Client
	queryParams *GetCatalogPaymentMethodQueryParams
	pathParams  *GetCatalogPaymentMethodPathParams
	method      string
	headers     http.Header
	requestBody GetCatalogPaymentMethodRequestBody
}

func (c *Client) NewGetCatalogPaymentMethodQueryParams() *GetCatalogPaymentMethodQueryParams {
	return &GetCatalogPaymentMethodQueryParams{}
}

type GetCatalogPaymentMethodQueryParams struct {
	Code                string    `schema:"code,omitempty"`
	IntRef              string    `schema:"intRef,omitempty"`
	DescriptionContains string    `schema:"descriptionContains,omitempty"`
	LastModifiedFrom    DateTime  `schema:"lastModifiedFrom,omitempty"`
	LastModifiedTo      DateTime  `schema:"lastModifiedTo,omitempty"`
	IsCustomer          *bool     `schema:"isCustomer,omitempty"`
	IsSupplier          *bool     `schema:"isSupplier,omitempty"`
	IsResource          *bool     `schema:"isResource,omitempty"`
	Include             []Include `schema:"include,omitempty"`
	Take                int       `schema:"take,omitempty"`
	Skip                int       `schema:"skip,omitempty"`
}

func (p GetCatalogPaymentMethodQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetCatalogPaymentMethodRequest) QueryParams() *GetCatalogPaymentMethodQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCatalogPaymentMethodPathParams() *GetCatalogPaymentMethodPathParams {
	return &GetCatalogPaymentMethodPathParams{}
}

type GetCatalogPaymentMethodPathParams struct{}

func (p *GetCatalogPaymentMethodPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCatalogPaymentMethodRequest) PathParams() *GetCatalogPaymentMethodPathParams {
	return r.pathParams
}

func (r *GetCatalogPaymentMethodRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetCatalogPaymentMethodRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCatalogPaymentMethodRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCatalogPaymentMethodRequestBody() GetCatalogPaymentMethodRequestBody {
	return GetCatalogPaymentMethodRequestBody{}
}

type GetCatalogPaymentMethodRequestBody struct {
}

func (r *GetCatalogPaymentMethodRequest) RequestBody() *GetCatalogPaymentMethodRequestBody {
	return &r.requestBody
}

func (r *GetCatalogPaymentMethodRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetCatalogPaymentMethodRequest) SetRequestBody(body GetCatalogPaymentMethodRequestBody) {
	r.requestBody = body
}

func (r *GetCatalogPaymentMethodRequest) NewResponseBody() *GetCatalogPaymentMethodResponseBody {
	return &GetCatalogPaymentMethodResponseBody{}
}

type GetCatalogPaymentMethodResponseBody PaymentMethods

func (r *GetCatalogPaymentMethodRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("Catalog/PaymentMethod", r.PathParams())
	return &u
}

func (r *GetCatalogPaymentMethodRequest) Do(ctx context.Context) (GetCatalogPaymentMethodResponseBody, error) {
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
