package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type updateUserSocialLinkData struct {
	ProfileLink string `json:"profileLink"`
}

func (socialLink updateUserSocialLinkData) Validate() error {
	validationErr := validation.ValidateStruct(
		&socialLink,
		validation.Field(&socialLink.ProfileLink, validation.Required, is.URL),
	)
	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

// admin shouldn't be able to update user social link
func (s *userService) updateUserSocialLink(ctx context.Context, userID, resourceID uuid.UUID, profileLink updateUserSocialLinkData) (*db.GlobalUserSocialLinks, error) {
	permissionErr := s.iam.CheckPermission(ctx, iam.NewPermissionRequiredFromSelfPermission(
		updateSelfProfilePermission,
		iam.PermissionRequiredResources{
			UserID: pointer.New(userID),
		},
	))
	if permissionErr != nil {
		return nil, permissionErr
	}

	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return nil, txErr
	}
	defer tx.Rollback(context.Background())

	updatedSocialLink, dbErr := qtx.Queries().UpdateUserSocialLink(ctx, db.UpdateUserSocialLinkParams{
		UserID:      userID,
		ID:          resourceID,
		ProfileLink: profileLink.ProfileLink,
	})
	if dbErr != nil {
		if errors.Is(dbErr, pgx.ErrNoRows) {
			return nil, &SocialLinkNotFoundError{}
		}
		return nil, dbErr
	}

	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		return nil, commitErr
	}

	// cache invalidate
	s.getUserCache.Delete(userID.String())
	return &updatedSocialLink, nil
}

func (m *userServiceMux) updateUserSelfSocialLink(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	userID, userIDErr := actor.GetActorIdFromContext(reqCtx)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	socialLinkID, socialLinkIDErr := m.service.getSocialLinkIDFromRequest(r)
	if socialLinkIDErr != nil {
		payload.EncodeError(w, socialLinkIDErr)
		return
	}

	updateSocialLinkDetails, payloadErr := payload.DecodeRequestBodyAndValidate[updateUserSocialLinkData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}

	updatedSocialLink, socialLinkErr := m.service.updateUserSocialLink(reqCtx, userID, socialLinkID, updateSocialLinkDetails)
	if socialLinkErr != nil {
		payload.EncodeError(w, socialLinkErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, updatedSocialLink)
}
