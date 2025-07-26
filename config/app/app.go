package app

import "github.com/Adgytec/adgytec-flow/utils/interfaces"

type IAppExternalServices interface {
	Auth() interfaces.IAuth
	Database() interfaces.IDatabase
}

type IAppInternalServices interface{}

type IApp interface {
	IAppExternalServices
	IAppInternalServices
}
