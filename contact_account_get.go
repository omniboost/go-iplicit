package iplicit

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/omniboost/go-iplicit/utils"
)

func (c *Client) NewGetContactAccountRequest() GetContactAccountRequest {
	return GetContactAccountRequest{
		client:      c,
		queryParams: c.NewGetContactAccountQueryParams(),
		pathParams:  c.NewGetContactAccountPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetContactAccountRequestBody(),
	}
}

type GetContactAccountRequest struct {
	client      *Client
	queryParams *GetContactAccountQueryParams
	pathParams  *GetContactAccountPathParams
	method      string
	headers     http.Header
	requestBody GetContactAccountRequestBody
}

func (c *Client) NewGetContactAccountQueryParams() *GetContactAccountQueryParams {
	return &GetContactAccountQueryParams{}
}

type GetContactAccountQueryParams struct {
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

func (p GetContactAccountQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetContactAccountRequest) QueryParams() *GetContactAccountQueryParams {
	return r.queryParams
}

func (c *Client) NewGetContactAccountPathParams() *GetContactAccountPathParams {
	return &GetContactAccountPathParams{}
}

type GetContactAccountPathParams struct {
	ContactID string `schema:"contact_id,omitempty"`
}

func (p *GetContactAccountPathParams) Params() map[string]string {
	return map[string]string{
		"contact_id": p.ContactID,
	}
}

func (r *GetContactAccountRequest) PathParams() *GetContactAccountPathParams {
	return r.pathParams
}

func (r *GetContactAccountRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetContactAccountRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetContactAccountRequest) Method() string {
	return r.method
}

func (s *Client) NewGetContactAccountRequestBody() GetContactAccountRequestBody {
	return GetContactAccountRequestBody{}
}

type GetContactAccountRequestBody struct {
}

func (r *GetContactAccountRequest) RequestBody() *GetContactAccountRequestBody {
	return &r.requestBody
}

func (r *GetContactAccountRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetContactAccountRequest) SetRequestBody(body GetContactAccountRequestBody) {
	r.requestBody = body
}

func (r *GetContactAccountRequest) NewResponseBody() *GetContactAccountResponseBody {
	return &GetContactAccountResponseBody{}
}

type GetContactAccountResponseBody []ContactAccountGet

func (v *GetContactAccountResponseBody) UnmarshalJSON(b []byte) (err error) {
	vc, vcs := ContactAccountGet{}, []ContactAccountGet{}
	if err = json.Unmarshal(b, &vc); err == nil {
		*v = make([]ContactAccountGet, 1)
		[]ContactAccountGet(*v)[0] = ContactAccountGet(vc)
		return
	}
	if err = json.Unmarshal(b, &vcs); err == nil {
		*v = GetContactAccountResponseBody(vcs)
		return
	}
	return
}

func (r *GetContactAccountRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("ContactAccount/{{.contact_id}}", r.PathParams())
	return &u
}

func (r *GetContactAccountRequest) Do(ctx context.Context) (GetContactAccountResponseBody, error) {
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
