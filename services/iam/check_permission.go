package iam

import (
	"context"
	"errors"

	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

func (pc *iamServicePC) CheckPermission(ctx context.Context, permissionRequired PermissionProvider) error {
	return pc.CheckPermissions(ctx, []PermissionProvider{permissionRequired})
}

// CheckPermissions checks a list of permissions and succeeds if any one of them is granted.
// If the permissionsRequired slice is empty, it returns an error.
func (pc *iamServicePC) CheckPermissions(ctx context.Context, permissionsRequired []PermissionProvider) error {
	if len(permissionsRequired) == 0 {
		return &PermissionResolutionFailedError{
			Cause: ErrMissingPermissionsToCheck,
		}
	}

	actorDetails, actorDetailsErr := helpers.GetActorDetailsFromContext(ctx)
	if actorDetailsErr != nil {
		return actorDetailsErr
	}

	permissionEntity := permissionEntity{
		id:         actorDetails.ID,
		entityType: actorDetails.Type,
	}

	var lastPermissionErr error
	for _, permission := range permissionsRequired {
		lastPermissionErr = pc.service.checkPermission(ctx, permissionEntity, permission)
		if lastPermissionErr == nil {
			// permission granted
			return nil
		}

		if !errors.Is(lastPermissionErr, ErrPermissionDenied) {
			// some other error than permission denied so return early
			return lastPermissionErr
		}
	}

	return lastPermissionErr
}

func (s *iamService) checkPermission(ctx context.Context, permissionEntity permissionEntity, permissionRequired PermissionProvider) error {
	actorTypeError := s.validateActorType(
		permissionEntity.entityType,
		permissionRequired.GetPermissionActorType(),
	)
	if actorTypeError != nil {
		return actorTypeError
	}

	switch permissionRequired.GetPermissionType() {
	case PermissionTypeSelf:
		return s.resolveSelfPermission(permissionEntity, permissionRequired)
	case PermissionTypeApplication:
		return s.resolveApplicationPermission(ctx, permissionEntity, permissionRequired)
	case PermissionTypeManagement:
		return s.resolveManagementPermission(ctx, permissionEntity, permissionRequired)
	default:
		return &PermissionResolutionFailedError{
			Cause: ErrUnknownPermissionType,
		}
	}
}
