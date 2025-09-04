package iam

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/go-chi/chi/v5"
)

type iamServiceMux struct {
	service    *iam
	middleware core.MiddlewarePC
}

func (m *iamServiceMux) BasePath() string {
	return "/iam"
}

func (m *iamServiceMux) Router() *chi.Mux {
	mux := chi.NewMux()
	return mux
}

func NewIAMMux(params iamServiceMuxParams) core.ServiceMux {
	log.Println("adding iam mux")
	return &iamServiceMux{
		service:    newIAMService(params),
		middleware: params.Middleware(),
	}
}
