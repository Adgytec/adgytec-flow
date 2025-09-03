package iam

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
)

func (s *iam) resolveSelfPermission(
	permissionEntity core.PermissionEntity,
	permissionRequired core.PermissionProvider,
) error {
	// invalid case
	if permissionRequired.GetPermissionRequiredResources().UserID == nil {
		return &PermissionResolutionFailedError{
			Cause: ErrMissingRequiredResourcesValue,
		}
	}

	if permissionEntity.ID != *permissionRequired.GetPermissionRequiredResources().UserID {

		return &PermissionDeniedError{
			Reason: "The resource is owned by a different user account than the one currently authenticated.",
		}
	}

	return nil
}
