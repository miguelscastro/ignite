package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewMux()

	// --- MIDDLEWARES GLOBAIS ---
	r.Use(middleware.Recoverer) // Recupera de panics inesperados em handlers.
	r.Use(middleware.RequestID) // Adiciona ID único por requisição aos headers e logs.
	r.Use(middleware.Logger)    // Loga o tempo, método e status das chamadas.

	// Repositório em memória compartilhado entre os handlers.
	db := map[int64]User{
		1: {Username: "admin", Password: "admin", Role: "admin", ID: 1},
	}

	// Agrupamento de rotas com middleware específico para JSON.
	r.Group(func(r chi.Router) {
		r.Use(JSONMiddleware)

		r.Get("/users/{id:[0-9]+}", handleGetUsers(db))
		r.Post("/users", handlePostUsers(db))
	})

	fmt.Println("Servidor Chi rodando em http://localhost:8080")
	// ListenAndServe inicia o servidor; trava a execução do main.
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
