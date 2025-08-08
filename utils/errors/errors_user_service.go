package app_errors

import "fmt"

var (
	ErrUserActionFailed = fmt.Errorf("user requested action failed")
)
