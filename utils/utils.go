package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	null "gopkg.in/guregu/null.v3"

	"github.com/gorilla/schema"
)

var (
	// register custom encoders
	EncodeSchemaMarshaler = func(v reflect.Value) string {
		if marshaler, ok := v.Interface().(SchemaMarshaler); ok {
			return marshaler.MarshalSchema()
		}

		if stringer, ok := v.Interface().(fmt.Stringer); ok {
			return stringer.String()
		}

		return ""
	}
)

type SchemaMarshaler interface {
	MarshalSchema() string
}

type ToURLValues interface {
	ToURLValues() (url.Values, error)
}

func AddQueryParamsToRequest(requestParams interface{}, req *http.Request, skipEmpty bool) error {
	var err error
	params := url.Values{}

	if to, ok := requestParams.(ToURLValues); ok {
		params, err = to.ToURLValues()
		if err != nil {
			return err
		}
	} else {
		encoder := NewSchemaEncoder()
		err := encoder.Encode(requestParams, params)
		if err != nil {
			return err
		}
	}

	return AddURLValuesToRequest(params, req, skipEmpty)
}

func AddURLValuesToRequest(params url.Values, req *http.Request, skipEmpty bool) error {
	query := req.URL.Query()
	for k, vals := range params {
		for _, v := range vals {
			if skipEmpty && v == "" {
				continue
			}

			if skipEmpty && v == "0" {
				continue
			}

			query.Add(k, v)
		}
	}

	req.URL.RawQuery = query.Encode()

	// force $ in query parameters
	req.URL.RawQuery = strings.Replace(req.URL.RawQuery, "%24", "$", -1)
	// req.URL.RawQuery = strings.Replace(req.URL.RawQuery, "+", "%20", -1)
	return nil
}

func NewSchemaEncoder() *schema.Encoder {
	encoder := schema.NewEncoder()

	encodeNullFloat := func(v reflect.Value) string {
		nullFloat, _ := v.Interface().(null.Float)
		if nullFloat.IsZero() {
			return ""
		}
		return strconv.FormatFloat(nullFloat.Float64, 'f', 6, 64)
	}

	encodeNullBool := func(v reflect.Value) string {
		nullBool, _ := v.Interface().(null.Bool)
		if nullBool.IsZero() {
			return ""
		}
		return strconv.FormatBool(nullBool.Bool)
	}

	encoder.RegisterEncoder(null.Float{}, encodeNullFloat)
	encoder.RegisterEncoder(null.Bool{}, encodeNullBool)
	return encoder
}

func PtrTo[T any](v T) *T {
	return &v
}
