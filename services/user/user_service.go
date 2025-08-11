package user

import "github.com/Adgytec/adgytec-flow/utils/core"

type iUserServiceParams interface {
	Database() core.IDatabase
	Auth() core.IAuth
	AccessManagement() core.IAccessManagementPC
	CacheClient() core.ICacheClient
}

type userService struct {
	db               core.IDatabase
	auth             core.IAuth
	accessManagement core.IAccessManagementPC
}
