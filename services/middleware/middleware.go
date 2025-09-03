package app_middleware

import (
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type appMiddlewareParams interface {
	Auth() core.Auth
	UserService() user.PC
}

type appMiddleware struct {
	auth        core.Auth
	userService user.PC
}

func newAppMiddleware(params appMiddlewareParams) *appMiddleware {
	return &appMiddleware{
		auth:        params.Auth(),
		userService: params.UserService(),
	}
}
