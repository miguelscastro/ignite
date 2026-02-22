package main

import (
	"fmt"
)

// --- 1. DEFINIÇÃO E IMPLEMENTAÇÃO IMPLÍCITA ---

// Interface define o CONTRATO. Não há "implements" em Go.
type Animal interface {
	Sound() string
}

type Dog struct {
	Breed string
}

// Dog implementa Animal automaticamente pois possui o método Sound() string.
func (d *Dog) Sound() string {
	// 3. TRATATIVA DE NIL NO RECEIVER
	// Permite que a interface chame o método mesmo que o valor seja nulo.
	if d == nil {
		return "<silêncio: cachorro nulo>"
	}
	return "Au! Au!"
}

// Exemplo de polimorfismo: aceita qualquer tipo que satisfaça 'Animal'
func whatDoesThisAnimalSay(a Animal) {
	fmt.Println("O animal diz:", a.Sound())
}

// --- 2. IMPLEMENTAÇÃO PARA PACOTES EXTERNOS ---
// Supondo que o pacote 'foo' tenha: type Interface interface { Interface() }
func (d *Dog) Interface() {
	fmt.Println("Dog satisfazendo uma interface de outro pacote!")
}

func Interfaces1() {
	fmt.Println("=== 1. Implementação Implícita ===")
	dog := &Dog{Breed: "Labrador"}
	whatDoesThisAnimalSay(dog) // Dog 'é um' Animal

	fmt.Println("\n=== 2. Nulidade de Interfaces ===")
	var a Animal // Interface vazia (Tipo: nil, Valor: nil)
	var d *Dog   // Ponteiro nulo (Valor: nil)

	fmt.Printf("A interface 'a' é nil? %v\n", a == nil) // true

	// A ARMADILHA: Atribuindo um ponteiro nil à interface
	a = d

	// Agora 'a' guarda o TIPO (*Dog), mas o VALOR é nil.
	// REGRA: Interface só é nil se Tipo E Valor forem nil.
	fmt.Printf("Após a=d, 'a' é nil? %v\n", a == nil) // false!

	// Isso funciona porque o método Sound() trata receiver nulo internamente.
	// Se não tratasse e acessasse d.Breed, daria PANIC (invalid memory address).
	fmt.Println("Chamada segura em interface não-nil com valor nil:", a.Sound())

	fmt.Println("\n=== 3. Interface Vazia (any) ===")
	// Aceita qualquer coisa pois qualquer tipo tem no mínimo ZERO métodos.
	takeEmptyInterface("Texto")
	takeAnything(100)
	takeAnything(dog)
}

// any é um alias para interface{}
func takeAnything(a any) {
	fmt.Printf("Recebi tipo: %T com valor: %v\n", a, a)
}

func takeEmptyInterface(i interface{}) {
	fmt.Println("Interface vazia recebeu:", i)
}
