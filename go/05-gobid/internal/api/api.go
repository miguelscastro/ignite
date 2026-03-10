package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/miguelscastro/ignite/go/05-gobid/internal/services"
)

type Api struct {
	Router      *chi.Mux
	UserService services.UserService
}
