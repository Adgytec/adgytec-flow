package access_management

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type accessManagementPC struct {
	service *accessManagement
}

func (pc *accessManagementPC) CheckPermission(permissionEntity core.IPermissionEntity, permissionRequired core.IPermissionRequired) error {
	return pc.service.checkPermission(permissionEntity, permissionRequired)
}

func (pc *accessManagementPC) CheckSelfPermission(currentUserId, userId string) error {
	return pc.service.selfPermissionCheck(currentUserId, userId)
}

func CreateAccessManagementPC(params iAccessManagementParams) core.IAccessManagementPC {
	log.Println("creating access-management PC")
	return &accessManagementPC{
		service: &accessManagement{
			db: params.Database(),
		},
	}
}
