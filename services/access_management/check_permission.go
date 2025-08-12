package access_management

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

func (pc *accessManagementPC) CheckPermission(ctx context.Context, permissionEntity core.IPermissionEntity, permissionRequired core.IPermissionRequired) error {
	return pc.service.checkPermission(ctx, permissionEntity, permissionRequired)
}

func (s *accessManagement) checkPermission(ctx context.Context, entity core.IPermissionEntity, requiredPermission core.IPermissionRequired) error {
	return nil
}
