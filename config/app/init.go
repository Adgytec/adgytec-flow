package app

import "log"

func NewApp() App {
	log.Println("Initializaing application external services.")
	externalServices := newExternalServices()

	log.Println("Initializing application services PC.")
	internalServices := newInternalService(externalServices)

	return &app{
		externalServices,
		internalServices,
	}
}
