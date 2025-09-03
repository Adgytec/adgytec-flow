package user

import (
	"context"
	"log"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/google/uuid"
)

type UserServicePC interface {
	NewUser(context.Context, string) (uuid.UUID, error)
	GetUserStatus(context.Context, uuid.UUID) (db_actions.GlobalUserStatus, error)
}

type userServicePC struct {
	service *userService
}

func NewUserServicePC(params userServiceParams) UserServicePC {
	log.Println("creating user-service PC")
	return &userServicePC{
		service: newUserService(params),
	}
}
