package usermanagement

import "github.com/Adgytec/adgytec-flow/utils/staging"

func InitService() staging.Details {
	return staging.Details{
		Service:                serviceDetails,
		ManagementPermissions:  nil,
		ApplicationPermissions: nil,
		ServiceRestrictions:    nil,
	}
}
