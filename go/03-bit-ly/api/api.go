package api

import (
	"encoding/json"
	"log/slog"
	"math/rand/v2" // rand/v2 é mais rápido e seguro que o v1 (Go 1.22+)
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewHandler centraliza a criação do roteador e registro de middlewares.
func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shorten", handlePost(db))
	r.Get("/{code}", handleGet(db))

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

// Response padroniza a saída da API. omitEmpty remove campos nulos do JSON final.
type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

// sendJSON é um helper para evitar repetição de cabeçalhos e Marshalling em cada handler.
func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error while doing json marshall", "error", err)
		// Recursão controlada para erro interno caso o marshal falhe.
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
		// Decoding: Lê o JSON de entrada para a struct PostBody.
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		// url.Parse valida se o que foi enviado é tecnicamente uma URL válida.
		if _, err := url.Parse(body.URL); err != nil {
			sendJSON(w, Response{Error: "invalid url passed"}, http.StatusBadRequest)
			return // Faltava o return aqui para interromper a execução!
		}

		code := genCode()
		db[code] = body.URL // Inserção no mapa (Race condition possível sem Mutex)
		sendJSON(w, Response{Data: code}, http.StatusCreated)
	}
}

const characters = "abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYWZ123456789"

// genCode gera uma string aleatória de 8 caracteres.
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
		// chi.URLParam extrai o código da URL: /{code}
		code := chi.URLParam(r, "code")

		data, ok := db[code]
		if !ok {
			http.Error(w, "url not found", http.StatusNotFound)
			return
		}

		// StatusPermanentRedirect (308) diz ao navegador para sempre usar a URL de destino.
		http.Redirect(w, r, data, http.StatusPermanentRedirect)
	}
}
