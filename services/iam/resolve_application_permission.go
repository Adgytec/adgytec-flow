package iam

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
)

func (s *iam) resolveApplicationPermission(
	ctx context.Context,
	permissionEntity core.PermissionEntity,
	permissionRequired core.PermissionProvider,
) error {
	return app_errors.ErrNotImplemented
}
