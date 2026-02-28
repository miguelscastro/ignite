package main

import (
	"log/slog"
	"net/http"
	"time"

	"04-bit-ly-redis/internal/api"
	"04-bit-ly-redis/repositories"

	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)

		return
	}
	slog.Info("all systems offline")
}

func run() error {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	repository := repositories.NewRepository(redisDB)
	handler := api.NewHandler(repository)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
