package user

import (
	"context"
	"log"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
)

type PC interface {
	NewUser(ctx context.Context, email string) (uuid.UUID, error)
	GetUserStatus(ctx context.Context, userID uuid.UUID) (db.GlobalUserStatus, error)
}

type pc struct {
	service *userService
}

func NewPC(params userServiceParams) PC {
	log.Println("creating user-service PC")
	return &pc{
		service: newUserService(params),
	}
}
