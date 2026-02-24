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

	createTableSQL := `--sql
		CREATE TABLE IF NOT EXISTS foo (
		id integer not null primary key,
		name text
		);
	`

	// Exec é usado para comandos que não retornam linhas (DDL, INSERT, UPDATE, DELETE)
	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

	insertSQL := `--sql 
		INSERT INTO foo (id, name) values (1, "miguel")`

	// res (Result) fornece metadados sobre a execução
	res, err := db.Exec(insertSQL)
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
	// --- MÉTODO SEGURO: QueryRow com Placeholders ---
	querySQL := `
	--sql
	SELECT id, name FROM foo WHERE id = ?;`

	var u user
	err = db.QueryRow(querySQL, 1).Scan(&u.ID, &u.Name)
	if err != nil {
		fmt.Println("Erro no Select:", err)
	} else {
		fmt.Println("Usuário recuperado:", u)
	}

	// --- DEMONSTRAÇÃO DE RISCO: SQL Injection ---

	// Exemplo de input malicioso que um atacante poderia enviar
	// Em vez de apenas "1", ele envia "1 OR 1=1"
	inputMalicioso := "1 OR 1=1"

	// PERIGOSO: O fmt.Sprintf apenas concatena strings.
	// O driver recebe: DELETE FROM foo WHERE id = 1 OR 1=1;
	// Isso apagaria a tabela INTEIRA.
	deleteInseguro := fmt.Sprintf(`
		--sql
		DELETE FROM foo WHERE id = %s;`, inputMalicioso)

	_ = deleteInseguro
	fmt.Println("Executando query insegura...")
	// _, err = db.Exec(deleteInseguro) // Comentado para segurança do seu DB local

	// --- MÉTODO CORRETO: Exec com Args ---

	deleteSeguro := `
		--sql
		DELETE FROM foo WHERE id = ?;`

	// O driver trata o valor separadamente da estrutura da query.
	// Se passarmos "1 OR 1=1", o driver tentará encontrar um ID que seja
	// literalmente essa string, falhando com segurança.
	if _, err := db.Exec(deleteSeguro, 1); err != nil {
		panic(err)
	}
	fmt.Println("Delete seguro executado.")
}
