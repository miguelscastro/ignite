package main

import (
	"database/sql"
	"fmt"

	// O "_" realiza o registro do driver no pacote sql via função init()
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// sql.Open não abre uma conexão real imediatamente, ele apenas valida os argumentos
	// e prepara o pool de conexões para o uso futuro.
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close() // Boa prática: fechar o pool ao encerrar o programa

	createTableSQL := `
		CREATE TABLE IF NOT EXISTS foo (
		id integer not null primary key,
		name text
		);
	`

	// Exec é usado para comandos que não retornam linhas (DDL, INSERT, UPDATE, DELETE)
	res, err := db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

	insertSQL := `INSERT INTO foo (id, name) values (1, "miguel")`

	// res (Result) fornece metadados sobre a execução
	res, err = db.Exec(insertSQL)
	if err != nil {
		// Note: rodar 2x causará panic por chave duplicada no SQLite
		panic(err)
	}

	// RowsAffected retorna quantas linhas foram alteradas
	count, _ := res.RowsAffected()
	fmt.Println("Linhas afetadas:", count)

	type user struct {
		ID   int64
		Name string
	}

	// O caractere "?" é o placeholder para evitar SQL Injection (varia conforme o driver)
	querySQL := `SELECT id, name FROM foo WHERE id = ?;`

	var u user
	// QueryRow é otimizado para quando esperamos apenas UM resultado.
	// O Scan mapeia as colunas da query para os campos da struct usando ponteiros.
	err = db.QueryRow(querySQL, 1).Scan(&u.ID, &u.Name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Usuário recuperado:", u)
}
