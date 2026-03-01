package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/miguelscastro/ignite/go/04-bitly-redis/repositories"
)

type shortenURLRequest struct {
	URL string `json:"url"`
}

type shortenURLResponse struct {
	Code string `json:"code"`
}

func handleShortenURL(repository repositories.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body shortenURLRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, APIResponse{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			sendJSON(w, APIResponse{Error: "invalid url passed"}, http.StatusBadRequest)
			return
		}

		code, err := repository.SaveShortenedURL(r.Context(), body.URL)
		if err != nil {
			slog.Error("failed to create code", "error", err)
			sendJSON(w, APIResponse{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}
		sendJSON(w, APIResponse{Data: code}, http.StatusCreated)
	}
}
