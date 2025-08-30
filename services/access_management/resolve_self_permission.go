package access_management

import (
	"errors"

	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
)

func (s *accessManagement) resolveSelfPermission(
	permissionEntity core.PermissionEntity,
	permissionRequired core.IPermissionRequired,
) error {
	// invalid case
	if permissionRequired.GetPermissionRequiredResources().UserID == nil {
		return &app_errors.PermissionResolutionFailedError{
			Cause: errors.New("Missing required resources userID value for self permission resolution."),
		}
	}

	if permissionEntity.ID != *permissionRequired.GetPermissionRequiredResources().UserID {

		return &app_errors.PermissionDeniedError{
			Reason: "The resource is owned by a different user account than the one currently authenticated.",
		}
	}

	return nil
}
