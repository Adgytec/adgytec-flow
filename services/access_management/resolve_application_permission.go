package access_management

import (
	"context"

	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
)

func (s *accessManagement) resolveApplicationPermission(
	ctx context.Context,
	permissionEntity core.PermissionEntity,
	permissionRequired core.IPermissionRequired,
) error {
	return app_errors.ErrNotImplemented
}
