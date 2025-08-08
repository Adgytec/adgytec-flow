package app

import "github.com/Adgytec/adgytec-flow/utils/core"

type iAppExternalServices interface {
	Auth() core.IAuth
	Database() core.IDatabase
	Communication() core.ICommunicaiton
	Storage() core.IStorage
	CDN() core.ICDN
}

type iAppInternalServices interface {
	AccessManagement() core.IAccessManagementPC
	UserService() core.IUserServicePC
}

type IApp interface {
	iAppExternalServices
	iAppInternalServices
}

type app struct {
	iAppExternalServices
	iAppInternalServices
}
