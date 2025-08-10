package user

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type userServicePC struct {
	service *userService
}

func (pc *userServicePC) CreateUser(username string) (string, error) {
	return "", nil
}

func (pc *userServicePC) UpdateLastAccessed(username string) error {
	return nil
}

func CreateUserServicePC(params iUserServiceParams) core.IUserServicePC {
	log.Println("creating user-service PC")
	return &userServicePC{
		service: &userService{
			db:               params.Database(),
			auth:             params.Auth(),
			accessManagement: params.AccessManagement(),
		},
	}
}
