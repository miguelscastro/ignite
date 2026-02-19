package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Measurement armazena o estado da agregação para cada localidade.
type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	// Abre o arquivo de medições. O retorno é um *os.File, que implementa io.Reader.
	measurements, err := os.Open("measurements.txt")
	if err != nil {
		panic(err) // Em Go, panic é usado apenas para erros fatais irrecuperáveis.
	}
	// Garante que o descritor do arquivo seja liberado ao fim da função main.
	defer measurements.Close()

	// Map para agrupar as medições por nome da localidade (chave).
	dados := make(map[string]Measurement)

	// bufio.Scanner é ideal para ler arquivos linha por linha de forma eficiente (Streaming).
	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rawData := scanner.Text()                // Obtém a linha atual como string.
		semicolon := strings.Index(rawData, ";") // Localiza o delimitador ponto e vírgula.

		// Fatia a string (Slice) para extrair localidade e temperatura sem criar novas strings pesadas.
		location := rawData[:semicolon]
		rawTemperature := rawData[semicolon+1:]

		// Converte a temperatura de string para float64.
		temperature, _ := strconv.ParseFloat(rawTemperature, 64)

		// Verifica se a localidade já existe no map (idioma "comma ok").
		measurement, ok := dados[location]
		if !ok {
			// Se não existe, inicializa com os valores da primeira leitura.
			measurement = Measurement{
				Min:   temperature,
				Max:   temperature,
				Sum:   temperature,
				Count: 1,
			}
		} else {
			// Se existe, atualiza os valores acumulados (agregação).
			measurement.Min = min(measurement.Min, temperature)
			measurement.Max = max(measurement.Max, temperature)
			measurement.Sum += temperature
			measurement.Count++
		}

		// Importante: structs em maps são copiadas, então reatribuímos para atualizar o valor real.
		dados[location] = measurement
	}

	// Criamos um slice para ordenar as chaves (localidades) alfabeticamente.
	locations := make([]string, 0, len(dados))
	for name := range dados {
		locations = append(locations, name)
	}
	sort.Strings(locations) // Ordenação in-place do slice de strings.

	// Formatação final da saída.
	fmt.Printf("{")
	for i, name := range locations {
		measurement := dados[name]
		// Calcula a média em tempo de execução: soma / quantidade.
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f",
			name,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)

		// Adiciona vírgula apenas entre os itens, não no último.
		if i < len(locations)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("}\n")
}
