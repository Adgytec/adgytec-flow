package user

import (
	"context"

	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
	"github.com/google/uuid"
)

func (s *userService) createUser(ctx context.Context, email string) (uuid.UUID, error) {
	var zero uuid.UUID
	userID := helpers.GetIDFromString(email)

	tx, txErr := s.db.NewTransaction(ctx)
	if txErr != nil {
		return zero, txErr
	}
	defer tx.Rollback(context.Background())
	qtx := s.db.Queries().WithTx(tx)

	inserted, dbErr := qtx.CreateGlobalUser(
		ctx,
		db_actions.CreateGlobalUserParams{
			ID:    userID,
			Email: email,
		},
	)
	if dbErr != nil {
		return zero, dbErr
	}

	// for newly inserted users also create the useraccount in auth service
	if inserted == 1 {
		authErr := s.auth.CreateUser(email)
		if authErr != nil {
			return zero, authErr
		}
	}

	txCommitErr := tx.Commit(context.Background())
	return userID, txCommitErr
}

func (pc *userServicePC) CreateUser(ctx context.Context, email string) (uuid.UUID, error) {
	return pc.service.createUser(ctx, email)
}
