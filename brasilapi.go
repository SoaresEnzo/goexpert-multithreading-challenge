package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type BrasilApiResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type BrasilApiService struct {
}

func NewBrasilApiService() *BrasilApiService {
	return &BrasilApiService{}
}

func (v *BrasilApiService) SearchZipcode(zipcode string, ctx context.Context) (*BrasilApiResponse, error) {
	brasilApiUrl := "https://brasilapi.com.br/api/cep/v1/" + zipcode

	req, err := http.NewRequestWithContext(ctx, "GET", brasilApiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var brasilApiResponse BrasilApiResponse
	err = json.NewDecoder(res.Body).Decode(&brasilApiResponse)
	if err != nil {
		return nil, err
	}
	return &brasilApiResponse, nil
}
