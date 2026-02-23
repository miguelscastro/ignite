package api

import (
	"encoding/json"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"net/url"

	"03-get-movies/movies"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer) // Recupera de pânicos para não derrubar o servidor
	r.Use(middleware.RequestID) // Atribui um ID único para cada requisição (bom para logs)
	r.Use(middleware.Logger)    // Loga as requisições no terminal

	r.Post("/api/shorten", handlePost(db))
	r.Get("/{code}", handleGet(db))

	r.Get("/", handleGetMovies())

	return r
}

func handleGetMovies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := movies.Search()
		if err != nil {
			// Se o Search falhar (ou der pânico), o Recoverer atua e caímos aqui
			sendJSON(w, Response{Error: "failed to fetch movies"}, http.StatusBadGateway)
			return
		}
		sendJSON(w, Response{Data: data}, http.StatusOK)
	}
}

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error while doing json marshall", "error", err)

		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("error while sending response", "error", err)
		return
	}
}

func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			sendJSON(w, Response{Error: "invalid url passed"}, http.StatusBadRequest)
			return
		}

		code := genCode()
		db[code] = body.URL
		sendJSON(w, Response{Data: code}, http.StatusCreated)
	}
}

const characters = "abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYWZ123456789"

func genCode() string {
	const n = 8
	byts := make([]byte, n)
	for i := range n {
		byts[i] = characters[rand.IntN(len(characters))]
	}
	return string(byts)
}

func handleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		data, ok := db[code]
		if !ok {
			http.Error(w, "url not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, data, http.StatusPermanentRedirect)
	}
}
