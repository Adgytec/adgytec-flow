package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/google/uuid"
)

func (s *userManagementService) getUserProfile(ctx context.Context, userID uuid.UUID) (*models.GlobalUser, error) {
	return nil, nil
}

func (m *serviceMux) getUserProfile(w http.ResponseWriter, r *http.Request) {
	userID, userIDErr := reqparams.GetUserIDFromRequest(r)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	userProfile, userErr := m.service.getUserProfile(r.Context(), userID)
	if userErr != nil {
		payload.EncodeError(w, userErr)
		return
	}
	payload.EncodeJSON(w, http.StatusOK, userProfile)
}
