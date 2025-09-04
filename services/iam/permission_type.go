package iam

// PermissionType defines 'Type' of permission
type PermissionType string

const (
	PermissionTypeSelf        PermissionType = "self"
	PermissionTypeManagement  PermissionType = "management"
	PermissionTypeApplication PermissionType = "application"
)
