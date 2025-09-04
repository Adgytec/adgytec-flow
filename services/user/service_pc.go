package user

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
)

type UserServicePC interface {
	NewUser(ctx context.Context, email string) (uuid.UUID, error)
	GetUserStatus(ctx context.Context, userID uuid.UUID) (db.GlobalUserStatus, error)
}

type userServicePC struct {
	service *userService
}

func NewUserServicePC(params userServiceParams) UserServicePC {
	log.Printf("creating %s-service PC", serviceName)
	return &userServicePC{
		service: newUserService(params),
	}
}
