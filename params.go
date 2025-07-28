package iplicit

import (
	"net/url"
)

type Request interface {
	Method() string
	//QueryParams() QueryParams // In this setup the implementations return their specific implementation of the QueryParams
	PathParamsInterface() PathParams
	RequestBodyInterface() interface{}
	URL() *url.URL
}

type QueryParams interface {
	ToURLValues() (url.Values, error)
}

type PathParams interface {
	Params() map[string]string
}
