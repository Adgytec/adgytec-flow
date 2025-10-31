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

	// users
	mux.Post("/user", m.newUser)
	mux.Delete("/user/{userID}", m.removeUser)
	mux.Get("/users", m.getManagementUsers)
	mux.Get("/user/{userID}", m.getUserProfile)

	// user-group
	mux.Post("/user-group", m.newUserGroup)
	mux.Patch("/user-group/{groupID}", m.updateUserGroup)

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
