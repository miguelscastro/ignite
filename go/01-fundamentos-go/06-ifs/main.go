package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. IF Básico
	// Diferença: Sem parênteses (), mas chaves {} são obrigatórias.
	if 1 < 2 {
		fmt.Println("Um é menor que dois")
	}

	// 2. Short Statement (Declaração curta + Condição)
	// A variável 'x' só existe dentro do bloco if/else if/else.
	if x := math.Sqrt(4); x < 1 {
		fmt.Println("Raiz menor que 1:", x)
	} else if x > 0 {
		fmt.Println("Raiz maior que 0:", x)
	} else {
		fmt.Println("Caiu no else")
	}

	// 3. Demonstração de Erro de Escopo (Descomente para testar)
	// fmt.Println(x) // Erro: undefined: x

	analisarComportamento()
}

func analisarComportamento() {
	fmt.Println("\n--- Dica de Engenharia ---")
	fmt.Println("Use short statements para variáveis temporárias.")
	fmt.Println("Isso evita que você use acidentalmente um valor obsoleto mais tarde na mesma função.")
}