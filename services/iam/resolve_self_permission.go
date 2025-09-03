package iam

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
)

func (s *iam) resolveSelfPermission(
	permissionEntity core.PermissionEntity,
	permissionRequired core.PermissionProvider,
) error {
	// invalid case
	if permissionRequired.GetPermissionRequiredResources().UserID == nil {
		return &app_errors.PermissionResolutionFailedError{
			Cause: app_errors.ErrMissingRequiredResourcesValue,
		}
	}

	if permissionEntity.ID != *permissionRequired.GetPermissionRequiredResources().UserID {

		return &app_errors.PermissionDeniedError{
			Reason: "The resource is owned by a different user account than the one currently authenticated.",
		}
	}

	return nil
}
