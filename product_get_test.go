package iplicit

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetProductRequest(t *testing.T) {
	req := client.NewGetProductRequest()
	// req.PathParams().ProductIdOrCode = "30130"
	// req.QueryParams().Take = 1
	// req.QueryParams().Include = []IncludeForProduct{IncludeExt}
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
