package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// 1. Switch Básico (Break Implícito)
	fmt.Println("--- Switch Básico ---")
	testarValor(1)

	// 2. Fallthrough (Forçar a descida para o próximo case)
	fmt.Println("\n--- Usando Fallthrough ---")
	testarFallthrough(1)

	// 3. Switch sem condição (Substituto para If/Else longo)
	fmt.Println("\n--- Switch sem Condição ---")
	fmt.Println("É fim de semana?", isWeekend(time.Now()))

	// 4. Switch com Short Statement e Tipos Diferentes
	fmt.Println("\n--- Switch com Short Statement ---")
	switch x := math.Sqrt(4); x {
	case 2: // Comparando float64 (resultado da Sqrt) com literal int
		fmt.Println("Resultado é 2")
	default:
		fmt.Println("Algo deu errado")
	}

	// 5. Type Switch (Descobrir o tipo de uma interface)
	fmt.Println("\n--- Type Switch ---")
	identificarTipo("Olá Go")
	identificarTipo(100)
}

func testarValor(x int) {
	switch x {
	case 1:
		fmt.Println("Caiu no 1")
	case 2:
		fmt.Println("Caiu no 2")
	default:
		fmt.Println("Outro valor")
	}
}

func testarFallthrough(x int) {
	switch x {
	case 1:
		fmt.Println("Entrou no 1")
		fallthrough // Executa o case 2 mesmo sem validar a condição
	case 2:
		fmt.Println("Entrou no 2 via fallthrough")
	}
}

func isWeekend(t time.Time) bool {
	// Múltiplos valores no mesmo case
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		return true
	default:
		return false
	}
}

// x any (ou interface{}) aceita qualquer coisa
func identificarTipo(x any) {
	switch t := x.(type) {
	case string:
		fmt.Printf("É uma string: %q\n", t)
	case int:
		fmt.Printf("É um inteiro: %d\n", t)
	default:
		fmt.Printf("Tipo desconhecido: %T\n", t)
	}
}