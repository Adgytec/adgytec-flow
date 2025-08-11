package access_management

import app_errors "github.com/Adgytec/adgytec-flow/utils/errors"

func (s *accessManagement) selfPermissionCheck(currentUserId, userId, action string) error {
	if userId != currentUserId {
		return &app_errors.PermissionDeniedError{
			Action: action,
		}

	}

	return nil

}
