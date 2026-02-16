package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. Inicialização e Erro de Nil Map
	// var m map[string]string
	// m["Pedro"] = "Pessoa" // PANIC: assignment to entry in nil map

	m := make(map[string]string, 10) // Inicializado com capacidade 10
	m["Pedro"] = "Pessoa"

	// 2. Comma ok idiom (Acesso seguro)
	valor, ok := m["Joaquim"]
	if !ok {
		fmt.Println("Chave não encontrada. Valor retornado é o 'zero value':", valor)
	}

	// 3. O problema do NaN (Not a Number)
	f := math.NaN()
	mapNaN := map[float64]string{
		f: "Valor Fantasma",
	}

	// Tentar acessar ou deletar via chave falha porque NaN != NaN
	_, ok = mapNaN[f]
	fmt.Println("Conseguiu achar o NaN via chave?", ok) // false

	// A única solução para limpar um mapa com NaNs é o clear()
	clear(mapNaN)
	fmt.Println("Mapa após clear:", mapNaN)

	// 4. Iteração e Deleção segura
	usuarios := map[string]string{
		"Pedro": "Admin",
		"Ana":   "User",
	}

	for k := range usuarios {
		if k == "Pedro" {
			delete(usuarios, k) // Seguro deletar durante o range em Go
		}
	}
}
