package iam

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type pc struct {
	service *iam
}

func NewPC(params iamParams) core.AccessManagementPC {
	log.Println("creating access-management PC")
	return &pc{
		service: newService(params),
	}
}
