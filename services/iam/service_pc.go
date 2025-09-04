package iam

import (
	"context"
	"log"
)

type PC interface {
	// CheckPermission checks a single permission and returns nil if it is granted.
	CheckPermission(context.Context, PermissionProvider) error

	// CheckPermissions returns nil if any of the provided permissions are granted.
	CheckPermissions(context.Context, []PermissionProvider) error
}

type pc struct {
	service *iam
}

func NewPC(params iamParams) PC {
	log.Println("creating access-management PC")
	return &pc{
		service: newIAMService(params),
	}
}
