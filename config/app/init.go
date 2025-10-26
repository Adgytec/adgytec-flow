package app

import (
	"github.com/rs/zerolog/log"
)

func NewApp() (App, error) {
	log.Info().Msg("Initializing application external services")
	externalServices, externalServiceErr := newExternalServices()
	if externalServiceErr != nil {
		return nil, externalServiceErr
	}

	log.Info().Msg("Initializing application internal services")
	internalServices, internalServiceErr := newInternalService(externalServices)
	if internalServiceErr != nil {
		return nil, internalServiceErr
	}

	return &app{
		appExternalServices: externalServices,
		appInternalServices: internalServices,
		services:            nil,
	}, nil
}
