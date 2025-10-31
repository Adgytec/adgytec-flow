package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/google/uuid"
)

func (s *userManagementService) getUserGroupByID(ctx context.Context, groupID uuid.UUID) (*models.UserGroup, error) {
	return nil, nil
}

func (m *serviceMux) getUserGroupByID(w http.ResponseWriter, r *http.Request) {
	groupID, groupIDErr := reqparams.GetUserGroupIDFromRequest(r)
	if groupIDErr != nil {
		payload.EncodeError(w, groupIDErr)
		return
	}

	groupDetails, groupErr := m.service.getUserGroupByID(r.Context(), groupID)
	if groupErr != nil {
		payload.EncodeError(w, groupErr)
		return
	}

	payload.EncodeJSON(w, http.StatusOK, groupDetails)
}
