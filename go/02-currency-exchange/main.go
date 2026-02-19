package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ExchangeRates mapeia a estrutura do arquivo JSON.
// As tags `json:"..."` são essenciais para o Unmarshal saber qual chave do JSON
// deve preencher cada campo da struct.
type ExchangeRates struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"` // Usamos map para lidar com chaves dinâmicas (moedas).
}

func main() {
	// 1. Abertura do Arquivo: os.Open retorna um *os.File (implementa io.Reader).
	file, err := os.Open("exchange-rate.json")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo: ", err)
		return
	}
	// O defer garante que o arquivo seja fechado após o processamento, liberando recursos.
	defer file.Close()

	// 2. Decodificação JSON: O NewDecoder lê diretamente do fluxo de dados (Reader).
	var data ExchangeRates
	// Passamos o endereço (&data) para que o Decoder possa preencher a struct original via ponteiro.
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON: ", err)
		return
	}

	// 3. Argumentos de Linha de Comando: os.Args[0] é o nome do programa.
	// Esperamos os.Args[1] (valor) e os.Args[2] (moeda).
	args := os.Args
	if len(args) != 3 {
		fmt.Println("É necessário passar valor e moeda de destino")
		return
	}

	amountInBrl := args[1]
	// Convertemos para maiúsculas para bater com as chaves do map (ex: "usd" -> "USD").
	currencyTo := strings.ToUpper(args[2])

	// 4. Conversão de Tipos: Transformamos a string do argumento em float64.
	amount, err := strconv.ParseFloat(amountInBrl, 64)
	if err != nil {
		fmt.Println("Valor inválido, erro: ", err)
		return
	}

	// 5. Cálculo e Validação: Acessamos o map de taxas usando a moeda como chave.
	result := amount * data.Rates[currencyTo]

	// Se a moeda não existir no map, o Go retorna o valor zero do tipo (0.0 para float64).
	if result <= 0 {
		fmt.Println("Algo deu errado (Moeda não encontrada ou valor inválido)")
		return
	}

	// Exibimos o resultado formatado com 2 casas decimais.
	fmt.Printf("O valor convertido é: %.2f %s\n", result, currencyTo)
}
