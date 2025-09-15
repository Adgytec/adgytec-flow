package user

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (m *userServiceMux) newProfilePictureUtil(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
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
