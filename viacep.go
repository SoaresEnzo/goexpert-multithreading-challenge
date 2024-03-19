package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type ViaCepService struct {
}

func NewViaCepService() *ViaCepService {
	return &ViaCepService{}
}

func (v *ViaCepService) SearchZipcode(zipcode string, ctx context.Context) (*ViaCepResponse, error) {
	viaCepUrl := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	req, err := http.NewRequestWithContext(ctx, "GET", viaCepUrl, nil)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var viaCepResponse ViaCepResponse
	err = json.NewDecoder(res.Body).Decode(&viaCepResponse)

	if err != nil {
		return nil, err
	}

	return &viaCepResponse, nil
}
