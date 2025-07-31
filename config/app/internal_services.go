package app

import (
	"github.com/Adgytec/adgytec-flow/services/access_management"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type internalServices struct {
	accessManagement core.IAccessManagementPC
}

func (s *internalServices) AccessManagement() core.IAccessManagementPC {
	return s.accessManagement
}

func createInternalService(externalService iAppExternalServices) iAppInternalServices {
	return &internalServices{
		accessManagement: access_management.CreateAccessManagementPC(externalService),
	}
}
