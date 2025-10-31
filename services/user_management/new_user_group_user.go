package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/google/uuid"
)

func (s *userManagementService) newUserGroupUser(ctx context.Context, groupID uuid.UUID, userData newUserData) (*uuid.UUID, error) {
	return nil, nil
}

func (m *serviceMux) newUserGroupUser(w http.ResponseWriter, r *http.Request) {
	groupID, groupIDErr := reqparams.GetUserGroupIDFromRequest(r)
	if groupIDErr != nil {
		payload.EncodeError(w, groupIDErr)
		return
	}

	newUserDetails, payloadErr := payload.DecodeRequestBodyAndValidate[newUserData](w, r)
	if payloadErr != nil {
		payload.EncodeError(w, payloadErr)
		return
	}

	userID, newUserErr := m.service.newUserGroupUser(r.Context(), groupID, newUserDetails)
	if newUserErr != nil {
		payload.EncodeError(w, newUserErr)
		return
	}

	payload.EncodeJSON(w, http.StatusCreated, map[string]any{
		"groupID": groupID,
		"userID":  userID,
		"message": "user added successfully",
	})
}
