package main

import "fmt"

func main() {
	// 1. Chamada Simples e Múltiplos Retornos
	soma := somar(10, 20)
	nome, sobrenome := obterNomeCompleto()
	fmt.Printf("Soma: %d | Nome: %s %s\n", soma, nome, sobrenome)

	// 2. Variádica (O parâmetro vira um Slice []int lá dentro)
	total := somarVarios(1, 2, 3, 4, 5)
	fmt.Println("Total Variádica:", total)

	// 3. Closure / Higher-Order Function
	// O 'multiplicador' captura o valor 2 e "guarda" com ele.
	duplicar := criarMultiplicador(2)
	fmt.Println("Dobro de 10:", duplicar(10)) 
}

// somar demonstra a simplificação de tipos (a, b int)
func somar(a, b int) int {
	return a + b
}

// obterNomeCompleto retorna múltiplos valores (comum para retornar valor + erro)
func obterNomeCompleto() (string, string) {
	return "Engenharia", "Whirlpool"
}

// somarVarios é uma função variádica. 
// Internamente, 'nums' é tratado como uma lista (slice).
func somarVarios(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// criarMultiplicador retorna uma função (Higher-Order)
// A função interna "lembra" da variável 'fator' (Closure).
func criarMultiplicador(fator int) func(int) int {
	return func(n int) int {
		return n * fator
	}
}