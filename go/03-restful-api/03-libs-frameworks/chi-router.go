package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ChiRouter() {
	r := chi.NewMux()

	// --- MIDDLEWARES GLOBAIS ---
	r.Use(middleware.Recoverer) // Recupera de pânicos (evita crash do server)
	r.Use(middleware.RequestID) // ID único por request para rastreabilidade
	r.Use(middleware.Logger)    // Log de acesso no terminal

	r.Get("/horario", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Fprintln(w, now.Format(time.RFC3339))
	})

	// --- ORGANIZAÇÃO POR ESCOPO (Route) ---
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Lista de usuários v1"))
			})

			// Uso de REGEX: id deve ser apenas números.
			r.Get("/users/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
				// No Chi, pegamos o parâmetro com URLParam
				id := chi.URLParam(r, "id")
				fmt.Fprintf(w, "Buscando usuário v1 com ID: %s", id)
			})
		})

		r.Route("/v2", func(r chi.Router) {
			// Futuras implementações da v2 aqui
		})

		// .With(): Aplica um middleware apenas para uma linha/rota específica
		r.With(middleware.RealIP).Get("/client-ip", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("IP rastreado para auditoria"))
		})

		// --- GRUPOS PROTEGIDOS (Group) ---
		// Útil para aplicar middlewares (como Auth) em várias rotas sem mudar o path
		r.Group(func(r chi.Router) {
			r.Use(middleware.BasicAuth("meu-app", map[string]string{
				"admin": "123456",
			}))

			r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "Banco de dados e sistemas: OK")
			})
		})
	})

	fmt.Println("Servidor Chi rodando em http://localhost:8080")

	// Em produção, substitua ListenAndServe por &http.Server com timeouts!
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
