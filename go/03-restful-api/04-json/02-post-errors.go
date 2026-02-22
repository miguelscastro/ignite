package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// handlePostUsers implementa a lógica de criação com limite de bytes (segurança de memória).
func handlePostUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// r.Body é um stream; http.MaxBytesReader limita a leitura para evitar estouro de RAM.
		r.Body = http.MaxBytesReader(w, r.Body, 1000)

		// data contém o payload bruto lido do stream.
		data, err := io.ReadAll(r.Body)
		if err != nil {
			var maxErr *http.MaxBytesError
			// Avalia se o erro foi causado pelo estouro do limite de bytes definido.
			if errors.As(err, &maxErr) {
				http.Error(w, "body too large", http.StatusRequestEntityTooLarge)
				return
			}
			fmt.Println(err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		var user User
		// Unmarshal converte os bytes em memória para a struct de destino.
		if err := json.Unmarshal(data, &user); err != nil {
			http.Error(w, "invalid body", http.StatusUnprocessableEntity)
			return
		}

		// Risco de Race Condition: Maps não são seguros para escrita concorrente sem Mutex.
		db[user.ID] = user
		w.WriteHeader(http.StatusCreated)
	}
}
