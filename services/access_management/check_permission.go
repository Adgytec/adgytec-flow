package access_management

import "github.com/Adgytec/adgytec-flow/utils/core"

func (pc *accessManagementPC) CheckPermission(permissionEntity core.IPermissionEntity, permissionRequired core.IPermissionRequired) error {
	return pc.service.checkPermission(permissionEntity, permissionRequired)
}

func (s *accessManagement) checkPermission(entity core.IPermissionEntity, requiredPermission core.IPermissionRequired) error {
	return nil
}
