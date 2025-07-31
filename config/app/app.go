package app

import "github.com/Adgytec/adgytec-flow/utils/core"

type iAppExternalServices interface {
	Auth() core.IAuth
	Database() core.IDatabase
}

type iAppInternalServices interface {
	AccessManagement() core.IAccessManagementPC
}

type IApp interface {
	iAppExternalServices
	iAppInternalServices
}
