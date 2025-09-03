package user

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type userServicePC struct {
	service *userService
}

func NewUserServicePC(params userServiceParams) core.UserServicePC {
	log.Println("creating user-service PC")
	return &userServicePC{
		service: newUserService(params),
	}
}
