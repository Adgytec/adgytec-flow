package access_management

import (
	"context"
	"errors"

	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

// CheckPermission is called for actions which requires secure access
func (pc *accessManagementPC) CheckPermission(ctx context.Context, permissionRequired core.IPermissionRequired) error {
	return pc.CheckPermissions(ctx, []core.IPermissionRequired{permissionRequired})
}

// CheckPermissions only require single permissionRequired to be successfull to successfully resolve the permission
// if by any chance permissionRequired slice is empty than its an invalid case and will implicitly deny the permission with MissingPermission = 'unknown'
func (pc *accessManagementPC) CheckPermissions(ctx context.Context, permissionsRequired []core.IPermissionRequired) error {
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

	for _, permissionRequired := range permissionsRequired {
		err = pc.service.checkPermission(ctx, permissionEntity, permissionRequired)
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

func (s *accessManagement) checkPermission(ctx context.Context, permissionEntity core.PermissionEntity, permissionRequired core.IPermissionRequired) error {
	actorTypeError := s.validateActorType(
		permissionEntity.EntityType,
		permissionRequired.GetPermissionActorType(),
	)
	if actorTypeError != nil {
		return actorTypeError
	}

	switch permissionRequired.GetPermissionType() {
	case core.PermissionTypeSelf:
		return s.resolveSelfPermission(permissionEntity, permissionRequired)
	case core.PermissionTypeApplication:
		return s.resolveApplicationPermission(ctx, permissionEntity, permissionRequired)
	case core.PermissionTypeManagement:
		return s.resolveManagementPermission(ctx, permissionEntity, permissionRequired)
	default:
		return &app_errors.PermissionResolutionFailedError{
			Cause: app_errors.ErrUnknownPermissionType,
		}
	}
}
