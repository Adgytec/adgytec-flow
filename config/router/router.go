package router

import (
	"github.com/go-chi/chi/v5"
)

func CreateApplicationRouter() *chi.Mux {
	mux := chi.NewMux()
	return mux
}
