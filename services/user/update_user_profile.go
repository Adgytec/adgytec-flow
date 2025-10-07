package user

import (
	"net/http"
	"unicode/utf8"

	"github.com/Adgytec/adgytec-flow/services/media"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/Adgytec/adgytec-flow/utils/types"
	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type updateUserProfileData struct {
	Name           types.NullableString           `json:"name"`
	ProfilePicture media.NullableNewMediaItemInfo `json:"profilePicture"`
	About          types.NullableString           `json:"about"`
	DateOfBirth    types.Nullable[pgtype.Date]    `json:"dateOfBirth"`
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
					if nameLen < 3 && nameLen > 100 {
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
					if aboutLen < 8 && aboutLen > 1024 {
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
					profilePictureInfo := val.(media.NullableNewMediaItemInfo)
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

// func (s *userService) updateUserProfile(ctx context.Context, userID uuid.UUID, userProfile updateUserProfileData) (*models.GlobalUser, error) {
// 	requiredPermissions := []iam.PermissionProvider{
// 		iam.NewPermissionRequiredFromSelfPermission(
// 			updateSelfProfilePermission,
// 			iam.PermissionRequiredResources{
// 				UserID: pointer.New(userID),
// 			},
// 		),
// 		iam.NewPermissionRequiredFromManagementPermission(
// 			updateUserProfilePermission,
// 			iam.PermissionRequiredResources{},
// 		),
// 	}
//
// 	permissionErr := s.iam.CheckPermissions(ctx, requiredPermissions)
// 	if permissionErr != nil {
// 		return nil, permissionErr
// 	}
//
// 	// start transaction
// 	qtx, tx, txErr := s.db.WithTransaction(ctx)
// 	if txErr != nil {
// 		return nil, txErr
// 	}
// 	defer tx.Rollback(context.Background())
//
// 	updatedUserProfileView, dbErr := qtx.Queries().UpdateGlobalUserProfile(
// 		ctx,
// 		db.UpdateGlobalUserProfileParams{
// 			ID:               userID,
// 			Name:             userProfile.Name,
// 			ProfilePictureID: userProfile.ProfilePicture,
// 			About:            userProfile.About,
// 			DateOfBirth:      userProfile.DateOfBirth,
// 		},
// 	)
// 	if dbErr != nil {
// 		if errors.Is(dbErr, pgx.ErrNoRows) {
// 			return nil, &UserNotFoundError{}
// 		}
// 		return nil, dbErr
// 	}
// 	txCommitErr := tx.Commit(ctx)
// 	if txCommitErr != nil {
// 		return nil, txCommitErr
// 	}
//
// 	updatedUserProfileModel := s.getUserResponseModel(updatedUserProfileView)
// 	s.getUserCache.Set(userID.String(), updatedUserProfileModel)
//
// 	return &updatedUserProfileModel, nil
// }

func (m *userServiceMux) updateUserProfileUtil(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	// userProfieDetails, payloadErr := payload.DecodeRequestBodyAndValidate[updateUserProfileData](w, r)
	// if payloadErr != nil {
	// 	payload.EncodeError(w, payloadErr)
	// 	return
	// }

	// updatedUserProfile, updateProfileErr := m.service.updateUserProfile(r.Context(), userID, userProfieDetails)
	// if updateProfileErr != nil {
	// 	payload.EncodeError(w, updateProfileErr)
	// 	return
	// }
	//
	// payload.EncodeJSON(w, http.StatusOK, updatedUserProfile)
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
