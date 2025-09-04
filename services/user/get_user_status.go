package user

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/google/uuid"
)

func (pc *userServicePC) GetUserStatus(ctx context.Context, userID uuid.UUID) (db.GlobalUserStatus, error) {
	// TODO: will implement this later
	return db.GlobalUserStatusDisabled, nil
}
