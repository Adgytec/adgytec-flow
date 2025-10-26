package iam

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/markdown"

	md "github.com/nao1215/markdown"
)

var managementPermissions = []db.AddManagementPermissionsIntoStagingParams{
	assignManagementPermission,
	removeManagementPermission,
	listManagementPermission,
}

var assignManagementPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:assign:management-permission", iamServiceDetails.Name),
	ServiceID: iamServiceDetails.ID,
	Name:      "Assign Permission",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Assign Permission").
			PlainText("Grants the ability to assign permissions to any user or group.")
	}),
	RequiredResources: []string{
		string(db.ManagementPermissionResourceTypeOrganization),
	},
	AssignableActor: db.GlobalAssignableActorTypeUser,
}

var removeManagementPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:remove:management-permission", iamServiceDetails.Name),
	ServiceID: iamServiceDetails.ID,
	Name:      "Remove Permission",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Remove Permission").
			PlainText("Grants the ability to remove permissions from any user or group.")
	}),
	RequiredResources: []string{
		string(db.ManagementPermissionResourceTypeOrganization),
	},
	AssignableActor: db.GlobalAssignableActorTypeUser,
}

var listManagementPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:list:management-permission", iamServiceDetails.Name),
	ServiceID: iamServiceDetails.ID,
	Name:      "List Permission",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("List Permission").
			PlainText("Grants the ability to list permissions to any user or group.")
	}),
	RequiredResources: []string{
		string(db.ManagementPermissionResourceTypeOrganization),
	},
	AssignableActor: db.GlobalAssignableActorTypeUser,
}
