package access_management

import (
	"github.com/Adgytec/adgytec-flow/utils/interfaces"
)

type accessManagementInit struct {
	db interfaces.IDatabase
}

func (i *accessManagementInit) InitService() error {
	if err := i.initServiceDetails(); err != nil {
		return err
	}

	if err := i.initServiceManagementPermissions(); err != nil {
		return err
	}

	if err := i.initServiceApplicationPermissions(); err != nil {
		return err
	}

	return nil
}

func (i *accessManagementInit) initServiceDetails() error {
	return nil
}

func (i *accessManagementInit) initServiceManagementPermissions() error {
	return nil
}

func (i *accessManagementInit) initServiceApplicationPermissions() error {
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
