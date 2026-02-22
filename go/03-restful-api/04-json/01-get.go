package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// User define a estrutura de dados com metadados para o encoder JSON.
type User struct {
	Username string
	// json:",string" faz com que o campo seja tratado como string no JSON final.
	// Útil para garantir compatibilidade com clientes que não suportam números de 64 bits.
	ID   int64 `json:"id,string"`
	Role string
	// json:"-" instrui o pacote encoding/json a ignorar este campo completamente.
	// Segurança: impede que dados sensíveis (senhas) saiam do servidor.
	Password string `json:"-"`
}

func Get() {
	r := chi.NewMux()

	// --- MIDDLEWARES GLOBAIS ---
	// Recoverer: Captura panics e evita que o servidor encerre abruptamente.
	r.Use(middleware.Recoverer)
	// RequestID: Injeta um ID único por requisição para rastreamento em logs.
	r.Use(middleware.RequestID)
	// Logger: Registra no console o tempo de execução, método e rota.
	r.Use(middleware.Logger)

	// db atua como um repositório em memória para persistência temporária.
	db := map[int64]User{
		1: {Username: "admin", Password: "admin", Role: "admin", ID: 1},
	}

	// Group permite isolar middlewares para conjuntos específicos de rotas.
	r.Group(func(r chi.Router) {
		r.Use(JSONMiddleware)
		// Aqui handleGetUsers é invocado para retornar a função HandlerFunc final.
		r.Get("/users/{id:[0-9]+}", handleGetUsers(db))
		r.Post("/users", handlePostUsers)
	})

	fmt.Println("Servidor Chi rodando em http://localhost:8080")
	// ListenAndServe inicia o loop de escuta TCP de forma bloqueante.
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

// JSONMiddleware define o cabeçalho de resposta como JSON antes de passar para o próximo handler.
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// handleGetUsers utiliza o conceito de Closure para injetar o mapa db no escopo do handler.
func handleGetUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		// Converte a string capturada na URL para o tipo da chave do mapa (int64).
		id, _ := strconv.ParseInt(idStr, 10, 64)

		user, ok := db[id]

		if ok {
			// json.Marshal transforma a struct em um slice de bytes ([]byte).
			data, err := json.Marshal(user)
			if err != nil {
				panic(err)
			}

			// Escreve os bytes do JSON no corpo da resposta HTTP.
			_, _ = w.Write(data)
		}
		// Nota: se !ok, o servidor não retorna body (404 implícito ou 200 vazio).
	}
}

// handlePostUsers (assinatura pronta para implementação futura).
func handlePostUsers(w http.ResponseWriter, r *http.Request) {}
