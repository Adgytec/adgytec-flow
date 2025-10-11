package iam

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/services"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type iamServiceMux struct {
	service    *iamService
	middleware core.MiddlewarePC
}

func (m *iamServiceMux) BasePath() string {
	return "/iam"
}

func (m *iamServiceMux) Router() *chi.Mux {
	mux := chi.NewMux()

	mux.Use(m.middleware.ValidateAndGetActorDetailsFromHttpRequest)
	mux.Use(m.middleware.ValidateActorTypeUserGlobalStatus)

	return mux
}

func NewIAMMux(params iamServiceMuxParams) services.Mux {
	log.Info().
		Str("service", serviceName).
		Msg("new service mux")
	return &iamServiceMux{
		service:    newIAMService(params),
		middleware: params.Middleware(),
	}
}
