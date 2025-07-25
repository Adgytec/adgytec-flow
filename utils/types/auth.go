package types

import "fmt"

type UserExistsError struct {
	username string
}

func (e *UserExistsError) Error() string {
	return fmt.Sprintf("user with username %s already present.", e.username)
}

type AuthActionFailedError struct {
	username string
	reason   string
}

func (e *AuthActionFailedError) Error() string {
	return e.reason
}

func (e *AuthActionFailedError) LoggingMessage() string {
	return fmt.Sprintf("user account creation failed for username: %s.\nerror: %s", e.username, e.reason)
}
