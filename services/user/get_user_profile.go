package user

import (
	"context"
	"errors"

	"github.com/Adgytec/adgytec-flow/database/models"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *userService) getUserProfile(ctx context.Context, currentUserId, userId string) (*models.GlobalUser, error) {
	permissionErr := s.accessManagement.CheckSelfPermission(currentUserId, userId, "get-user-profile")
	if permissionErr != nil {
		return nil, permissionErr
	}

	userUUID, userIdErr := uuid.Parse(userId)
	if userIdErr != nil {
		return nil, &app_errors.InvalidUserIdError{
			InvalidUserId: userId,
		}
	}

	userProfile, dbErr := s.db.Queries().GetUserById(ctx, userUUID)
	if dbErr != nil {
		if errors.Is(dbErr, pgx.ErrNoRows) {
			return nil, &app_errors.UserNotFoundError{}
		}

		return nil, dbErr
	}

	userModel := s.getUserResponseModel(userProfile)
	return &userModel, nil
}
