package app_middleware

import "github.com/Adgytec/adgytec-flow/utils/core"

type iAppMiddlewareParams interface {
	Auth() core.Auth
	UserService() core.UserServicePC
}

type appMiddleware struct {
	auth        core.Auth
	userService core.UserServicePC
}

func newAppMiddleware(params iAppMiddlewareParams) *appMiddleware {
	return &appMiddleware{
		auth:        params.Auth(),
		userService: params.UserService(),
	}
}
