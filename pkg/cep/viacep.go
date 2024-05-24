package cep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type viaCEPResponse struct {
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
	Erro        string `json:"erro"`
}

const cep_url = "http://viacep.com.br/ws/%s/json/"

type ViaCep struct{}

func InstanceViaCep() *ViaCep {
	return &ViaCep{}
}

func (v *ViaCep) FindLocation(cep string) (string, error) {
	url := fmt.Sprintf(cep_url, cep)

	req, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("FAILED TO REQUEST CEP API: %v", err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return "", fmt.Errorf("FAILED TO READ CEP API RESPONSE: %v", err)
	}

	var data viaCEPResponse
	err = json.Unmarshal(res, &data)
	if err != nil {
		return "", fmt.Errorf("FAILED TO PARSE CEP API RESPONSE: %v", err)
	}

	if data.Erro != "" {
		return "", fmt.Errorf("CEP NOT FOUND")
	}

	return fmt.Sprintf("%s,%s", data.Localidade, data.Uf), nil
}
