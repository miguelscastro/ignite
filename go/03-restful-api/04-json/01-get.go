package main

import (
	"encoding/json"
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
			// Retorno manual de erro em formato JSON para manter o contrato da API.
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"error":"user not found"}`))
			return
		}

		data, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(data)
	}
}
