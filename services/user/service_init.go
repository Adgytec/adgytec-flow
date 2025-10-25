package user

import "github.com/Adgytec/adgytec-flow/utils/staging"

func InitUserService() staging.Details {
	return staging.Details{
		Service:                userServiceDetails,
		ManagementPermissions:  managementPermissions,
		ApplicationPermissions: nil,
		ServiceRestrictions:    nil,
	}
}
