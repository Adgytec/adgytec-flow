package user

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s *userService) removeUserSocialLink(ctx context.Context, userID, resourceID uuid.UUID) error {
	requiredPermissions := []iam.PermissionProvider{
		iam.NewPermissionRequiredFromSelfPermission(
			updateSelfProfilePermission,
			iam.PermissionRequiredResources{
				UserID: pointer.New(userID),
			},
		),
		iam.NewPermissionRequiredFromManagementPermission(
			updateUserProfilePermission,
			iam.PermissionRequiredResources{},
		),
	}

	permissionErr := s.iam.CheckPermissions(ctx, requiredPermissions)
	if permissionErr != nil {
		return permissionErr
	}

	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback(context.Background())

	rowsAffected, dbErr := qtx.Queries().RemoveUserSocialLink(ctx, db.RemoveUserSocialLinkParams{
		UserID: userID,
		ID:     resourceID,
	})
	if dbErr != nil {
		return dbErr
	}

	if rowsAffected == 0 {
		return &SocialLinkNotFoundError{}
	}

	return tx.Commit(ctx)
}

func (m *userServiceMux) removeUserSocialLinkUtil(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	socialLinkID, socialLinkIDErr := m.service.getSocialLinkIDFromRequest(r)
	if socialLinkIDErr != nil {
		payload.EncodeError(w, socialLinkIDErr)
		return
	}

	removeSocialLinkErr := m.service.removeUserSocialLink(r.Context(), userID, socialLinkID)
	if removeSocialLinkErr != nil {
		payload.EncodeError(w, removeSocialLinkErr)
		return
	}
}

func (m *userServiceMux) removeUserSelfSocialLink(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	userID, userIDErr := actor.GetActorIdFromContext(reqCtx)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	m.removeUserSocialLinkUtil(w, r, userID)
}

func (m *userServiceMux) removeUserSocialLink(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	userUUID, userIDErr := m.service.getUserUUIDFromString(userID)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	m.removeUserSocialLinkUtil(w, r, userUUID)
}
