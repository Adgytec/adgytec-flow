package user

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/services/media"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s *userService) newProfilePicture(ctx context.Context, userID uuid.UUID, profilePictureDetails media.NewMediaItemInput) (*media.NewMediaItemOutput, error) {
	return nil, nil
}

func (m *userServiceMux) newProfilePictureUtil(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	profilePictureDetails, reqBodyErr := payload.DecodeRequestBodyAndValidate[media.NewMediaItemInput](w, r)
	if reqBodyErr != nil {
		payload.EncodeError(w, reqBodyErr)
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
