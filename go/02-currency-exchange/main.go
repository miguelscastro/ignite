package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ExchangeRates struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	file, err := os.Open("exchange-rate.json")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo: ", err)
		return
	}
	defer file.Close()

	var data ExchangeRates
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON: ", err)
		return
	}

	args := os.Args
	if len(args) != 3 {
		fmt.Println("É necessário passar valor e moeda de destino")
		return
	}

	amountInBrl := args[1]
	currencyTo := strings.ToUpper(args[2])

	amount, err := strconv.ParseFloat(amountInBrl, 64)
	if err != nil {
		fmt.Println("Valor inválido, erro: ", err)
		return
	}

	result := amount * data.Rates[currencyTo]
	if result <= 0 {
		fmt.Println("Algo deu errado")
	}
	fmt.Printf("O valor convertido é: %.2f %s\n", result, currencyTo)
}
