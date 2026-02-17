package main

import "fmt"

func Slices3() {
	fmt.Println("--- Aula 3: Full Slice, Bounds Check e Performance ---")

	// 1. Full Slice Expression [baixo:alto:max]
	// O terceiro índice controla a CAPACIDADE do novo slice.
	// Regra: baixo <= alto <= max <= cap(origem)
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:3:4] // Pega índices 1 e 2, mas limita a capacidade até o índice 4.

	fmt.Printf("Full Slice: %v | Len: %d | Cap: %d\n", slice, len(slice), cap(slice))

	// 2. Bounds Check & Performance (Hints)
	// O Go verifica todos os acesso a índice para evitar que você acesse memória errada.
	numeros := []int{10, 20, 30, 40}

	// Dica para o compilador (Hint):
	// Ao acessar o índice 3 primeiro, o compilador entende que os índices
	// 0, 1 e 2 também são seguros, eliminando 3 verificações automáticas.
	_ = numeros[3]
	fmt.Println(numeros[0], numeros[1], numeros[2], numeros[3])

	// 3. Passagem por Referência vs Valor
	fmt.Println("Original antes de foo:", numeros)
	foo(numeros) // Slices passam o cabeçalho (aponta para a mesma memória)
	fmt.Println("Original depois de foo:", numeros)

	copiaArr := [4]int{1, 2, 3, 4}
	foo2(copiaArr) // Arrays passam uma CÓPIA de todos os dados
	fmt.Println("Array original após foo2 (não muda):", copiaArr)
}

func foo(s []int) {
	s[0] = 999
}

func foo2(a [4]int) {
	a[0] = 999
}
