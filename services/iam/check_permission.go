package iam

import (
	"context"
	"errors"

	"github.com/Adgytec/adgytec-flow/utils/actor"
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

	actorDetails, actorDetailsErr := actor.GetActorDetailsFromContext(ctx)
	if actorDetailsErr != nil {
		return actorDetailsErr
	}

	var lastPermissionErr error
	for _, permission := range permissionsRequired {
		lastPermissionErr = pc.service.checkPermission(ctx, actorDetails, permission)
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

func (s *iamService) checkPermission(ctx context.Context, actorDetails actor.ActorDetails, permissionRequired PermissionProvider) error {
	// check if current actor is system and permission can be resolved by system actor
	if actorDetails.IsSystem() {
		if permissionRequired.SystemAllowed() {
			// permission granted
			return nil
		}

		return &PermissionDeniedError{
			Reason: "system actor is not permitted to perform this action",
		}
	}

	actorTypeError := s.validateActorType(
		actorDetails.Type,
		permissionRequired.GetPermissionActorType(),
	)
	if actorTypeError != nil {
		return actorTypeError
	}

	switch permissionRequired.GetPermissionType() {
	case PermissionTypeSelf:
		return s.resolveSelfPermission(actorDetails, permissionRequired)
	case PermissionTypeApplication:
		return s.resolveApplicationPermission(ctx, actorDetails, permissionRequired)
	case PermissionTypeManagement:
		return s.resolveManagementPermission(ctx, actorDetails, permissionRequired)
	default:
		return &PermissionResolutionFailedError{
			Cause: ErrUnknownPermissionType,
		}
	}
}
