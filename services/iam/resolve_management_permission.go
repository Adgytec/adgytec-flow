package iam

import (
	"context"

	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
)

func (s *iam) resolveManagementPermission(
	ctx context.Context,
	permissionEntity permissionEntity,
	permissionRequired PermissionProvider,
) error {
	return app_errors.ErrNotImplemented
}
