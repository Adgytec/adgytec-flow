package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *userService) getUserProfile(ctx context.Context, userID uuid.UUID) (*models.GlobalUser, error) {
	requiredPermissions := []iam.PermissionProvider{
		iam.NewPermissionRequiredFromSelfPermission(
			getSelfProfilePermission,
			iam.PermissionRequiredResources{
				UserID: pointer.New(userID),
			},
		),
		iam.NewPermissionRequiredFromManagementPermission(
			getUserProfilePermission,
			iam.PermissionRequiredResources{},
		),
	}

	permissionErr := s.iam.CheckPermissions(
		ctx,
		requiredPermissions,
	)
	if permissionErr != nil {
		return nil, permissionErr
	}

	userModel, userError := s.getUserCache.Get(userID.String(), func() (models.GlobalUser, error) {
		var zero models.GlobalUser

		// get user profile
		userProfile, dbErr := s.db.Queries().GetUserById(ctx, userID)
		if dbErr != nil {
			if errors.Is(dbErr, pgx.ErrNoRows) {
				return zero, &UserNotFoundError{}
			}

			return zero, dbErr
		}

		// get user profile social links
		userSocialLinks, dbErr := s.db.Queries().GetUserSocialLinks(ctx, userID)
		if dbErr != nil {
			return zero, dbErr
		}

		userModel := s.getUserResponseModel(userProfile)
		userModel.SocialLinks = userSocialLinks

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

	userID, userIDErr := actor.GetActorIdFromContext(reqCtx)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	m.getUserProfileUtil(reqCtx, w, userID)
}

func (m *userServiceMux) getUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()
	userID := chi.URLParam(r, "userID")

	userUUID, userIDErr := m.service.getUserUUIDFromString(userID)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	m.getUserProfileUtil(reqCtx, w, userUUID)
}
