package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
)

// Password é um tipo customizado para permitir masking de logs
type Password string

func (p Password) String() string {
	return "[REDACTED]"
}

// LogValue precisa ter o V maiúsculo para satisfazer a interface slog.LogValuer
func (p Password) LogValue() slog.Value {
	return slog.StringValue("[REDACTED]")
}

type User struct {
	Username string
	ID       int64 `json:"id,string"`
	Role     string
	Password Password `json:"-"` // Usando o tipo customizado aqui
}

const LevelFoo = slog.Level(-50)

func main() {
	// 1. Configuração do Logger Customizado
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     LevelFoo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.String()
				// O slog retorna strings como "DEBUG-46" para níveis customizados
				if level == "DEBUG-46" {
					a.Value = slog.StringValue("FOO")
				}
			}
			return a
		},
	}

	// Inicializa o logger (JSON no Stdout)
	l := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(l)

	// Exemplo de integração com Uber Zap (se necessário no projeto)
	z, _ := zap.NewProduction()
	_ = slog.New(zapslog.NewHandler(z.Core(), nil))

	// 2. Teste do Masking (Redaction)
	p := Password("secret123")
	u := User{Username: "miguel", Password: "my-password", ID: 1}

	// Agora u.Password chamará automaticamente o LogValue()
	slog.Info("dados do usuario", "user_struct", u, "password_direto", p)

	// 3. Logger com contexto e grupos (Herdando atributos)
	l.With(slog.Group("app_info", slog.String("version", "1.0.0")))

	slog.LogAttrs(
		context.Background(),
		slog.LevelInfo, "request_info",
		slog.Group("http_data",
			slog.String("method", http.MethodDelete),
			slog.Int("status", http.StatusOK),
		),
		slog.Duration("latency", time.Second),
	)

	// 4. Configuração do Roteador Chi
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger) // Middleware padrão do Chi (usa o log do Go interno)

	db := map[int64]User{
		1: {Username: "admin", Password: "admin_password", Role: "admin", ID: 1},
	}

	r.Group(func(r chi.Router) {
		r.Use(JSONMiddleware) // Definido nos outros arquivos
		r.Get("/users/{id:[0-9]+}", handleGetUsers(db))
		r.Post("/users", handlePostUsers(db))
	})

	slog.Info("servidor rodando", "port", "8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error("falha ao iniciar servidor", "error", err)
		os.Exit(1)
	}
}
