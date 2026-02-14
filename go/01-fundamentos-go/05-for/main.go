package main

import "fmt"

func main() {
	// 1. For Clássico (C-style)
	fmt.Println("--- For Clássico ---")
	contagem := 0
	for i := 0; i < 5; i++ {
		contagem++
	}
	fmt.Println("Resultado do contador:", contagem)

	// 2. For como While
	fmt.Println("\n--- For como While ---")
	j := 0
	for j < 5 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()

	// 3. For Range (Iteração sobre coleções)
	fmt.Println("\n--- For Range ---")
	numeros := [3]int{10, 20, 30}
	for i, v := range numeros {
		fmt.Printf("Índice: %d | Valor: %d\n", i, v)
	}

	// 4. For Range sobre Inteiro (Go 1.22+)
	fmt.Println("\n--- For Range sobre Inteiro ---")
	for range 3 {
		fmt.Println("Repetição simplificada")
	}

	// 5. O Perigo da Cópia e Endereço de Memória
	fmt.Println("\n--- Endereços de Memória no Range ---")
	analisarMemoriaRange()
}

func analisarMemoriaRange() {
	arr := [3]int{1, 2, 3}
	
	// 'i' e 'v' são declarados uma vez e reutilizados.
	// O valor muda, mas o endereço de memória da variável de iteração é o mesmo.
	for i, v := range arr {
		fmt.Printf("Valor: %d | Endereço de i: %p | Endereço de v: %p\n", v, &i, &v)
	}
}