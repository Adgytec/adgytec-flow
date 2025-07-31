package access_management

import "github.com/Adgytec/adgytec-flow/utils/core"

type accessManagementPC struct {
	service *accessManagement
}

func (b *accessManagementPC) CheckPermission(core.IPermissionEntity, core.IPermissionRequired) error {
	return nil
}

func CreateAccessManagementPC(params iAccessManagementParams) core.IAccessManagementPC {
	return &accessManagementPC{
		service: &accessManagement{
			db: params.Database(),
		},
	}
}
