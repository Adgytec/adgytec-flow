package iam

import "github.com/Adgytec/adgytec-flow/utils/staging"

func InitIAMService() staging.Details {
	return staging.Details{
		Service:                iamServiceDetails,
		ManagementPermissions:  managementPermissions,
		ApplicationPermissions: applicationPermissions,
		ServiceRestrictions:    nil,
	}
}
