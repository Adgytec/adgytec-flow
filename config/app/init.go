package app

import "log"

func NewApp() (App, error) {
	log.Println("Initializaing application external services.")
	externalServices, externalServiceErr := newExternalServices()
	if externalServiceErr != nil {
		return nil, externalServiceErr
	}

	log.Println("Initializing application services PC.")
	internalServices := newInternalService(externalServices)

	return &app{
		externalServices,
		internalServices,
	}, nil
}
