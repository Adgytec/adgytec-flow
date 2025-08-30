package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *userService) getUserProfile(ctx context.Context, userID uuid.UUID) (*models.GlobalUser, error) {
	requiredPermissions := []core.IPermissionRequired{
		helpers.CreatePermissionRequiredFromSelfPermission(
			getSelfProfilePermission,
			core.PermissionRequiredResources{
				UserID: helpers.ValuePtr(userID),
			},
		),
		helpers.CreatePermissionRequiredFromManagementPermission(
			getUserProfilePermission,
			core.PermissionRequiredResources{},
		),
	}

	permissionErr := s.accessManagement.CheckPermission(
		ctx,
		requiredPermissions,
	)
	if permissionErr != nil {
		return nil, permissionErr
	}

	cachedUser, cacheOk := s.getUserCache.Get(userID.String())
	if cacheOk {
		return &cachedUser, nil
	}

	userProfile, dbErr := s.db.Queries().GetUserById(ctx, userID)
	if dbErr != nil {
		if errors.Is(dbErr, pgx.ErrNoRows) {
			return nil, &app_errors.UserNotFoundError{}
		}

		return nil, dbErr
	}

	userModel := s.getUserResponseModel(userProfile)
	s.getUserCache.Set(userID.String(), userModel)

	return &userModel, nil
}

func (m *userServiceMux) getUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	userID, userIDErr := helpers.GetActorIdFromContext(reqCtx)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	user, userErr := m.service.getUserProfile(reqCtx, userID)
	if userErr != nil {
		payload.EncodeError(w, userErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, user)
}
