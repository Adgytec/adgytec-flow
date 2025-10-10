package user

import (
	"context"
	"errors"
	"net/http"
	"path"
	"unicode/utf8"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/media"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	"github.com/Adgytec/adgytec-flow/utils/types"
	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const (
	nameMinLength = 3
	nameMaxLength = 100

	aboutMinLength = 8
	aboutMaxLength = 1024
)

type updateUserProfileResponse struct {
	User                 *models.GlobalUser        `json:"user"`
	NextStep             string                    `json:"nextStep"`
	ProfileUploadDetails *media.MediaUploadDetails `json:"profileUploadDetails,omitempty"`
}

type updateUserProfileData struct {
	Name           types.NullableString                   `json:"name"`
	ProfilePicture types.Nullable[media.NewMediaItemInfo] `json:"profilePicture"`
	About          types.NullableString                   `json:"about"`
	DateOfBirth    types.Nullable[pgtype.Date]            `json:"dateOfBirth"`
}

func (userProfile updateUserProfileData) Validate() error {
	validationErr := validation.ValidateStruct(
		&userProfile,
		validation.Field(
			&userProfile.Name,
			validation.By(
				func(val any) error {
					name := val.(types.NullableString)
					if name.Null() {
						return nil
					}

					nameLen := utf8.RuneCountInString(name.Value)
					if nameLen < nameMinLength || nameLen > nameMaxLength {
						return ErrNameLength
					}

					return nil
				},
			),
		),
		validation.Field(
			&userProfile.About,
			validation.By(
				func(val any) error {
					about := val.(types.NullableString)
					if about.Null() {
						return nil
					}

					aboutLen := utf8.RuneCountInString(about.Value)
					if aboutLen < aboutMinLength || aboutLen > aboutMaxLength {
						return ErrAboutLength
					}

					return nil
				},
			),
		),
		validation.Field(
			&userProfile.DateOfBirth,
			validation.By(
				func(val any) error {
					dob := val.(types.Nullable[pgtype.Date])
					if dob.Null() {
						return nil
					}

					if !dob.Value.Valid {
						return ErrInvalidDateOfBirth
					}

					return nil
				},
			),
		),
		validation.Field(
			&userProfile.ProfilePicture,
			validation.By(
				func(val any) error {
					profilePictureInfo := val.(types.Nullable[media.NewMediaItemInfo])
					if profilePictureInfo.Null() {
						return nil
					}

					if err := profilePictureInfo.Value.Validate(); err != nil {
						return err
					}

					return nil
				},
			),
		),
	)

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (s *userService) updateUserProfile(ctx context.Context, userID uuid.UUID, userProfile updateUserProfileData) (*models.GlobalUser, *media.MediaUploadDetails, error) {
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
		return nil, nil, permissionErr
	}

	// get existing user detail
	existingUser, existingUserErr := s.getUserProfile(ctx, userID)
	if existingUserErr != nil {
		return nil, nil, existingUserErr
	}

	// start transaction
	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return nil, nil, txErr
	}
	defer tx.Rollback(context.Background())

	// update user obj
	updatedUser := db.UpdateGlobalUserProfileParams{
		ID: userID,
	}

	// name check
	if userProfile.Name.Missing() {
		updatedUser.Name = existingUser.Name
	} else if !userProfile.Name.Null() {
		updatedUser.Name = &userProfile.Name.Value
	}

	// about check
	if userProfile.About.Missing() {
		updatedUser.About = existingUser.About
	} else if !userProfile.About.Null() {
		updatedUser.About = &userProfile.About.Value
	}

	// dob check
	if userProfile.DateOfBirth.Missing() {
		updatedUser.DateOfBirth = existingUser.DateOfBirth
	} else if !userProfile.DateOfBirth.Null() {
		updatedUser.DateOfBirth = userProfile.DateOfBirth.Value
	}

	// profile picture check
	var profilePictureUploadDetails *media.MediaUploadDetails
	if userProfile.ProfilePicture.Missing() && existingUser.ProfilePicture != nil {
		updatedUser.ProfilePictureID = &existingUser.ProfilePicture.MediaID
	} else if !userProfile.ProfilePicture.Null() {
		// new profile picture
		updatedUser.ProfilePictureID = &userProfile.ProfilePicture.Value.ID

		// create new profile picture upload details
		mediaService := s.media.WithTransaction(qtx)
		uploadDetails, profilePictureUploadErr := mediaService.NewMediaItem(
			ctx,
			media.NewMediaItemInfoWithStorageDetails{
				NewMediaItemInfo: userProfile.ProfilePicture.Value,
				RequiredMime:     media.ImageMime,
				BucketPrefix:     path.Join(userID.String(), "profile"),
			},
		)
		if profilePictureUploadErr != nil {
			return nil, nil, profilePictureUploadErr
		}
		profilePictureUploadDetails = uploadDetails
	}

	updatedUserProfileView, dbErr := qtx.Queries().UpdateGlobalUserProfile(
		ctx,
		updatedUser,
	)
	if dbErr != nil {
		if errors.Is(dbErr, pgx.ErrNoRows) {
			return nil, nil, &UserNotFoundError{}
		}
		return nil, nil, dbErr
	}
	txCommitErr := tx.Commit(ctx)
	if txCommitErr != nil {
		return nil, nil, txCommitErr
	}

	updatedUserProfileModel := s.getUserResponseModel(updatedUserProfileView)
	s.getUserCache.Set(userID.String(), updatedUserProfileModel)

	return &updatedUserProfileModel, profilePictureUploadDetails, nil
}

func (m *userServiceMux) updateUserProfileUtil(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	userProfieDetails, payloadErr := payload.DecodeRequestBodyAndValidate[updateUserProfileData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}

	updatedUserProfile, profilePictureUpdateDetails, updateProfileErr := m.service.updateUserProfile(r.Context(), userID, userProfieDetails)
	if updateProfileErr != nil {
		payload.EncodeError(w, updateProfileErr)
		return
	}

	updateProfileNextStep := "Done"
	if profilePictureUpdateDetails != nil {
		updateProfileNextStep = "Upload profile picture"
	}

	updateProfileResponse := updateUserProfileResponse{
		User:                 updatedUserProfile,
		ProfileUploadDetails: profilePictureUpdateDetails,
		NextStep:             updateProfileNextStep,
	}

	payload.EncodeJSON(w, http.StatusOK, updateProfileResponse)
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
