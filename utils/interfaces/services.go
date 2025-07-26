package interfaces

import "github.com/go-chi/chi/v5"

type IService interface {
	ServiceName() string
	BasePath() string
	Router() *chi.Mux
}

type IServiceInit interface{}

type IServicePC interface{}
