package access_management

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type accessManagementPC struct {
	service *accessManagement
}

func (b *accessManagementPC) CheckPermission(permissionEntity core.IPermissionEntity, permissionRequired core.IPermissionRequired) error {
	return b.service.checkPermission(permissionEntity, permissionRequired)
}

func (b *accessManagementPC) CheckSelfPermission(currentUserId, userId string) error {
	return b.service.selfPermissionCheck(currentUserId, userId)
}

func CreateAccessManagementPC(params iAccessManagementParams) core.IAccessManagementPC {
	log.Println("creating access-management PC")
	return &accessManagementPC{
		service: &accessManagement{
			db: params.Database(),
		},
	}
}
