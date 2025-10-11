package user

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type UserServicePC interface {
	NewUser(ctx context.Context, email string) (uuid.UUID, error)
	GetUserStatus(ctx context.Context, userID uuid.UUID) (db.GlobalUserStatus, error)
}

type userServicePC struct {
	service *userService
}

func NewUserServicePC(params userServiceParams) UserServicePC {
	log.Info().
		Str("service", serviceName).
		Msg("new service pc")
	return &userServicePC{
		service: newUserService(params),
	}
}
