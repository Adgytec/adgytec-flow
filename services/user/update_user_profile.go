package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type updateUserProfileData struct {
	Name           string
	ProfilePicture *uuid.UUID
	About          *string
	DateOfBirth    pgtype.Date
}

func (userProfile updateUserProfileData) Validate() error {
	validationErr := validation.ValidateStruct(&userProfile,
		validation.Field(&userProfile.Name, validation.Required, validation.Length(3, 100)),
		validation.Field(&userProfile.About, validation.NilOrNotEmpty, validation.Length(1, 1024)),
		validation.Field(&userProfile.DateOfBirth, validation.Required),
	)

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (s *userService) updateUserProfile(ctx context.Context, userID uuid.UUID, userProfile updateUserProfileData) (*models.GlobalUser, error) {
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
		return nil, permissionErr
	}

	// start transaction
	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return nil, txErr
	}
	defer tx.Rollback(context.Background())

	if userProfile.ProfilePicture != nil {
		// complete media upload
		mediaUploadErr := s.media.WithTransaction(qtx).CompleteMediaItemUpload(ctx, *userProfile.ProfilePicture)
		if mediaUploadErr != nil {
			return nil, mediaUploadErr
		}
	}

	updatedUserProfileView, dbErr := qtx.Queries().UpdateGlobalUserProfile(
		ctx,
		db.UpdateGlobalUserProfileParams{
			ID:               userID,
			Name:             userProfile.Name,
			ProfilePictureID: userProfile.ProfilePicture,
			About:            userProfile.About,
			DateOfBirth:      userProfile.DateOfBirth,
		},
	)
	if dbErr != nil {
		if errors.Is(dbErr, pgx.ErrNoRows) {
			return nil, &UserNotFoundError{}
		}
		return nil, dbErr
	}
	txCommitErr := tx.Commit(context.Background())
	if txCommitErr != nil {
		return nil, txCommitErr
	}

	updatedUserProfileModel := s.getUserResponseModel(updatedUserProfileView)
	s.getUserCache.Set(userID.String(), updatedUserProfileModel)

	return &updatedUserProfileModel, nil
}

func (m *userServiceMux) updateUserProfileUtil(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	userProfieDetails, payloadErr := payload.DecodeRequestBodyAndValidate[updateUserProfileData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}

	updatedUserProfile, updateProfileErr := m.service.updateUserProfile(r.Context(), userID, userProfieDetails)
	if updateProfileErr != nil {
		payload.EncodeError(w, updateProfileErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, updatedUserProfile)
}

func (m *userServiceMux) updateSelfProfile(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	userID, userIDErr := actor.GetActorIdFromContext(reqCtx)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	m.updateUserProfileUtil(w, r, userID)
}

func (m *userServiceMux) updateUserProfile(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	userUUID, userIDErr := m.service.getUserUUIDFromString(userID)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	m.updateUserProfileUtil(w, r, userUUID)
}
