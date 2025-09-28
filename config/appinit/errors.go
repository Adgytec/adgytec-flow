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
	serviceName string
	cause       error
}

func (e *AddingServiceDetailsError) Error() string {
	return fmt.Sprintf("failed to add service details for service %s: %v", e.serviceName, e.cause)
}

func (e *AddingServiceDetailsError) Is(target error) bool {
	return target == ErrAddingServiceDetails
}

type AddingPermissionError struct {
	serviceName     string
	permissionKey   string
	permissionType  permissionType
	cause           error
}

func (e *AddingPermissionError) Error() string {
	return fmt.Sprintf("failed to add %s permission '%s' for service %s: %v", e.permissionType, e.permissionKey, e.serviceName, e.cause)
}

func (e *AddingPermissionError) Is(target error) bool {
	return target == ErrAddingPermission
}
