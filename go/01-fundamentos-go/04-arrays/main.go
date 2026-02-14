package main

import "fmt"

func main() {
	// 1. Declaração Simples
	// O tamanho faz parte do tipo. [3]int != [5]int.
	var a [3]int
	a[0] = 10
	a[1] = 20
	a[2] = 30
	fmt.Println("Array simples:", a)

	// 2. Inicialização Curta e Elipse (...)
	// O compilador conta quantos elementos existem para definir o tamanho.
	b := [...]string{"Go", "Python", "Java"}
	fmt.Printf("Array B: %v | Tamanho: %d | Tipo: %T\n", b, len(b), b)

	// 3. Inicialização por Índice (Arrays Esparsos)
	// Útil para quando você sabe a posição exata de alguns valores.
	exemploIndices()
}

func exemploIndices() {
	// Array de 10 inteiros: índice 5 é 200, índice 7 é 400.
	// Os demais são preenchidos com o 'zero value' (0).
	arrInt := [10]int{5: 200, 7: 400}
	fmt.Println("Array por índice (int):", arrInt)

	// Array de 10 strings: índice 5 recebe valor, demais ficam vazios ("").
	arrString := [10]string{5: "Hello World"}
	fmt.Printf("Array por índice (string): %q\n", arrString) // %q mostra as aspas
}
