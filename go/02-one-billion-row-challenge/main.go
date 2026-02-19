package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	measurements, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}
	defer measurements.Close()

	dados := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		rawTemperature := rawData[semicolon+1:]

		temperature, _ := strconv.ParseFloat(rawTemperature, 64)

		measurement, ok := dados[location]
		if !ok {
			measurement = Measurement{
				Min:   temperature,
				Max:   temperature,
				Sum:   temperature,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temperature)
			measurement.Max = max(measurement.Max, temperature)
			measurement.Sum += temperature
			measurement.Count++
		}

		dados[location] = measurement
	}

	locations := make([]string, 0, len(dados))
	for name := range dados {
		locations = append(locations, name)
	}
	sort.Strings(locations)

	fmt.Printf("{")
	for _, name := range locations {
		measurement := dados[name]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f, ",
			name,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max)
	}
	fmt.Printf("}\n")
}
