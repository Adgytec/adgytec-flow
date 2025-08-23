package user

import (
	"context"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/google/uuid"
)

func (pc *userServicePC) GetUserStatus(ctx context.Context, userID uuid.UUID) (db_actions.GlobalUserStatus, error) {
	// TODO: will implement this later
	return db_actions.GlobalUserStatusDisabled, nil
}
