package access_management

import (
	"github.com/Adgytec/adgytec-flow/utils/interfaces"
)

type accessManagementInit struct {
	db interfaces.IDatabase
}

func (i *accessManagementInit) InitService() error {
	return nil
}

type iAccessManagementInitParams interface {
	Database() interfaces.IDatabase
}

func InitAccessManagement(params iAccessManagementInitParams) interfaces.IServiceInit {
	return &accessManagementInit{
		db: params.Database(),
	}
}
