package user

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/go-chi/chi/v5"
)

type userServiceMux struct {
	service    *userService
	middleware core.IMiddlewarePC
}

func (m *userServiceMux) BasePath() string {
	return "/user"
}

func (m *userServiceMux) Router() *chi.Mux {
	mux := chi.NewMux()

	mux.Group(func(router chi.Router) {
		router.Use(m.middleware.EnsureActorTypeUserOnly)

		router.Get("/profile", m.getUserProfileHandler)
	})

	mux.Group(func(router chi.Router) {
		router.Use(m.middleware.ActorManagementAccessCheck)

		router.Get("/all", m.getGlobalUsers)
		router.Patch("/{userID}/enable", m.enableGlobalUser)
		router.Patch("/{userID}/disable", m.disableGlobalUser)
	})

	return mux
}

func CreateUserServiceMux(params iUserServiceMuxParams) core.IServiceMux {
	log.Println("adding user-service mux")
	return &userServiceMux{
		service:    createUserService(params),
		middleware: params.Middleware(),
	}
}
