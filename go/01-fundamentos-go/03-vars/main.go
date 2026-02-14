package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. Declaração e Inferência
	nome := "GoLang"
	var versao int = 1

	// Usando as variáveis para evitar erro de "unused var"
	fmt.Println("Estudo de:", nome, "Versão:", versao)

	// 2. Conversão de Tipos (Sempre Explícita)
	var x int = 65
	var f float64 = float64(x)

	// Conversão Direta: int -> string resulta no caractere Unicode
	s := string(x)
	fmt.Printf("Inteiro: %d | Float: %f | String Unicode: %s\n", x, f, s)

	// 3. Conversão Idiomática com strconv
	texto := "10084"
	// Atoi retorna (valor, erro). Usamos _ para ignorar o erro neste exemplo.
	numero, _ := strconv.Atoi(texto)
	fmt.Println("String para Int (Atoi):", numero)

	// 4. Constantes e Constantes Não Tipadas
	const constanteNaoTipada = 10
	const constanteTipada int32 = 10

	// Constantes não tipadas se adaptam ao contexto
	var a int64 = constanteNaoTipada   // Funciona
	var b float32 = constanteNaoTipada // Funciona

	// Já a tipada exige cast se o destino for diferente
	var c int64 = int64(constanteTipada)

	fmt.Println("Valores das constantes:", a, b, c)
}
