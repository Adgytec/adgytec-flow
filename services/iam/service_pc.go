package iam

import (
	"context"
	"log"
)

type IAMServicePC interface {
	// CheckPermission checks a single permission and returns nil if it is granted.
	CheckPermission(context.Context, PermissionProvider) error

	// CheckPermissions returns nil if any of the provided permissions are granted.
	CheckPermissions(context.Context, []PermissionProvider) error
}

type iamServicePC struct {
	service *iam
}

func NewIAMServicePC(params iamParams) IAMServicePC {
	log.Println("creating iam PC")
	return &iamServicePC{
		service: newIAMService(params),
	}
}
