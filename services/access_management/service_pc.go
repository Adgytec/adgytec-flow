package access_management

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type accessManagementPC struct {
	service *accessManagement
}

func CreateAccessManagementPC(params iAccessManagementParams) core.IAccessManagementPC {
	log.Println("creating access-management PC")
	return &accessManagementPC{
		service: createAccessManagementService(params),
	}
}
