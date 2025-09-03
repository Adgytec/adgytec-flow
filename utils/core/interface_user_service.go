package core

import (
	"context"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/google/uuid"
)

type UserServicePC interface {
	NewUser(context.Context, string) (uuid.UUID, error)
	GetUserStatus(context.Context, uuid.UUID) (db_actions.GlobalUserStatus, error)
}
