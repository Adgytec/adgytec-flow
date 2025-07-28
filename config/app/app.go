package app

import "github.com/Adgytec/adgytec-flow/utils/interfaces"

type iAppExternalServices interface {
	Auth() interfaces.IAuth
	Database() interfaces.IDatabase
}

type iAppInternalServices interface{}

type IApp interface {
	iAppExternalServices
	iAppInternalServices
}
