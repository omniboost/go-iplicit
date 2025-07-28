package iplicit

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetContactAccountRequest(t *testing.T) {
	req := client.NewGetContactAccountRequest()
	req.PathParams().ContactID = "CODE"
	// req.QueryParams().Include = []Include{IncludeAllBankDetails}
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
