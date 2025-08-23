package app_middleware

import "github.com/Adgytec/adgytec-flow/utils/core"

type iAppMiddlewareParams interface {
	Auth() core.IAuth
	UserService() core.IUserServicePC
}

type appMiddleware struct {
	auth        core.IAuth
	userService core.IUserServicePC
}

func createAppMiddleware(params iAppMiddlewareParams) *appMiddleware {
	return &appMiddleware{
		auth:        params.Auth(),
		userService: params.UserService(),
	}
}
