package app

import "github.com/rs/zerolog/log"

func NewApp() (App, error) {
	log.Info().Msg("external services init")
	externalServices, externalServiceErr := newExternalServices()
	if externalServiceErr != nil {
		return nil, externalServiceErr
	}

	log.Info().Msg("internal service pc init")
	internalServices, internalServiceErr := newInternalService(externalServices)
	if internalServiceErr != nil {
		return nil, internalServiceErr
	}

	return &app{
		externalServices,
		internalServices,
	}, nil
}
