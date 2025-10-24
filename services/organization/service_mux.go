package org

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/services"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type orgServiceMux struct {
	middleware core.MiddlewarePC
}

func (m *orgServiceMux) BasePath() string {
	return "/organization"
}

func (m *orgServiceMux) Router() *chi.Mux {
	mux := chi.NewMux()

	mux.Use(m.middleware.ValidateAndGetActorDetailsFromHttpRequest)
	mux.Use(m.middleware.ValidateActorTypeUserGlobalStatus)

	return mux
}

func NewOrgMux(params any) services.Mux {
	log.Info().
		Str("service", serviceName).
		Msg("new service mux")

	return &orgServiceMux{}
}
