package access_management

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

func (pc *accessManagementPC) CheckPermission(ctx context.Context, permissionRequired core.PermissionRequired) error {
	actorDetails, actorDetailsErr := helpers.GetActorDetailsFromContext(ctx)
	if actorDetailsErr != nil {
		return actorDetailsErr
	}

	permissionEntity := core.PermissionEntity{
		ID:         actorDetails.ID,
		EntityType: actorDetails.Type,
	}
	return pc.service.checkPermission(ctx, permissionEntity, permissionRequired)
}

func (s *accessManagement) checkPermission(ctx context.Context, entity core.PermissionEntity, requiredPermission core.PermissionRequired) error {
	// TODO: Implement actual permission checking logic here.
	// For now, returning an error to prevent accidental use in production.
	return app_errors.ErrNotImplemented
}
