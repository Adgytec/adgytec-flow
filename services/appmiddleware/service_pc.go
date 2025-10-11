package appmiddleware

import (
	"github.com/Adgytec/adgytec-flow/utils/core"
	"github.com/rs/zerolog/log"
)

type appMiddlewarePC struct {
	service *appMiddleware
}

func NewAppMiddlewarePC(params appMiddlewareParams) core.MiddlewarePC {
	log.Info().
		Str("service", serviceName).
		Msg("new service pc")
	return &appMiddlewarePC{
		service: newAppMiddleware(params),
	}
}
