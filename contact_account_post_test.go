package iplicit

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestPostContactAccountRequest(t *testing.T) {
	req := client.NewPostContactAccountRequest()
	req.requestBody = PostContactAccountRequestBody{
		Description: "Omniboost General Revenue Bucket 2",
		Code:        "OMNI-REVBUCKET-2",
		CountryCode: "GB",
		Customer: CustomerPost{
			ContactGroupCustomerID: "",
			PayTermID:              "",
			IsActive:               true,
		},
		Contact: ContactPost{
			CompanyName: "Omniboost General Revenue Bucket",
			FirstName:   "First name",
			MiddleName:  "Middle name",
			LastName:    "Hart",
			Emails: []EmailsPost{
				{
					Email: "",
				},
			},
			Phones: []PhonesPost{
				{
					Phone: "",
				},
			},
			Addresses: []AddressesPost{
				{
					Address:     "123 Example Street",
					PostCode:    "PostalCode",
					City:        "City",
					County:      "Country",
					CountryCode: "CountryCode",
				},
			},
		},
	}

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
