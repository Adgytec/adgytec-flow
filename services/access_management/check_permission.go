package access_management

import (
	"context"
	"fmt"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

func (pc *accessManagementPC) CheckPermission(ctx context.Context, permissionEntity core.IPermissionEntity, permissionRequired core.IPermissionRequired) error {
	return pc.service.checkPermission(ctx, permissionEntity, permissionRequired)
}

func (s *accessManagement) checkPermission(ctx context.Context, entity core.IPermissionEntity, requiredPermission core.IPermissionRequired) error {
	// TODO: Implement actual permission checking logic here.
	// For now, returning an error to prevent accidental use in production.
	return fmt.Errorf("permission check not implemented")
}
