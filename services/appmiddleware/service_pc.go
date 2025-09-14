package appmiddleware

import (
	"log"

	"github.com/Adgytec/adgytec-flow/utils/core"
)

type appMiddlewarePC struct {
	service *appMiddleware
}

func NewAppMiddlewarePC(params appMiddlewareParams) core.MiddlewarePC {
	log.Println("creating app-middleware PC")
	return &appMiddlewarePC{
		service: newAppMiddleware(params),
	}
}
