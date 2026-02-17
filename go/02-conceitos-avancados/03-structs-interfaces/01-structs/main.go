package main

import (
	"encoding/json"
	"fmt"
)

// 1. Definição de Tipos e Structs
type User struct {
	Name string `json:"name"` // Struct Tags: definem o nome no JSON
	ID   uint64 `json:"id"`
}

func main() {
	// Inicialização por Nome (Recomendado: mais legível e permite omitir campos)
	user1 := User{Name: "Pedro Pessoa"} // ID será 0 (zero value)

	// Inicialização por Posição (Obriga todos os campos e é menos clara)
	user2 := User{"Joaquim", 20}

	// 2. Chamada de Métodos
	user1.Foo()

	// Alterando via Pointer Receiver (funciona com valor ou ponteiro devido ao pointer indirection)
	user1.UpdateName("Pedro Alterado")
	fmt.Println("Após UpdateName (Método):", user1.Name)

	// 3. Alterando via Função (Exige que o parâmetro seja exatamente o esperado)
	UpdateNameFunc(&user1, "Nome via Função")
	fmt.Println("Após UpdateName (Função):", user1.Name)

	// 4. JSON Marshal (usando as Tags)
	res, _ := json.Marshal(user2)
	fmt.Println("JSON Gerado:", string(res))
}

// --- MÉTODOS (Receivers) ---

// Value Receiver: Recebe uma CÓPIA. Útil para leitura (Performance: cópia pode ser lenta em structs grandes)
func (u User) Foo() {
	fmt.Println("Valor no Receiver:", u.Name)
}

// Pointer Receiver: Recebe o ENDEREÇO. Útil para alteração ou evitar cópias pesadas.
func (u *User) UpdateName(newName string) {
	u.Name = newName
}

// --- FUNÇÕES (Independentes) ---
func UpdateNameFunc(u *User, newName string) {
	u.Name = newName
}
