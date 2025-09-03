package app

import "log"

func NewApp() IApp {
	log.Println("Initializaing application external services.")
	externalServices := newExternalServices()

	log.Println("Initializing application services PC.")
	internalServices := newInternalService(externalServices)

	return &app{
		externalServices,
		internalServices,
	}
}
