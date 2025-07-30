package core

import "fmt"

type UserExistsError struct {
	username string
}

func (e *UserExistsError) Error() string {
	return fmt.Sprintf("User with username %s already exists.", e.username)
}

type AuthActionFailedError struct {
	username   string
	reason     string
	actionType AuthActionType
}

func (e *AuthActionFailedError) Error() string {
	return fmt.Sprintf("Auth action failed.\nAction Type: '%s' \nUsername: %s\nReason: %s", e.actionType, e.username, e.reason)
}
