package iam

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

func (s *iamService) resolveApplicationPermission(
	ctx context.Context,
	permissionEntity permissionEntity,
	permissionRequired PermissionProvider,
) error {
	return core.ErrNotImplemented
}
