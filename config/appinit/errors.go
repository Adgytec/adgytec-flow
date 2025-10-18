package appinit

import (
	"errors"
	"fmt"
)

var (
	ErrAddingServiceDetails = errors.New("error adding service detail")
	ErrAddingPermission     = errors.New("error adding permission")
)

type AddingServiceDetailsError struct {
	cause error
}

func (e *AddingServiceDetailsError) Error() string {
	return fmt.Sprintf("failed to add service details: %v", e.cause)
}

func (e *AddingServiceDetailsError) Is(target error) bool {
	return target == ErrAddingServiceDetails
}

type AddingPermissionError struct {
	permissionType permissionType
	cause          error
}

func (e *AddingPermissionError) Error() string {
	return fmt.Sprintf("failed to add %s permissions: %v", e.permissionType, e.cause)
}

func (e *AddingPermissionError) Is(target error) bool {
	return target == ErrAddingPermission
}
