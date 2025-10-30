package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/models"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/google/uuid"
)

func (s *userManagementService) getUserProfile(ctx context.Context, userID uuid.UUID) (*models.GlobalUser, error) {
	permissionErr := s.iam.CheckPermission(ctx,
		iam.NewPermissionRequiredFromManagementPermission(getManagementUserProfilePermission,
			iam.PermissionRequiredResources{},
		),
	)
	if permissionErr != nil {
		return nil, permissionErr
	}

	// check mangement user existence
	exists, dbErr := s.db.Queries().ManagementUserExists(ctx, userID)
	if dbErr != nil {
		return nil, dbErr
	}

	if !exists {
		return nil, &UserNotExistsInManagementError{}
	}

	// escale user privilage to system to get user profile from user service
	userDetails, userErr := s.getUserProfile(actor.NewSystemActorContext(ctx), userID)
	return userDetails, userErr
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
