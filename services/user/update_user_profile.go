package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type updateUserProfileData struct {
	Name           string
	ProfilePicture *string
	About          *string
	DateOfBirth    string
}

func (userProfile updateUserProfileData) Validate() error {
	validationErr := validation.ValidateStruct(&userProfile,
		validation.Field(&userProfile.Name, validation.Required, validation.Length(3, 100)),
		validation.Field(&userProfile.ProfilePicture, validation.NilOrNotEmpty, is.UUID),
		validation.Field(&userProfile.About, validation.NilOrNotEmpty, validation.Length(1, 1024)),
		validation.Field(&userProfile.DateOfBirth, validation.Required, validation.Date("2006-01-02")),
	)

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (userProfile updateUserProfileData) GetProfilePicture() (*uuid.UUID, error) {
	if userProfile.ProfilePicture == nil {
		return nil, nil
	}

	profilePictureUUID, parseErr := uuid.Parse(*userProfile.ProfilePicture)
	if parseErr != nil {
		return nil, core.ErrRequestBodyParsingFailed
	}

	return &profilePictureUUID, nil
}

func (userProfile updateUserProfileData) GetDateOfBirth() (pgtype.Date, error) {
	var zero pgtype.Date

	timeVal, parsingErr := time.Parse("2006-01-02", userProfile.DateOfBirth)
	if parsingErr != nil {
		return zero, core.ErrRequestBodyParsingFailed
	}

	return pgtype.Date{Time: timeVal, Valid: true}, nil
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

	// get parsed value
	profilePicture, profilePictureParsingErr := userProfile.GetProfilePicture()
	if profilePictureParsingErr != nil {
		return nil, profilePictureParsingErr
	}

	dob, dobParsingErr := userProfile.GetDateOfBirth()
	if dobParsingErr != nil {
		return nil, dobParsingErr
	}

	if userProfile.ProfilePicture != nil {
		// complete media upload
		mediaUploadErr := s.media.CompleteMediaItemUpload(ctx, *profilePicture)
		if mediaUploadErr != nil {
			return nil, mediaUploadErr
		}
	}

	// start transaction
	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return nil, txErr
	}
	defer tx.Rollback(context.Background())

	updatedUserProfileView, dbErr := qtx.UpdateGlobalUserProfile(ctx, db.UpdateGlobalUserProfileParams{
		ID:               userID,
		Name:             userProfile.Name,
		ProfilePictureID: profilePicture,
		About:            userProfile.About,
		DateOfBirth:      dob,
	})
	if dbErr != nil {
		if errors.Is(dbErr, pgx.ErrNoRows) {
			return nil, &UserNotFoundError{}
		}
		return nil, dbErr
	}
	txCommitErr := tx.Commit(context.Background())
	if txCommitErr != nil {
		// TODO: when media service is implemented will also clean complete media upload for failed profile update
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
