package main

import (
	"fmt"
)

func main() {
	// 1. Comportamento Básico e Retorno
	fmt.Println("--- Exemplo 1: Fluxo de Retorno ---")
	valor := exemploRetorno()
	fmt.Println("Valor retornado para main:", valor)

	// 2. LIFO (Last In, First Out)
	fmt.Println("\n--- Exemplo 2: Pilha LIFO ---")
	pilhaDefer()

	// 3. Snapshot de Parâmetros
	fmt.Println("\n--- Exemplo 3: Snapshot de Variáveis ---")
	snapshotDefer()

	// 4. Defer em Loops (O jeito correto)
	fmt.Println("\n--- Exemplo 4: Defer em Loop com Closure ---")
	deferEmLoopCorreto()
}

func exemploRetorno() int {
	defer fmt.Println("  [Defer] Executado após o return")
	fmt.Println("  Executando lógica da função...")
	return 10
}

func pilhaDefer() {
	defer fmt.Println("  Primeiro defer declarado (sai por último)")
	defer fmt.Println("  Segundo defer declarado (sai por primeiro)")
	fmt.Println("  Função principal rodando...")
}

func snapshotDefer() {
	x := 10
	// O valor de x (10) é copiado para o argumento 'y' IMEDIATAMENTE.
	defer func(y int) {
		fmt.Printf("  [Defer] Valor de y capturado: %d\n", y)
	}(x)

	x = 50
	fmt.Printf("  Valor de x alterado para: %d\n", x)
}

func deferEmLoopCorreto() {
	arquivos := []string{"teste1.txt", "teste2.txt"}

	for _, f := range arquivos {
		// Criamos um escopo de função anônima para que o defer
		// execute ao fim de cada iteração, e não ao fim da função pai.
		func(nome string) {
			fmt.Printf("  Abrindo arquivo: %s\n", nome)
			// Simulação de abertura
			// file, _ := os.Open(nome) 
			defer fmt.Printf("  [Defer] Fechando arquivo: %s\n", nome)
		}(f)
	}
}