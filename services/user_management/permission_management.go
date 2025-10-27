package usermanagement

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/markdown"

	md "github.com/nao1215/markdown"
)

var managementPermissions = []db.AddManagementPermissionsIntoStagingParams{
	newManagementUserPermission,
	newUserGroupPermission,
	updateUserGroupPermission,
	deleteUserGroupPermission,
	addUserInUserGroupPermission,
	removeUserFromUserGroupPermission,
}

var newManagementUserPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:add:user", serviceDetails.Name),
	ServiceID: serviceDetails.ID,
	Name:      "Add User",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Add User").
			PlainText("Grants the ability to add user for management purposes.")
	}),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var newUserGroupPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:create:user-group", serviceDetails.Name),
	ServiceID: serviceDetails.ID,
	Name:      "Create User Group",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Create User Group").
			PlainText("Grants the ability to create user groups.")
	}),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var updateUserGroupPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:update:user-group", serviceDetails.Name),
	ServiceID: serviceDetails.ID,
	Name:      "Update User Group",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Update User Group").
			PlainText("Grants the ability to update user group details.")
	}),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var deleteUserGroupPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:delete:user-group", serviceDetails.Name),
	ServiceID: serviceDetails.ID,
	Name:      "Delete User Group",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Delete User Group").
			PlainText("Grants the ability to delete user group.")
	}),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var addUserInUserGroupPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:add:user-group-user", serviceDetails.Name),
	ServiceID: serviceDetails.ID,
	Name:      "Add User To User Group",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Add User To User Group").
			PlainText("Grants the ability to add user to a user group.")
	}),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var removeUserFromUserGroupPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:remove:user-group-user", serviceDetails.Name),
	ServiceID: serviceDetails.ID,
	Name:      "Remove User From User Group",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Remove User From User Group").
			PlainText("Grants the ability to remove user from a user group.")
	}),
	RequiredResources: []string{},
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}
