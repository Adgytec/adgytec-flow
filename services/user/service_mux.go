package user

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/go-chi/chi/v5"
)

type mux struct {
	service    *userService
	middleware core.MiddlewarePC
}

func (m *mux) BasePath() string {
	return "/user"
}

func (m *mux) Router() *chi.Mux {
	mux := chi.NewMux()

	mux.Group(func(router chi.Router) {
		router.Use(m.middleware.EnsureActorTypeUserOnly)

		router.Get("/profile", m.getUserSelfProfileHandler)
	})

	mux.Group(func(router chi.Router) {
		router.Use(m.middleware.ActorManagementAccessCheck)

		router.Get("/list", m.getGlobalUsers)
		router.Get("/{userID}", m.getUserProfileHandler)
		router.Patch("/{userID}/enable", m.enableGlobalUser)
		router.Patch("/{userID}/disable", m.disableGlobalUser)
	})

	return mux
}

func NewMux(params muxParams) core.ServiceMux {
	log.Println("adding user-service mux")
	return &mux{
		service:    newService(params),
		middleware: params.Middleware(),
	}
}
