package main

import "fmt"

func Slices2() {
	// --- BASE DE DADOS E SIMULAÇÃO DE API ---
	filmesNoDB := []string{
		"O Poderoso Chefão", "Titanic", "O Senhor dos Anéis", "Matrix",
		"Forrest Gump", "O Rei Leão", "Harry Potter", "Gladiador",
		"O Sexto Sentido", "Benjamin Button", "Pulp Fiction",
		"Esqueceram de Mim", "Exterminador do Futuro 2", "Amélie Poulain",
		"O Labirinto do Fauno",
	}

	resultsFromAPI := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// --- 1. A FUNÇÃO append() (Dinâmica) ---
	// O append adiciona elementos e gerencia o crescimento da memória.
	// Usamos make com len 0 e cap 10 para evitar realocações desnecessárias.
	filmes := make([]string, 0, 10)

	for _, id := range resultsFromAPI {
		filme := filmesNoDB[id]
		// O append não é "in place"; ele retorna um novo slice header.
		filmes = append(filmes, filme)
	}
	fmt.Println("Slice após Appends:", filmes)

	// --- 2. A FUNÇÃO copy() (Rígida/Segura) ---
	// O copy é usado para criar clones independentes (Deep Copy).
	// REGRA DA PROVA: Ele copia apenas o MENOR tamanho entre origem e destino.

	// Caso A: Destino com tamanho exato (Cópia completa)
	backupCompleto := make([]string, len(filmes))
	n1 := copy(backupCompleto, filmes)
	fmt.Printf("\nCopy Total: Copiou %d itens para o backup.\n", n1)

	// Caso B: Destino menor que a origem (Truncamento)
	// Se o destino tem tamanho 3, o Go copia apenas os 3 primeiros e NÃO aumenta o slice.
	destinoCurto := make([]string, 3)
	n2 := copy(destinoCurto, filmes)
	fmt.Printf("Copy Curto: Copiou apenas %d itens. Conteúdo: %v\n", n2, destinoCurto)

	// --- 3. PROVANDO A INDEPENDÊNCIA (Passagem por Referência) ---
	// Como usamos copy(), alterar o backup não afeta o original.
	backupCompleto[0] = "ALTERADO NO BACKUP"
	fmt.Println("\nOriginal[0]:", filmes[0])     // Continua original
	fmt.Println("Backup[0]:", backupCompleto[0]) // Alterado apenas aqui

	// --- 4. MATRIZES (Slices de Slices) ---
	// matrix := [][]int{}
	// matrix3d := [][][]int{}
}
