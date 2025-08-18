package core

import (
	"context"

	"github.com/google/uuid"
)

type IUserServicePC interface {
	CreateUser(context.Context, string) (uuid.UUID, error)
}
