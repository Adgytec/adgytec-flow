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

func (s *userService) getUserProfile(ctx context.Context, userID string) (*models.GlobalUser, error) {
	cachedUser, cacheOk := s.getUserCache.Get(userID)
	if cacheOk {
		return &cachedUser, nil
	}

	userUUID := uuid.MustParse(userID)
	userProfile, dbErr := s.db.Queries().GetUserById(ctx, userUUID)
	if dbErr != nil {
		if errors.Is(dbErr, pgx.ErrNoRows) {
			return nil, &app_errors.UserNotFoundError{}
		}

		return nil, dbErr
	}

	userModel := s.getUserResponseModel(userProfile)
	s.getUserCache.Set(userID, userModel)

	return &userModel, nil
}

func (m *userServiceMux) getUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()
	userID, userIDOk := helpers.GetContextValue(reqCtx, helpers.ActorIDKey)
	if !userIDOk {
		payload.EncodeError(w, fmt.Errorf("Can't find current user."))
		return
	}

	user, userErr := m.service.getUserProfile(reqCtx, userID)
	if userErr != nil {
		payload.EncodeError(w, userErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, user)
}
