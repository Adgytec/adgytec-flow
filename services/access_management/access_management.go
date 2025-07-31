package access_management

import "github.com/Adgytec/adgytec-flow/utils/core"

type iAccessManagementParams interface {
	Database() core.IDatabase
}

type accessManagement struct {
	db core.IDatabase
}
