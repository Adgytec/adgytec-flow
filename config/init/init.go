package app_init

import (
	"log"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/services/access_management"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type serviceFactory func(params app.IApp) core.IServiceInit

var services = []serviceFactory{
	func(appConfig app.IApp) core.IServiceInit {
		return access_management.InitAccessManagement(appConfig)
	},
}

func EnsureServicesInitialization(appConfig app.IApp) {
	log.Println("Ensuring application initialization.")
	for _, factory := range services {
		serviceInit := factory(appConfig)
		if err := serviceInit.InitService(); err != nil {
			log.Fatalf("error initializaing services. Can't proceed without resolving: %v", err)
		}
	}

}
