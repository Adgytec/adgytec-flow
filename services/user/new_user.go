package user

import (
	"context"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/google/uuid"
)

func (s *userService) newUser(ctx context.Context, email string) (uuid.UUID, error) {
	var zero uuid.UUID
	userID := core.GetIDFromPayload([]byte(email))

	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return zero, txErr
	}
	defer tx.Rollback(context.Background())

	inserted, dbErr := qtx.CreateGlobalUser(
		ctx,
		db.CreateGlobalUserParams{
			ID:    userID,
			Email: email,
		},
	)
	if dbErr != nil {
		return zero, dbErr
	}

	// for newly inserted users also create the useraccount in auth service
	if inserted == 1 {
		authErr := s.auth.NewUser(email)
		if authErr != nil {
			return zero, authErr
		}
	}

	txCommitErr := tx.Commit(context.Background())
	return userID, txCommitErr
}

func (pc *userServicePC) NewUser(ctx context.Context, email string) (uuid.UUID, error) {
	return pc.service.newUser(ctx, email)
}
