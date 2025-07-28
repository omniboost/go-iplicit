package iplicit

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestPostSaleInvoiceRequest(t *testing.T) {
	req := client.NewPostSaleInvoiceRequest()
	req.requestBody = PostSaleInvoiceRequestBody{
		DocNo:            "",
		DocSerieID:       "",
		DocTypeID:        "",
		LegalEntityID:    "",
		ContactAccountID: "",
		Description:      "SNT Sales Import - Guest Name",
		DocDate:          DateTime{time.Date(2025, 03, 12, 0, 0, 0, 0, time.UTC)},
		Currency:         "",
		CurrencyRate:     1,
		IsNetEntry:       true,
		Details: []DetailsSaleRequest{
			{
				ProductID: "",
				DetailsPost: DetailsPostSaleRequest{
					CostCentre: "",
					Department: "",
					RateCode:   "",
				},
				Description:          "",
				IsNetEntry:           true,
				TaxBandID:            "S",
				NetCurrencyUnitPrice: 110,
				Quantity:             1,
				NetCurrencyAmount:    110,
			},
			{
				ProductID: "",
				DetailsPost: DetailsPostSaleRequest{
					CostCentre: "",
					Department: "",
					RateCode:   "",
				},
				Description:          "",
				IsNetEntry:           true,
				TaxBandID:            "S",
				NetCurrencyUnitPrice: 13.33,
				Quantity:             1,
				NetCurrencyAmount:    13.33,
			},
			{
				ProductID: "",
				DetailsPost: DetailsPostSaleRequest{
					CostCentre: "",
					Department: "",
					RateCode:   "",
				},
				Description:          "",
				IsNetEntry:           true,
				TaxBandID:            "S",
				NetCurrencyUnitPrice: 60,
				Quantity:             1,
				NetCurrencyAmount:    60,
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
