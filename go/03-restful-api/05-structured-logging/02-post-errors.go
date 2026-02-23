package main

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
)

func handlePostUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1000)

		data, err := io.ReadAll(r.Body)
		if err != nil {
			var maxErr *http.MaxBytesError

			if errors.As(err, &maxErr) {
				sendJSON(w, Response{Error: "body too large"}, http.StatusRequestEntityTooLarge)
				return
			}
			slog.Error("error during json read", "error", err)
			sendJSON(
				w,
				Response{Error: "something went wrong"},
				http.StatusInternalServerError,
			)
			return
		}

		var user User

		if err := json.Unmarshal(data, &user); err != nil {
			sendJSON(
				w,
				Response{Error: "invalid body"},
				http.StatusUnprocessableEntity,
			)
			return
		}

		db[user.ID] = user
		w.WriteHeader(http.StatusCreated)
	}
}
