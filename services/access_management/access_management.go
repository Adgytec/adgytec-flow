package access_management

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	app_errors "github.com/Adgytec/adgytec-flow/utils/errors"
)

type iAccessManagementParams interface {
	Database() core.IDatabase
	CacheClient() core.ICacheClient
}

type accessManagement struct {
	db              core.IDatabase
	permissionCache core.ICache[bool]
}

func (s *accessManagement) checkPermission(entity core.IPermissionEntity, requiredPermission core.IPermissionRequired) error {
	return nil
}

func (s *accessManagement) selfPermissionCheck(currentUserId, userId string) error {
	if userId != currentUserId {
		return app_errors.ErrPermissionDenied
	}

	return nil

}
