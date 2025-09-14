package appinit

import (
	"log"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/services/iam"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/services"
)

type serviceFactory func(params app.App) services.Init

var appServices = []serviceFactory{
	func(appConfig app.App) services.Init {
		return iam.InitIAMService(appConfig)
	},
	func(appConfig app.App) services.Init {
		return user.InitUserService(appConfig)
	},
}

func EnsureServicesInitialization(appConfig app.App) {
	log.Println("Ensuring application initialization.")
	for _, factory := range appServices {
		serviceInit := factory(appConfig)
		if err := serviceInit.InitService(); err != nil {
			log.Fatalf("error initializaing services. Can't proceed without resolving: %v", err)
		}
	}

}
