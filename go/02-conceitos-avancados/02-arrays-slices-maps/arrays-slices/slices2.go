package main

import "fmt"

func Slices2() {
	filmesNoDB := []string{
		"O Poderoso Chefão",
		"Titanic",
		"O Senhor dos Anéis: O Retorno do Rei",
		"Matrix",
		"Forrest Gump: O Contador de Histórias",
		"O Rei Leão",
		"Harry Potter e a Pedra Filosofal",
		"Gladiador",
		"O Sexto Sentido",
		"O Curioso Caso de Benjamin Button",
		"Pulp Fiction: Tempo de Violência",
		"Esqueceram de Mim",
		"O Exterminador do Futuro 2: O Julgamento Final",
		"O Fabuloso Destino de Amélie Poulain",
		"O Labirinto do Fauno",
	}

	resultsFromAPI := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	filmes := make([]string, 0, 10)

	for _, id := range resultsFromAPI {
		filme := filmesNoDB[id]
		filmes = append(filmes, filme)
	}
	fmt.Println(filmes)

	// matrix := [][]int{}
	// matrix3d := [][][]int{}
}
