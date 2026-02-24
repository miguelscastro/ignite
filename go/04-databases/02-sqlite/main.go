package main

import (
	"fmt"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

func main() {
	// ":memory:" é uma funcionalidade do SQLite que mantém o banco apenas na RAM
	// Útil para testes unitários ultrarrápidos ou caches temporários
	conn, err := sqlite.OpenConn(":memory:", sqlite.OpenReadWrite)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// ExecuteTransient é uma forma rápida de rodar SQL sem preparar um Statement manualmente
	// ResultFunc atua como um callback para cada linha retornada
	err = sqlitex.ExecuteTransient(conn, "SELECT 'hello, world';", &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			// ColumnText(0) acessa diretamente o valor da primeira coluna
			fmt.Println(stmt.ColumnText(0))
			return nil
		},
	})
	if err != nil {
		panic(err)
	}
}
