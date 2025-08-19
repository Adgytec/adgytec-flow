package core

import "context"

type IAccessManagementPC interface {
	CheckPermission(context.Context, PermissionRequired) error
	CheckSelfPermission(currentUserId string, userId string, action string) error
}
