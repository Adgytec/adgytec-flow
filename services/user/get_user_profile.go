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
	"github.com/go-chi/chi/v5"
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

	userModel, userError := s.getUserCache.Get(userID.String(), func() (models.GlobalUser, error) {
		var zero models.GlobalUser
		userProfile, dbErr := s.db.Queries().GetUserById(ctx, userID)
		if dbErr != nil {
			if errors.Is(dbErr, pgx.ErrNoRows) {
				return zero, &app_errors.UserNotFoundError{}
			}

			return zero, dbErr
		}

		userModel := s.getUserResponseModel(userProfile)
		return userModel, nil
	})
	if userError != nil {
		return nil, userError
	}

	return &userModel, nil
}

func (m *userServiceMux) getUserProfileUtil(ctx context.Context, w http.ResponseWriter, userID uuid.UUID) {
	user, userErr := m.service.getUserProfile(ctx, userID)
	if userErr != nil {
		payload.EncodeError(w, userErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, user)
}

func (m *userServiceMux) getUserSelfProfileHandler(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	userID, userIDErr := helpers.GetActorIdFromContext(reqCtx)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	m.getUserProfileUtil(reqCtx, w, userID)
}

func (m *userServiceMux) getUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()
	userID := chi.URLParam(r, "userID")

	userUUID, userIdErr := m.service.getUserUUIDFromString(userID)
	if userIdErr != nil {
		payload.EncodeError(w, userIdErr)
		return
	}

	m.getUserProfileUtil(reqCtx, w, userUUID)
}
