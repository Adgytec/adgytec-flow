package user

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type userServicePC struct {
	service *userService
}

func (pc *userServicePC) CreateUser(ctx context.Context, username string) (string, error) {
	return "", nil
}

func (pc *userServicePC) UpdateLastAccessed(ctx context.Context, username string) error {
	return nil
}

func CreateUserServicePC(params iUserServiceParams) core.IUserServicePC {
	log.Println("creating user-service PC")
	return &userServicePC{
		service: createUserService(params),
	}
}
