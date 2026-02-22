package main

import (
	"fmt"
)

// --- 1. INTERFACE STRINGER (BIBLIOTECA PADRÃO) ---

type Pedro struct{}

// Implementação implícita da interface fmt.Stringer.
// O pacote fmt usa isso internamente para decidir como imprimir o tipo.
func (Pedro) String() string {
	return "Sou a struct Pedro, mas o Println me chama de 'teste'"
}

// --- 2. TIPO E MÉTODO PARA TESTE DE POINTER RECEIVER ---

type Cat struct {
	Name string
}

// Definido com POINTER RECEIVER (*Cat)
func (c *Cat) Sound() string {
	if c == nil {
		return "miau (silencioso)"
	}
	return "Miau!"
}

// --- 3. TYPE ASSERTIONS E TYPE SWITCHES ---

func processAnything(a any) {
	fmt.Printf("\n--- Processando tipo: %T ---\n", a)

	// TYPE ASSERTION: "Eu acho que esse 'any' é uma string"
	// str recebe o valor, ok recebe true se for realmente string
	str, ok := a.(string)
	if ok {
		fmt.Println("Sucesso no Type Assertion! Valor:", str)
	}

	// TYPE SWITCH: Forma limpa de tratar múltiplos tipos concretos
	switch t := a.(type) {
	case *Cat:
		// Aqui 't' já é do tipo *Cat (ponteiro), podemos acessar Sound()
		fmt.Println("Type Switch detectou um Gato:", t.Sound())
	case Pedro:
		// Aqui 't' é do tipo Pedro (valor)
		fmt.Println("Type Switch detectou Pedro:", t.String())
	case int:
		fmt.Println("Recebi um número inteiro:", t)
	case nil:
		fmt.Println("Recebi um valor nil")
	default:
		fmt.Println("Tipo desconhecido pelo switch")
	}
}

func Interfaces2() {
	// Exemplo Stringer
	p := Pedro{}
	fmt.Println("=== Exemplo Stringer ===")
	fmt.Println(p) // fmt.Println detecta o método String() implicitamente

	// Exemplo Type Switch e Method Sets
	fmt.Println("\n=== Exemplo Type Switch ===")

	gato := &Cat{Name: "Garfield"}
	processAnything(gato) // Passando o ponteiro (*Cat)

	// Por que não passar Cat direto?
	// Como Sound() tem receiver (*Cat), a struct pura Cat não satisfaz Animal.

	processAnything("Uma string qualquer")
	processAnything(100)
}
