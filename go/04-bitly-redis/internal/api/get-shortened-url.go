package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/miguelscastro/ignite/go/04-bitly-redis/repositories"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type getShortenedURLResponse struct {
	FullURL string `json:"full_url"`
}

func handleGetShortenedURL(repository repositories.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		fullURL, err := repository.GetFullURL(r.Context(), code)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				sendJSON(w, APIResponse{Error: "code not found"}, http.StatusNotFound)
				return
			}
			slog.Error("failed to retrieve code", "error", err)
			sendJSON(w, APIResponse{Error: "something went wrong"}, http.StatusInternalServerError)
		}
		sendJSON(w, APIResponse{Data: getShortenedURLResponse{FullURL: fullURL}}, http.StatusOK)
	}
}
