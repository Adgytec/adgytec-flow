package user

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/markdown"

	md "github.com/nao1215/markdown"
)

var managementPermissions = []db.AddManagementPermissionsIntoStagingParams{
	listAllUsersPermission,
	disableUserPermission,
	enableUserPermission,
	getUserProfilePermission,
	updateUserProfilePermission,
}

var listAllUsersPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:list:users", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "List All Users",

	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("List All Users").
			PlainText("Grants the ability to list all the users that are part of Adgytec Studio.").
			PlainTextf("%s", md.Italic("Note: This allows viewing all users regardless of whether they are part of any organization or management."))
	}),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var disableUserPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:disable:users", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "Disable Users",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Disable Users").
			PlainText("Grants the ability to disable users' access to Adgytec Studio.").
			PlainTextf("%s", md.Italic("Note: This disables users globally regardless of the organization they belong to."))
	}),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var enableUserPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:enable:users", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "Enable Users",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Enable Users").
			PlainText("Grants the ability to enable users' access to Adgytec Studio.")
	}),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var getUserProfilePermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:get:user-profile", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "Get User Profile",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Get User Profile").
			PlainText("Grants the ability to get individual user profile details.")
	}),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var updateUserProfilePermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:update:user-profile", userServiceDetails.Name),
	ServiceID: userServiceDetails.ID,
	Name:      "Update User Profile",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Update User Profile").
			PlainText("Grants the ability to update individual user profile.").
			PlainTextf("%s", md.Italic("Note: This also requires 'Get User Profile' permission."))
	}),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}
