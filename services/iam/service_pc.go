package iam

import (
	"context"

	"github.com/rs/zerolog/log"
)

type IAMServicePC interface {
	// CheckPermission checks a single permission and returns nil if it is granted.
	CheckPermission(context.Context, PermissionProvider) error

	// CheckPermissions returns nil if any of the provided permissions are granted.
	CheckPermissions(context.Context, []PermissionProvider) error
}

type iamServicePC struct {
	service *iamService
}

func NewIAMServicePC(params iamServiceParams) IAMServicePC {
	log.Info().
		Str("service", serviceName).
		Msg("new service pc")
	return &iamServicePC{
		service: newIAMService(params),
	}
}
