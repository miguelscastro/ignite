package main

import "fmt"

func Slices1() {
	// --- 1. SLICES COMO PONTEIROS PARA ARRAYS ---
	// Um slice não armazena dados, ele apenas descreve uma seção de um array.
	arr := [5]int{1, 2, 3, 4, 5}

	// Criando um slice usando half-open interval [inferior:superior]
	// O índice 1 é incluído, o 4 é excluído (pega 1, 2 e 3).
	slice := arr[1:4]

	// Como o slice é um ponteiro, alterar um altera o outro.
	arr[2] = 15    // Altera o array original
	slice[0] = 123 // Altera o índice 1 do array original (que é o índice 0 do slice)

	fmt.Println("Array original:", arr) // [1, 123, 15, 4, 5]

	// --- 2. OMISSÃO DE LIMITES ---
	// Padrões: Inferior = 0 | Superior = tamanho total
	fatiaInteira := arr[:]
	fatiaInicio := arr[2:] // Do índice 2 até o fim
	fatiaFim := arr[:3]    // Do início até o índice 2 (3 excluído)

	fmt.Println("Fatias:", fatiaInteira, fatiaInicio, fatiaFim)

	// --- 3. SLICE LITERAL (Independente) ---
	// O Go cria um array escondido por baixo automaticamente.
	literal := []int{10, 20, 30}
	fmt.Println("Slice Literal:", literal)

	// --- 4. LEN (Tamanho) vs CAP (Capacidade) ---
	numeros := []int{1, 2, 3, 4, 5}
	resumo := numeros[:0] // Reduzimos o tamanho para 0, mas mantemos a capacidade

	// Saída: [] 0 5
	fmt.Printf("Slice2: %v | Len: %d | Cap: %d\n", resumo, len(resumo), cap(resumo))

	// --- 5. NIL SLICE vs SLICE VAZIO ---
	// Var declarada sem valor não aponta para nada (nil).
	var sNil []int
	fmt.Println("sNil é nulo?", sNil == nil) // true

	// Slice literal vazio {} aponta para um endereço mas não tem elementos.
	sVazio := []int{}
	fmt.Println("sVazio é nulo?", sVazio == nil) // false
}
