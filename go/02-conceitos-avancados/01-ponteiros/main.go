package main

import "fmt"

func main() {
	// 1. Conceito Básico: Endereço (&) e Desreferenciação (*)
	x := 10
	p := &x
	fmt.Printf("Endereço: %p | Valor via ponteiro: %d\n", p, *p)

	// 2. Passagem por Referência (Mutate)
	fmt.Println("Antes do take:", x)
	take(&x)
	fmt.Println("Depois do take:", x)

	// 3. O Escape Analysis (Retornando ponteiro local)
	// Em C++ isso seria um erro fatal. Em Go, o compilador move 'x' para a Heap.
	ptrCriado := create()
	fmt.Println("Valor criado na Heap via Escape Analysis:", *ptrCriado)

	// 4. PERIGO: Ponteiro Nulo (nil)
	// Descomentar a linha abaixo causará um PANIC em runtime.
	// exemploPanic()
}

func take(x *int) {
	if x != nil {
		*x = 100
	}
}

func create() *int {
	// Variável 'y' parece local (stack), mas como retornamos o endereço,
	// o Go a promove para a Heap.
	y := 42
	return &y
}

func exemploPanic() {
	var nulo *int
	// Tentar desreferenciar um ponteiro nil gera erro de execução.
	// Sempre verifique se o ponteiro é != nil antes de usar.
	fmt.Println(*nulo)
}
