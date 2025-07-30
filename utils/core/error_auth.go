package core

import "fmt"

var (
	ErrUserExists       = fmt.Errorf("user exists")
	ErrAuthActionFailed = fmt.Errorf("auth action failed")
)

type UserExistsError struct {
	username string
}

func (e *UserExistsError) Error() string {
	return fmt.Sprintf("User with username %s already exists.", e.username)
}

func (e *UserExistsError) Is(target error) bool {
	return target == ErrUserExists
}

type AuthActionFailedError struct {
	username   string
	reason     string
	actionType AuthActionType
	cause      error
}

func (e *AuthActionFailedError) Error() string {
	return fmt.Sprintf("Auth action failed.\nAction Type: '%s' \nUsername: %s\nReason: %s", e.actionType, e.username, e.reason)
}

func (e *AuthActionFailedError) Is(target error) bool {
	return target == ErrAuthActionFailed
}

func (e *AuthActionFailedError) Unwrap() error {
	return e.cause
}
