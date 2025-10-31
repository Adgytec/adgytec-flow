package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/google/uuid"
)

func (s *userManagementService) deleteUserGroup(ctx context.Context, groupID uuid.UUID) error {
	permissionErr := s.iam.CheckPermission(ctx,
		iam.NewPermissionRequiredFromManagementPermission(
			deleteUserGroupPermission,
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

	dbErr := qtx.Queries().DeleteUserGroup(ctx, groupID)
	if dbErr != nil {
		return dbErr
	}

	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		return commitErr
	}

	// invalidate cache
	s.userGroupCache.Delete(groupID.String())
	return nil
}

func (m *serviceMux) deleteUserGroup(w http.ResponseWriter, r *http.Request) {
	groupID, groupIDErr := reqparams.GetUserGroupIDFromRequest(r)
	if groupIDErr != nil {
		payload.EncodeError(w, groupIDErr)
		return
	}

	removeErr := m.service.deleteUserGroup(r.Context(), groupID)
	if removeErr != nil {
		payload.EncodeError(w, removeErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
