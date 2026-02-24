package main

import (
	"context"
	"fmt"

	// O pgxpool é o padrão para gerenciar múltiplas conexões simultâneas
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// DSN (Data Source Name): Define o protocolo, credenciais, endereço e banco
	// sslmode=disable é necessário para conexões locais sem certificado SSL
	databaseURL := "postgres://postgres:root@localhost:5432/tests"

	// pgxpool.New: Cria o gerenciador de conexões
	// O contexto permite definir um timeout para a tentativa de criação do pool
	db, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		panic(err) // Em produção, aqui você faria log.Fatal
	}
	// Importante: Fechar o pool ao encerrar o programa para liberar recursos do SO
	defer db.Close()

	// db.Ping: Verifica se a conexão física realmente funciona e se o banco está pronto
	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	// DROP TABLE IF EXISTS: Comando DDL seguro para resetar o ambiente de teste
	resetDB := `--sql
		DROP TABLE IF EXISTS foo
	`
	if _, err := db.Exec(context.Background(), resetDB); err != nil {
		panic(err)
	}

	// CREATE TABLE: id bigserial cria um inteiro de 64 bits auto-incremental
	query := `CREATE TABLE foo (id bigserial primary key, bar varchar(255));`
	if _, err := db.Exec(context.Background(), query); err != nil {
		panic(err)
	}

	// INSERT: Utilizamos $1 (placeholder) para prevenir SQL Injection
	// O driver do Postgres não aceita "?" como o MySQL/SQLite
	query = `INSERT INTO foo (bar) VALUES ($1)`
	if _, err := db.Exec(context.Background(), query, "abcd"); err != nil {
		panic(err)
	}

	// SELECT: LIMIT 1 é usado quando queremos apenas o primeiro registro
	query = `SELECT id, bar FROM foo LIMIT 1;`

	type foobar struct {
		id  int64
		bar string
	}

	var res foobar
	// db.QueryRow: Otimizado para consultas que retornam exatamente uma linha
	// .Scan(): Mapeia as colunas da query para os endereços de memória (&) da nossa struct
	if err := db.QueryRow(context.Background(), query).Scan(&res.id, &res.bar); err != nil {
		panic(err)
	}

	// %+v imprime os nomes dos campos da struct, %# imprime o tipo
	fmt.Printf("%#+v\n", res)
}
