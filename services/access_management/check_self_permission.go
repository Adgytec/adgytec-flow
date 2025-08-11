package access_management

import app_errors "github.com/Adgytec/adgytec-flow/utils/errors"

func (pc *accessManagementPC) CheckSelfPermission(currentUserId, userId, action string) error {
	return pc.service.checkSelfPermission(currentUserId, userId, action)
}

func (s *accessManagement) checkSelfPermission(currentUserId, userId, action string) error {
	if userId != currentUserId {
		return &app_errors.PermissionDeniedError{
			Action: action,
		}

	}

	return nil

}
