package appinit

import (
	"errors"
	"fmt"
)

var (
	ErrAddingServiceDetails      = errors.New("error adding service detail")
	ErrAddingPermission          = errors.New("error adding permission")
	ErrAddingServiceRestrictions = errors.New("error adding service restrictions")
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

type AddServiceRestrictionsError struct {
	cause error
}

func (e *AddServiceRestrictionsError) Error() string {
	return fmt.Sprintf("failed to add service restrictions: %v", e.cause)
}

func (e *AddServiceRestrictionsError) Is(target error) bool {
	return target == ErrAddingServiceRestrictions
}
