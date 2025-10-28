package usermanagement

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/services"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type serviceMux struct {
	service    *userManagementService
	middleware core.MiddlewarePC
}

func (m *serviceMux) BasePath() string {
	return "/user-management"
}

func (m *serviceMux) Router() *chi.Mux {
	mux := chi.NewMux()

	mux.Use(m.middleware.ValidateAndGetActorDetailsFromHttpRequest)
	mux.Use(m.middleware.ValidateActorTypeUserGlobalStatus)

	mux.Post("/user", m.newUser)
	mux.Delete("/user/{userID}", m.removeUser)

	return mux
}

func NewMux(params serviceMuxParams) services.Mux {
	log.Info().
		Str("service", serviceName).
		Msg("new service mux")
	return &serviceMux{
		service:    newService(params),
		middleware: params.Middleware(),
	}
}
