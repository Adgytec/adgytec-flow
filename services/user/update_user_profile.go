package user

import (
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type updateUserProfileData struct {
	Name           string
	ProfilePicture string
	About          string
	DateOfBirth    string
}

func (userProfile updateUserProfileData) Validate() error {
	validationErr := validation.ValidateStruct(&userProfile,
		validation.Field(&userProfile.Name, validation.Required, validation.Length(3, 100)),
		validation.Field(&userProfile.ProfilePicture, is.UUID),
		validation.Field(&userProfile.About, validation.Length(1, 1024)),
		validation.Field(&userProfile.DateOfBirth, validation.Required, validation.Date("2006-01-02")),
	)

	if validationErr != nil {
		return &core.FieldValidationError{
			ValidationErrors: validationErr,
		}
	}

	return nil
}

func (m *userServiceMux) updateUserProfileUtil(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	userProfieDetails, payloadErr := payload.DecodeRequestBodyAndValidate[updateUserProfileData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}
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
