package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Log é um Middleware que intercepta a requisição.
// Ele recebe o próximo Handler na fila e retorna um Handler customizado.
func Log(next http.Handler) http.Handler {
	// http.HandlerFunc converte uma função comum na interface http.Handler.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()

		// ServeHTTP passa o controle para o próximo item (pode ser outro middleware ou a rota final).
		next.ServeHTTP(w, r)

		// Após a execução do handler final, o log é impresso (pós-processamento).
		fmt.Printf("Log: %s | %s | %v\n", r.Method, r.URL.Path, time.Since(begin))
	})
}

func Routing() {
	mux := http.NewServeMux()

	// 1. Path Variables: {id} captura uma parte da URL.
	// Disponível nativamente a partir do Go 1.22.
	mux.HandleFunc("/api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id") // Extrai o valor definido entre chaves.
		fmt.Printf("Buscando usuário: %s\n", id)
		fmt.Fprintf(w, "ID do Usuário: %s\n", id)
	})

	// 2. Restrição de Método: Só aceita POST.
	// Se tentar GET, o Go retorna 405 automaticamente.
	mux.HandleFunc("POST /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Sistema saudável!")
	})

	srv := &http.Server{
		Addr: ":8080",

		// Envolvemos nosso multiplexer (roteador) com o middleware de log.
		// O fluxo será: Request -> Log Middleware -> ServeMux -> Handler Final.
		Handler: Log(mux),

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	fmt.Println("Servidor iniciado em http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(fmt.Sprintf("Erro fatal: %v", err))
		}
	}
}
