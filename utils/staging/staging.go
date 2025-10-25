package staging

import "github.com/Adgytec/adgytec-flow/database/db"

type Details struct {
	Service                db.AddServicesIntoStagingParams
	ManagementPermissions  []db.AddManagementPermissionsIntoStagingParams
	ApplicationPermissions []db.AddApplicationPermissionsIntoStagingParams
	ServiceRestrictions    []db.AddServiceRestrictionIntoStagingParams
}
