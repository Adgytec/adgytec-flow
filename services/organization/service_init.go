package org

import "github.com/Adgytec/adgytec-flow/database/db"

func InitOrgService() (db.AddServicesIntoStagingParams, []db.AddManagementPermissionsIntoStagingParams, []db.AddApplicationPermissionsIntoStagingParams, []db.AddServiceRestrictionIntoStagingParams) {
	return orgServiceDetails, nil, nil, nil
}
