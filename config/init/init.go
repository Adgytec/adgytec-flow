package app_init

import (
	"log"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/services/access_management"
	"github.com/Adgytec/adgytec-flow/services/user"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

type serviceFactory func(params app.App) core.ServiceInit

var services = []serviceFactory{
	func(appConfig app.App) core.ServiceInit {
		return access_management.InitAccessManagement(appConfig)
	},
	func(appConfig app.App) core.ServiceInit {
		return user.InitUserService(appConfig)
	},
}

func EnsureServicesInitialization(appConfig app.App) {
	log.Println("Ensuring application initialization.")
	for _, factory := range services {
		serviceInit := factory(appConfig)
		if err := serviceInit.InitService(); err != nil {
			log.Fatalf("error initializaing services. Can't proceed without resolving: %v", err)
		}
	}

}
