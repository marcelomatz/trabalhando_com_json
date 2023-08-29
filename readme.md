# Projeto JSON Benchmark em Go
## Benchmark Go: Marshal x Encoder com requisição http

Este projeto é um simples exemplo de como podemos benchmark diferentes maneiras de lidar com JSON em Go. Ele testa a eficiência do pacote padrão encoding/json e da biblioteca jsoniter.

### Descrição
Temos um struct chamado Info que representa a estrutura do JSON que recebemos de uma chamada à API.

```go
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
```
Há também uma função fetchData(url string) que faz uma chamada à API e retorna o corpo da resposta em bytes.
Os benchmarks são feitos para as funções:
* json.Unmarshal (BenchmarkEncodingJsonUnmarshal)
* json.NewDecoder (BenchmarkEncodingJsonDecoder)
* jsoniter.Unmarshal (BenchmarkJsoniterUnmarshal)
* jsoniter.NewDecoder (BenchmarkJsoniterDecoder)

### Objetivo
Nosso objetivo é testar a eficiência de cada pacote/método. As métricas do benchmark incluem tempo médio por operação e número de operações realizadas.

### Uso
Para rodar os benchmarks, execute:
```bash
go test -bench=.
```
Isso irá rodar todos os benchmarks definidos no pacote atual.

### Resultados
Os benchmarks mostraram que a biblioteca **jsoniter** é mais rápida do que o pacote padrão encoding/json. 

No entanto, outras considerações, como a estabilidade da API e a quantidade de recursos disponíveis, devem ser levadas em conta ao decidir qual pacote usar
