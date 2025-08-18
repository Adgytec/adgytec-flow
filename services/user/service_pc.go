package user

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type userServicePC struct {
	service *userService
}

func CreateUserServicePC(params iUserServiceParams) core.IUserServicePC {
	log.Println("creating user-service PC")
	return &userServicePC{
		service: createUserService(params),
	}
}
