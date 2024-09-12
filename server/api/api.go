package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type Quotation struct {
	Code       string  `json:"code"`
	CodeIn     string  `json:"codein"`
	Name       string  `json:"name"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	VarBid     float64 `json:"varBid"`
	PctChange  float64 `json:"pctChange"`
	Bid        float64 `json:"bid"`
	Ask        float64 `json:"ask"`
	Timestamp  string  `json:"timestamp"`
	CreateDate string  `json:"create_date"`
}

func GetQuotation(ctx context.Context) (*Quotation, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/all/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("API request failed")
	}

	var quotation Quotation
	err = json.NewDecoder(resp.Body).Decode(&quotation)
	if err != nil {
		return nil, err
	}

	return &quotation, nil
}
