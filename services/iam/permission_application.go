package iam

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/markdown"
	md "github.com/nao1215/markdown"
)

var applicationPermissions = []db.AddApplicationPermissionsIntoStagingParams{
	assignApplicationPermission,
	removeApplicationPermission,
	listApplicationPermission,
}

var assignApplicationPermission = db.AddApplicationPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:assign:permission", iamServiceDetails.Name),
	ServiceID: iamServiceDetails.ID,
	Name:      "Assign Permission",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Assign Permission").
			PlainText("Grants the ability to assign permissions to any user or group.")
	}),
	RequiredResources: nil,
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var removeApplicationPermission = db.AddApplicationPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:remove:permission", iamServiceDetails.Name),
	ServiceID: iamServiceDetails.ID,
	Name:      "Remove Permission",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Remove Permission").
			PlainText("Grants the ability to remove permissions from any user or group.")
	}),
	RequiredResources: nil,
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var listApplicationPermission = db.AddApplicationPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:list:permission", iamServiceDetails.Name),
	ServiceID: iamServiceDetails.ID,
	Name:      "List Permission",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("List Permission").
			PlainText("Grants the ability to list permissions to any user or group.")
	}),
	RequiredResources: nil,
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}
