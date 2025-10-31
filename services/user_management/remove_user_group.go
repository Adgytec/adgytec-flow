package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/google/uuid"
)

func (s *userManagementService) removeUserGroup(ctx context.Context, groupID uuid.UUID) error {
	return nil
}

func (m *serviceMux) removeUserGroup(w http.ResponseWriter, r *http.Request) {
	groupID, groupIDErr := reqparams.GetUserGroupIDFromRequest(r)
	if groupIDErr != nil {
		payload.EncodeError(w, groupIDErr)
		return
	}

	removeErr := m.service.removeUserGroup(r.Context(), groupID)
	if removeErr != nil {
		payload.EncodeError(w, removeErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
