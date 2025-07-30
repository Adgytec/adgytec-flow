package core

import "fmt"

var (
	ErrPermissionDenied      = fmt.Errorf("permission denied")
	ErrPermissionCheckFailed = fmt.Errorf("permission check failed")
)

type PermissionDeniedError struct {
	permission IPermissionRequired
}

func (e *PermissionDeniedError) Error() string {
	return fmt.Sprintf("Permission denied for action: '%s'.", e.permission.Action())
}

func (e *PermissionDeniedError) Is(target error) bool {
	return target == ErrPermissionDenied
}

type PermissionCheckFailedError struct {
	cause error
}

func (e *PermissionCheckFailedError) Error() string {
	return "Permission check failed."
}

func (e *PermissionCheckFailedError) Is(target error) bool {
	return target == ErrPermissionCheckFailed
}

func (e *PermissionCheckFailedError) Unwrap() error {
	return e.cause
}
