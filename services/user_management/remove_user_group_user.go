package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/utils/payload"
	reqparams "github.com/Adgytec/adgytec-flow/utils/req_params"
	"github.com/google/uuid"
)

func (s *userManagementService) removeUserGroupUser(ctx context.Context, groupID, userID uuid.UUID) error {
	permissionErr := s.iam.CheckPermission(ctx,
		iam.NewPermissionRequiredFromManagementPermission(
			removeUserFromUserGroupPermission,
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

	dbErr := qtx.Queries().RemoveUserGroupUser(ctx,
		db.RemoveUserGroupUserParams{
			UserGroupID: groupID,
			UserID:      userID,
		},
	)
	if dbErr != nil {
		return dbErr
	}

	return tx.Commit(ctx)

}

func (m *serviceMux) removeUserGroupUser(w http.ResponseWriter, r *http.Request) {
	userID, userIDErr := reqparams.GetUserIDFromRequest(r)
	if userIDErr != nil {
		payload.EncodeError(w, userIDErr)
		return
	}

	groupID, groupIDErr := reqparams.GetUserGroupIDFromRequest(r)
	if groupIDErr != nil {
		payload.EncodeError(w, groupIDErr)
		return
	}

	removeErr := m.service.removeUserGroupUser(r.Context(), groupID, userID)
	if removeErr != nil {
		payload.EncodeError(w, removeErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
