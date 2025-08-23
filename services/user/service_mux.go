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

	// TODO: add middleware to ensure actor type is user
	mux.Get("/profile", m.getUserProfileHandler)
	mux.Get("/all", m.getGlobalUsers)
	mux.Patch("/{userID}/enable", m.enableGlobalUser)
	mux.Patch("/{userID}/disable", m.disableGlobalUser)

	return mux
}

func CreateUserServiceMux(params iUserServiceMuxParams) core.IServiceMux {
	log.Println("adding user-service mux")
	return &userServiceMux{
		service:    createUserService(params),
		middleware: params.Middleware(),
	}
}
