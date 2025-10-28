package usermanagement

import (
	"context"
	"net/http"

	"github.com/Adgytec/adgytec-flow/services/iam"
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

func (m *serviceMux) removeUser(w http.ResponseWriter, r *http.Request) {}
