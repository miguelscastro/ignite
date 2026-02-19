package main

import "fmt"

// Constraint customizada com Aproximação (~)
type MyConstraint interface {
	~int | ~string | ~[]int
}

type MeuTipo string // Core type é string

// Função Genérica com Constraint
func imprimir[T MyConstraint](arg T) {
	fmt.Println("Valor:", arg)
}

// Struct Genérica
type MyStruct[T any] struct {
	Dado T
}

// Método de Struct Genérica (usa o T da struct)
func (s MyStruct[T]) Processar(val T) {
	fmt.Printf("Processando %T: %v\n", val, val)
}

// Exemplo Prático: Contains
func Contains[T comparable](s []T, cmp T) bool {
	for _, v := range s {
		if v == cmp {
			return true
		}
	}
	return false
}

func main() {
	// Uso da Aproximação (~)
	var mt MeuTipo = "Sou um core type string"
	imprimir(mt)  // Funciona por causa do ~string
	imprimir(100) // Funciona por causa do ~int

	// Instanciando Struct Genérica
	ms := MyStruct[int]{Dado: 10}
	ms.Processar(20)

	// Testando Contains
	nomes := []string{"Go", "Generics", "API"}
	fmt.Println("Existe 'Go'?", Contains(nomes, "Go"))
}
