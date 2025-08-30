package access_management

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

func (pc *accessManagementPC) CheckPermission(ctx context.Context, permissionRequired []core.IPermissionRequired) error {
	actorDetails, actorDetailsErr := helpers.GetActorDetailsFromContext(ctx)
	if actorDetailsErr != nil {
		return actorDetailsErr
	}

	permissionEntity := core.PermissionEntity{
		ID:         actorDetails.ID,
		EntityType: actorDetails.Type,
	}

	// if permissionRequired is empty slice than MissingPermission is unknown
	var err error = &app_errors.PermissionDeniedError{
		MissingPermission: "Unknown",
	}

	for _, perm := range permissionRequired {
		err = pc.service.checkPermission(ctx, permissionEntity, perm)

		// permission successfully resolved when any of the permission is resolved successfully
		if err == nil {
			return nil
		}
	}

	return err
}

func (s *accessManagement) checkPermission(ctx context.Context, entity core.PermissionEntity, requiredPermission core.IPermissionRequired) error {
	// TODO: Implement actual permission checking logic here.
	// For now, returning an error to prevent accidental use in production.
	return app_errors.ErrNotImplemented
}
