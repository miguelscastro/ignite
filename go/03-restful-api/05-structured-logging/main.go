package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Inicializa um novo logger estruturado que escreve em formato JSON no Stdout.
	// O Handler é o componente que processa os registros de log (JSON, Texto, etc).
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Define o logger criado como o padrão para o serviço.
	// A partir daqui, chamadas como slog.Info usarão o formato JSON definido acima.
	slog.SetDefault(l)

	// Exemplo de log simples com pares chave-valor.
	slog.Info("Serviço sendo iniciado", "time", time.Now(), "version", "1.0.0")

	// Mistura de pares chave-valor com Attr tipados (slog.Int).
	slog.Info("incoming HTTP request", "method", http.MethodDelete, "time_taken", time.Second, slog.Int("status", http.StatusOK))

	// LogAttrs: Abordagem com tipagem rígida (Strongly Typed).
	// É mais eficiente (evita alocações extras) e garante que tipos inválidos gerem erro de compilação.
	slog.LogAttrs(
		context.Background(),
		slog.LevelInfo, "incoming HTTP request",
		slog.String("method", http.MethodDelete),
		slog.Duration("time_taken", time.Second),
		slog.String("user_agent", "hasuida"),
		slog.Int("status", http.StatusOK),
	)

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	db := map[int64]User{
		1: {Username: "admin", Password: "admin", Role: "admin", ID: 1},
	}

	r.Group(func(r chi.Router) {
		r.Use(JSONMiddleware)
		r.Get("/users/{id:[0-9]+}", handleGetUsers(db))
		r.Post("/users", handlePostUsers(db))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
