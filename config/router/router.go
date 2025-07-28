package router

import (
	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/go-chi/chi/v5"
)

func CreateApplicationRouter(appConfig app.IApp) *chi.Mux {
	mux := chi.NewMux()
	return mux
}
