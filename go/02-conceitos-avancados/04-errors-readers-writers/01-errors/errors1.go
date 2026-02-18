package main

import (
	"errors"
	"fmt"
)

// --- SEÇÃO 1: ERROS COMO VALORES ---
// Em Go, não causamos exceptions (panic), retornamos erros.
// Erro é uma interface padrão: type error interface { Error() string }

type User struct{ foo string }

func (u User) Foo() {
	fmt.Println(u.foo)
}

// NewUser demonstra a convenção: retornar o valor zero (nil) e o erro.
// O retorno *User é um ponteiro (endereço).
func NewUser(wantError bool) (*User, error) {
	if wantError {
		// Retornamos nil para o ponteiro e um valor para o erro.
		return nil, errors.New("erro ao criar usuário")
	}
	// &User{} gera o endereço para cumprir o contrato *User.
	return &User{foo: "valor de teste"}, nil
}

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("não pode dividir por zero") // 0 é o valor zero de int.
	}
	return a / b, nil
}

func Errors1() {
	// --- Exemplo de Divisão ---
	res, errDiv := dividir(10, 0)
	if errDiv != nil {
		fmt.Println("Erro de matemática:", errDiv)
	} else {
		fmt.Println("Resultado:", res)
	}

	// --- Exemplo de User (Pointer e Nil check) ---
	user, errUser := NewUser(true)

	// A tratativa abaixo impede o panic().
	// Se ignorássemos o erro, user.Foo() tentaria acessar 'foo' em um endereço nil.
	if errUser != nil {
		fmt.Println("Tratativa ativa: Não podemos prosseguir com user nil.")
	} else {
		user.Foo() // Pointer Indirection: Go desreferencia user automaticamente.
	}
}
