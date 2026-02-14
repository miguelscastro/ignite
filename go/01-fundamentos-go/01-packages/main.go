// 01-packages/main.go
package main

/*
  DICA DE ENGENHARIA:
  O package 'main' é o único que gera um executável.
  Tudo o que você definir aqui não poderá ser importado por outros projetos.
*/

import (
	"fmt"
	// Alias Import: Útil para evitar conflitos ou encurtar nomes gigantes.
	f "fmt"

	// Blank Import (_): Executa apenas a função init() do pacote.
	// Muito usado em drivers de banco de dados ou configurações globais.
	_ "image/png"
	// Dot Import (.): Importa o pacote para o namespace atual.
	// CUIDADO: O 'Head of Engineering' rejeitaria isso em um Code Review profissional!
	// . "math"
)

func main() {
	// 1. Usando Alias
	f.Println("Explorando pacotes em Go!")

	// 2. Visibilidade (Exported vs Unexported)
	exibirConceitoVisibilidade()
}

// exibirConceitoVisibilidade começa com letra minúscula, então é PRIVADA.
// Só pode ser chamada por outros arquivos dentro desta mesma pasta (package main).
func exibirConceitoVisibilidade() {
	fmt.Println("--- Regras de Ouro ---")
	fmt.Println("1. Maiúscula (Nome): Público/Exportado")
	fmt.Println("2. Minúscula (nome): Privado/Interno")
	fmt.Println("3. 'internal/': Pacotes protegidos que só vizinhos podem usar.")
}
