package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type Info struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}

func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func BenchmarkEncodingJsonUnmarshal(b *testing.B) {
	data, _ := fetchData("https://viacep.com.br/ws/01001000/json/")
	info := &Info{}
	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal(data, &info); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncodingJsonDecoder(b *testing.B) {
	data, _ := fetchData("https://viacep.com.br/ws/01001000/json/")
	info := &Info{}
	for i := 0; i < b.N; i++ {
		if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(&info); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJsoniterUnmarshal(b *testing.B) {
	data, _ := fetchData("https://viacep.com.br/ws/01001000/json/")
	info := &Info{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal(data, &info); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJsoniterDecoder(b *testing.B) {
	data, _ := fetchData("https://viacep.com.br/ws/01001000/json/")
	info := &Info{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	for i := 0; i < b.N; i++ {
		if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(&info); err != nil {
			b.Fatal(err)
		}
	}
}
