package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/tests")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	query := `CREATE TABLE foo (id bigint auto_increment primary key, bar varchar(255));`
	if _, err := db.Exec(query); err != nil {
		panic(err)
	}

	query = `INSERT INTO foo (bar) VALUES (?)`
	if _, err := db.Exec(query, "abcd"); err != nil {
		panic(err)
	}

	query = `SELECT * FROM foo LIMIT 1;`

	type foobar struct {
		id  int64
		bar string
	}

	var res foobar
	if err := db.QueryRow(query).Scan(&res.id, &res.bar); err != nil {
		panic(err)
	}
	fmt.Printf("%#+v", res)
}
