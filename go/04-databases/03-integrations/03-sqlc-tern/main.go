package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	databaseURL := "postgres://postgres:root@localhost:5432/tests"

	db, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	queries := New(db)
	ctx := context.Background()

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(authors)

	author, err := queries.CreateAuthor(ctx, CreateAuthorParams{
		Name: "Miguel Castro",
		Bio:  pgtype.Text{String: "Platform Engineer", Valid: true}, // o tipo é pgtype.Text pq pode ser nil, ele recebe também um Valid que diz se o campo é nil ou não
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(author)
}
