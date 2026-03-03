package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router *chi.Mux
}

func (api *Api) handleCreateUser(w http.ResponseWriter, r *http.Request) {
}
