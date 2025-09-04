package app_middleware

import (
	"github.com/Adgytec/adgytec-flow/config/auth"
	"github.com/Adgytec/adgytec-flow/services/user"
)

type appMiddlewareParams interface {
	Auth() auth.Auth
	UserService() user.UserServicePC
}

type appMiddleware struct {
	auth        auth.Auth
	userService user.UserServicePC
}

func newAppMiddleware(params appMiddlewareParams) *appMiddleware {
	return &appMiddleware{
		auth:        params.Auth(),
		userService: params.UserService(),
	}
}
