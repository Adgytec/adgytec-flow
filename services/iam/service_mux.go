package iam

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/go-chi/chi/v5"
)

type mux struct {
	service    *iam
	middleware core.MiddlewarePC
}

func (m *mux) BasePath() string {
	return "/access-management"
}

func (m *mux) Router() *chi.Mux {
	mux := chi.NewMux()
	return mux
}

func NewMux(params muxParams) core.ServiceMux {
	log.Println("adding access-managment mux")
	return &mux{
		service:    newIAMService(params),
		middleware: params.Middleware(),
	}
}
