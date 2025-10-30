package user

import (
	"context"
	"errors"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *userService) getUserStatus(ctx context.Context, userID uuid.UUID) (db.GlobalUserStatus, error) {
	var zero db.GlobalUserStatus

	permissionErr := s.iam.CheckPermission(ctx,
		iam.NewPermissionRequiredFromSelfPermission(getSelfProfileStatusPermission,
			iam.PermissionRequiredResources{
				UserID: pointer.New(userID),
			},
		),
	)
	if permissionErr != nil {
		return zero, permissionErr
	}

	userStatus, statusErr := s.userStatusCache.Get(userID.String(), func() (db.GlobalUserStatus, error) {
		status, dbErr := s.db.Queries().GetUserStatus(ctx, userID)
		if dbErr != nil {
			if errors.Is(dbErr, pgx.ErrNoRows) {
				return zero, &UserNotFoundError{}
			}
			return zero, dbErr
		}

		return status, nil
	})

	return userStatus, statusErr
}

// GetUserStatus() is called always from middleware so user can only access their own status
func (pc *userServicePC) GetUserStatus(ctx context.Context, userID uuid.UUID) (db.GlobalUserStatus, error) {
	return pc.service.getUserStatus(ctx, userID)
}
