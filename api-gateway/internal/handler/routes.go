package handler 

import (
	"github.com/go-chi/chi"
)

func(h *Handler) Routes() *chi.Mux{
	r := chi.NewRouter()

	r.Post("/sign-up", h.userSignUp)
	r.Post("/sign-in", h.userSignIn)

	return r
}