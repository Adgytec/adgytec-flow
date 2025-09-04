package iam

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

func (s *iamService) resolveManagementPermission(
	ctx context.Context,
	permissionEntity permissionEntity,
	permissionRequired PermissionProvider,
) error {
	return core.ErrNotImplemented
}
