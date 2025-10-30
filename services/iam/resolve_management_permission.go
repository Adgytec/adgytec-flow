package iam

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/actor"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

func (s *iamService) resolveManagementPermission(
	ctx context.Context,
	permissionEntity actor.ActorDetails,
	permissionRequired PermissionProvider,
) error {
	return core.ErrNotImplemented
}
