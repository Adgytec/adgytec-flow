package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/media"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s *userService) newProfilePicture(ctx context.Context, userID uuid.UUID, profilePictureDetails media.NewMediaItemInput) (*media.NewMediaItemOutput, error) {
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

	profilePictureDetailsWithBucketPrefix := media.NewMediaItemInputWithBucketPrefix{
		NewMediaItemInput: profilePictureDetails,
		BucketPrefix: fmt.Sprintf(
			"/%s/profile",
			userID.String(),
		),
	}
	profilePictureDetailsOutput, mediaErr := s.media.NewMediaItem(ctx, profilePictureDetailsWithBucketPrefix)
	if mediaErr != nil {
		return nil, mediaErr
	}

	return profilePictureDetailsOutput, nil
}

func (m *userServiceMux) newProfilePictureUtil(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	profilePictureDetails, reqBodyErr := payload.DecodeRequestBodyAndValidate[media.NewMediaItemInput](w, r)
	if reqBodyErr != nil {
		payload.EncodeError(w, reqBodyErr)
		return
	}

	profilePictureErr := profilePictureDetails.EnsureMediaItemIsImage()
	if profilePictureErr != nil {
		payload.EncodeError(w, profilePictureErr)
		return
	}

	profilePictureDetailsOutput, uploadReqErr := m.service.newProfilePicture(r.Context(), userID, profilePictureDetails)
	if uploadReqErr != nil {
		payload.EncodeError(w, uploadReqErr)
		return
	}

	payload.EncodeJSON(w, http.StatusCreated, profilePictureDetailsOutput)
}

func (m *userServiceMux) newSelfProfilePicture(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	userID, userIDErr := actor.GetActorIdFromContext(reqCtx)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	m.newProfilePictureUtil(w, r, userID)
}

func (m *userServiceMux) newUserProfilePicture(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	userUUID, userIDErr := m.service.getUserUUIDFromString(userID)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	m.newProfilePictureUtil(w, r, userUUID)
}
