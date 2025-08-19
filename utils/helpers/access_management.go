package helpers

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

func CreatePermissionRequiredFromManagementPermission(permission db_actions.AddManagementPermissionParams, requiredResourcesId []string) core.PermissionRequired {
	var requiredResources []string
	for _, resourceType := range permission.RequiredResources {
		requiredResources = append(requiredResources, string(resourceType))
	}

	return core.PermissionRequired{
		Key:                 permission.Key,
		Management:          true,
		RequiredResources:   requiredResources,
		RequiredResourcesId: requiredResourcesId,
	}
}

func CreatePermissionRequiredFromApplicationPermission(permission db_actions.AddApplicationPermissionParams, orgId string, requiredResourcesId []string) core.PermissionRequired {
	var requiredResources []string
	for _, resourceType := range permission.RequiredResources {
		requiredResources = append(requiredResources, string(resourceType))
	}

	return core.PermissionRequired{
		Key:                 permission.Key,
		Management:          false,
		OrgId:               orgId,
		RequiredResources:   requiredResources,
		RequiredResourcesId: requiredResourcesId,
	}
}
