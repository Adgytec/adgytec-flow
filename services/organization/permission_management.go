package org

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/pointer"
)

var managementPermissions = []db.AddManagementPermissionsIntoStagingParams{
	createOrganizationPermission,
}

var createOrganizationPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:create:organization", orgServiceDetails.Name),
	ServiceID: orgServiceDetails.ID,
	Name:      "Create Organization",
	Description: pointer.New(`
### Create Organization

Grants the ability to create new organization.`),
	RequiredResources: nil,
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}
