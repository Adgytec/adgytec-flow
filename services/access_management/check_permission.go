package access_management

import (
	"context"
	"errors"

	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

// CheckPermission is called for actions which requires secure access
// only require single permissionRequired to be successfull to successfully resolve the permission
// if by any chance permissionRequired slice is empty than its an invalid case and will implicitly deny the permission with MissingPermission = 'unknown'
func (pc *accessManagementPC) CheckPermission(ctx context.Context, permissionRequired []core.IPermissionRequired) error {
	actorDetails, actorDetailsErr := helpers.GetActorDetailsFromContext(ctx)
	if actorDetailsErr != nil {
		return actorDetailsErr
	}

	permissionEntity := core.PermissionEntity{
		ID:         actorDetails.ID,
		EntityType: actorDetails.Type,
	}

	var err error = &app_errors.PermissionDeniedError{
		MissingPermission: "Unknown",
	}

	for _, perm := range permissionRequired {
		err = pc.service.checkPermission(ctx, permissionEntity, perm)
		if err == nil {
			// permission granted
			return nil
		}

		if !errors.Is(err, app_errors.ErrPermissionDenied) {
			// some other error than permission denied so return early
			return err
		}
	}

	return err
}

func (s *accessManagement) checkPermission(ctx context.Context, entity core.PermissionEntity, requiredPermission core.IPermissionRequired) error {
	// TODO: Implement actual permission checking logic here.
	// For now, returning an error to prevent accidental use in production.
	return app_errors.ErrNotImplemented
}
