package iam

func (s *iam) resolveSelfPermission(
	permissionEntity permissionEntity,
	permissionRequired PermissionProvider,
) error {
	// invalid case
	if permissionRequired.GetPermissionRequiredResources().UserID == nil {
		return &PermissionResolutionFailedError{
			Cause: ErrMissingRequiredResourcesValue,
		}
	}

	if permissionEntity.id != *permissionRequired.GetPermissionRequiredResources().UserID {

		return &PermissionDeniedError{
			Reason: "The resource is owned by a different user account than the one currently authenticated.",
		}
	}

	return nil
}
