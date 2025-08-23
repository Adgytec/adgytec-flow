package core

import "context"

type IAccessManagementPC interface {
	CheckPermission(context.Context, PermissionRequired) error
}
