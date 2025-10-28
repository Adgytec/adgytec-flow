package user

import (
	"context"
	"errors"

	"github.com/Adgytec/adgytec-flow/config/auth"
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

	// user is never delete from global users only their authentication detials from auth provider is removed
	dbErr := qtx.Queries().CreateGlobalUser(
		ctx,
		db.CreateGlobalUserParams{
			ID:    userID,
			Email: email,
		},
	)
	if dbErr != nil {
		return zero, dbErr
	}

	// always call auth provider method for new user
	// this ensures that previously removed user from auth provider is created again
	authErr := s.auth.NewUser(ctx, email)
	if authErr != nil {
		if !errors.Is(authErr, auth.ErrUserExists) {
			return zero, authErr
		}
	}

	txCommitErr := tx.Commit(context.Background())
	return userID, txCommitErr
}

func (pc *userServicePC) NewUser(ctx context.Context, email string) (uuid.UUID, error) {
	return pc.service.newUser(ctx, email)
}
