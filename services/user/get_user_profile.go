package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/models"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *userService) getUserProfile(ctx context.Context, currentUserId, userId string) (*models.GlobalUser, error) {
	permissionErr := s.accessManagement.CheckSelfPermission(currentUserId, userId, "get-user-profile")
	if permissionErr != nil {
		return nil, permissionErr
	}

	cachedUser, cacheOk := s.getUserCache.Get(userId)
	if cacheOk {
		return &cachedUser, nil
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
	s.getUserCache.Set(userId, userModel)

	return &userModel, nil
}

func (m *userServiceMux) getUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()
	userID, userIdOk := helpers.GetContextValue(reqCtx, helpers.ActorIDKey)
	if !userIdOk {
		payload.EncodeError(w, fmt.Errorf("Can't find current user."))
		return
	}

	user, userErr := m.service.getUserProfile(reqCtx, userID, userID)
	if userErr != nil {
		payload.EncodeError(w, userErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, user)

}
