package app

import "log"

func InitApp() IApp {
	log.Println("Initializaing application external services.")
	externalServices := createExternalServices()

	log.Println("Initializing application services PC.")
	internalServices := createInternalService(externalServices)

	return &app{
		externalServices,
		internalServices,
	}
}
