package usermanagement

import (
	"fmt"

	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/markdown"

	md "github.com/nao1215/markdown"
)

var managementPermissions = []db.AddManagementPermissionsIntoStagingParams{
	newManagementUserPermission,
	listManagementUsersPermission,
	removeManagementUserPermission,
	newUserGroupPermission,
	listUserGroupsPermission,
	updateUserGroupPermission,
	deleteUserGroupPermission,
	addUserInUserGroupPermission,
	listUserGroupUsersPermission,
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
	RequiredResources: nil,
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var removeManagementUserPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:remove:user", serviceDetails.Name),
	ServiceID: serviceDetails.ID,
	Name:      "Remove User",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("Remove User").
			PlainText("Grants the ability to remove user from management.")
	}),
	RequiredResources: nil,
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var listManagementUsersPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:list:users", serviceDetails.Name),
	ServiceID: serviceDetails.ID,
	Name:      "List Users",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("List Users").
			PlainText("Grants the ability to list all management users.")
	}),
	RequiredResources: nil,
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
	RequiredResources: nil,
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var listUserGroupsPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:list:user-groups", serviceDetails.Name),
	ServiceID: serviceDetails.ID,
	Name:      "List User Groups",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("List User Groups").
			PlainText("Grants the ability to list all user groups.")
	}),
	RequiredResources: nil,
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
	RequiredResources: nil,
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
	RequiredResources: nil,
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
	RequiredResources: nil,
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}

var listUserGroupUsersPermission = db.AddManagementPermissionsIntoStagingParams{
	Key:       fmt.Sprintf("%s:list:user-group-users", serviceDetails.Name),
	ServiceID: serviceDetails.ID,
	Name:      "List User Group Users",
	Description: markdown.BuildMarkdown(func(m *md.Markdown) {
		m.H3("List User Group Users").
			PlainText("Grants the ability to list user group users.")
	}),
	RequiredResources: nil,
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
	RequiredResources: nil,
	AssignableActor:   db.GlobalAssignableActorTypeUser,
}
