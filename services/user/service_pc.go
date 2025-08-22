package user

import (
	"context"
	"log"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/google/uuid"
)

type userServicePC struct {
	service *userService
}

func (pc *userServicePC) GetUserStatus(ctx context.Context, userID uuid.UUID) (db_actions.GlobalUserStatus, error) {
	// TODO: will implement this later
	return db_actions.GlobalUserStatusDisabled, nil
}

func CreateUserServicePC(params iUserServiceParams) core.IUserServicePC {
	log.Println("creating user-service PC")
	return &userServicePC{
		service: createUserService(params),
	}
}
