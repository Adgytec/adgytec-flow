package user

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/go-chi/chi/v5"
)

type userServiceMux struct {
	service *userService
}

func (m *userServiceMux) BasePath() string {
	return "/user"
}

func (m *userServiceMux) Router() *chi.Mux {
	mux := chi.NewMux()
	return mux
}

func CreateUserServiceMux(params iUserServiceParams) core.IServiceMux {
	log.Println("adding user-service mux")
	return &userServiceMux{
		service: &userService{
			db:               params.Database(),
			auth:             params.Auth(),
			accessManagement: params.AccessManagement(),
		},
	}
}
