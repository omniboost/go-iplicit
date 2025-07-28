package iplicit

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestPostReceiptRequest(t *testing.T) {
	req := client.NewPostReceiptRequest()
	req.requestBody = PostReceiptRequestBody{
		DocNo:            "",
		DocSerieID:       "",
		DocTypeID:        "",
		LegalEntityID:    "",
		ContactAccountID: "",
		DocDate:          DateTime{time.Date(2025, 03, 12, 0, 0, 0, 0, time.UTC)},
		Description:      "SNT AR Payment Import - Guest Name",
		TaxDate:          DateTime{time.Date(2025, 03, 12, 0, 0, 0, 0, time.UTC)},
		PaymentMethodID:  "",
		Post: PostReceiptReq{
			CostCentre: "",
		},
		Currency:           "",
		CurrencyRate:       1,
		IsNetEntry:         true,
		BankAccountID:      "",
		BankCurrency:       "",
		BankCurrencyAmount: 220,
		BankCurrencyRate:   1,
		CurrencyAmount:     220,
	}

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
