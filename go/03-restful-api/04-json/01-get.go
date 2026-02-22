package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// User define a estrutura de dados com metadados para o encoder JSON.
type User struct {
	Username string
	// json:",string" força a conversão do int64 para string no JSON final.
	ID   int64 `json:"id,string"`
	Role string
	// json:"-" oculta o campo no JSON por segurança.
	Password string `json:"-"`
}

// JSONMiddleware define o cabeçalho de resposta como JSON para o grupo de rotas.
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// handleGetUsers utiliza Closure para acessar o mapa db injetado na inicialização.
func handleGetUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		// ParseInt converte o parâmetro da URL para o tipo da chave do mapa.
		id, _ := strconv.ParseInt(idStr, 10, 64)

		user, ok := db[id]
		if !ok {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}
