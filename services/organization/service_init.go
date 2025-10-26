package org

import "github.com/Adgytec/adgytec-flow/utils/staging"

func InitOrgService() staging.Details {
	return staging.Details{
		Service:                orgServiceDetails,
		ManagementPermissions:  managementPermissions,
		ApplicationPermissions: nil,
		ServiceRestrictions:    nil,
	}
}
