package main

import (
	"log/slog"
	"net/http"
	"time"

	"03-bitly/api"
)

func main() {
	// A função run() encapsula a lógica principal para que possamos tratar erros de inicialização.
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		// os.Exit(1) -> util a depender do pipeline ou algo do tipo que opera com exitCodes
		// caso não definido, será 0
		return
	}
	slog.Info("all systems offline")
}

func run() error {
	// Map como banco de dados em memória.
	db := make(map[string]string)

	handler := api.NewHandler(db)

	// Customização do http.Server com Timeouts para evitar conexões "zumbis".
	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	// ListenAndServe é bloqueante. Ele só retorna erro se o servidor cair.
	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
