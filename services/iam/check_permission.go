package iam

import (
	"context"
	"errors"

	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

func (pc *pc) CheckPermission(ctx context.Context, permissionRequired core.PermissionProvider) error {
	return pc.CheckPermissions(ctx, []core.PermissionProvider{permissionRequired})
}

// CheckPermissions checks a list of permissions and succeeds if any one of them is granted.
// If the permissionsRequired slice is empty, it returns an error.
func (pc *pc) CheckPermissions(ctx context.Context, permissionsRequired []core.PermissionProvider) error {
	if len(permissionsRequired) == 0 {
		return &app_errors.PermissionResolutionFailedError{
			Cause: app_errors.ErrMissingPermissionsToCheck,
		}
	}

	actorDetails, actorDetailsErr := helpers.GetActorDetailsFromContext(ctx)
	if actorDetailsErr != nil {
		return actorDetailsErr
	}

	permissionEntity := core.PermissionEntity{
		ID:         actorDetails.ID,
		EntityType: actorDetails.Type,
	}

	var lastPermissionErr error
	for _, permission := range permissionsRequired {
		lastPermissionErr = pc.service.checkPermission(ctx, permissionEntity, permission)
		if lastPermissionErr == nil {
			// permission granted
			return nil
		}

		if !errors.Is(lastPermissionErr, app_errors.ErrPermissionDenied) {
			// some other error than permission denied so return early
			return lastPermissionErr
		}
	}

	return lastPermissionErr
}

func (s *iam) checkPermission(ctx context.Context, permissionEntity core.PermissionEntity, permissionRequired core.PermissionProvider) error {
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
