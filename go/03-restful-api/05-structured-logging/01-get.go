package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type User struct {
	Username string

	ID   int64 `json:"id,string"`
	Role string

	Password string `json:"-"`
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func handleGetUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, _ := strconv.ParseInt(idStr, 10, 64)

		user, ok := db[id]
		if !ok {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}
