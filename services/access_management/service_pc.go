package access_management

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type accessManagementPC struct {
	service *accessManagement
}

func NewAccessManagementPC(params accessManagementParams) core.AccessManagementPC {
	log.Println("creating access-management PC")
	return &accessManagementPC{
		service: newAccessManagementService(params),
	}
}
