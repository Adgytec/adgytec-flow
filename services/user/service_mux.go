package user

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/services"
	"github.com/go-chi/chi/v5"
)

type userServiceMux struct {
	service    *userService
	middleware core.MiddlewarePC
}

func (m *userServiceMux) BasePath() string {
	return "/user"
}

func (m *userServiceMux) Router() *chi.Mux {
	mux := chi.NewMux()

	mux.Group(func(router chi.Router) {
		router.Use(m.middleware.EnsureActorTypeUserOnly)

		router.Get("/profile", m.getUserSelfProfileHandler)

		router.Post("/profile/update", m.updateSelfProfile)
	})

	mux.Group(func(router chi.Router) {
		router.Use(m.middleware.ActorManagementAccessCheck)

		router.Get("/list", m.getGlobalUsers)
		router.Get("/{userID}", m.getUserProfileHandler)

		router.Patch("/{userID}/enable", m.enableGlobalUser)
		router.Patch("/{userID}/disable", m.disableGlobalUser)

		router.Post("/{userID}/update", m.updateUserProfile)
	})

	return mux
}

func NewUserServiceMux(params userServiceMuxParams) services.Mux {
	log.Printf("adding %s-service mux", serviceName)
	return &userServiceMux{
		service:    newUserService(params),
		middleware: params.Middleware(),
	}
}
