package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/google/uuid"
)

func (s *userManagementService) removeUser(ctx context.Context, userID uuid.UUID) error {
	permissionErr := s.iam.CheckPermission(ctx,
		iam.NewPermissionRequiredFromManagementPermission(
			removeManagementUserPermission,
			iam.PermissionRequiredResources{},
		),
	)
	if permissionErr != nil {
		return permissionErr
	}

	qtx, tx, txErr := s.db.WithTransaction(ctx)
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback(context.Background())

	dbErr := qtx.Queries().RemoveManagementUser(ctx, userID)
	if dbErr != nil {
		return dbErr
	}

	return tx.Commit(ctx)
}

func (m *serviceMux) removeUser(w http.ResponseWriter, r *http.Request) {
	userID, userIDErr := reqparams.GetUserIDFromRequest(r)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	removeUserErr := m.service.removeUser(r.Context(), userID)
	if removeUserErr != nil {
		payload.EncodeError(w, removeUserErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
