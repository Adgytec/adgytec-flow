package appcron

import (
	"context"
	"log"
	"time"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/services/media"
	"github.com/Adgytec/adgytec-flow/utils/services"
)

const cronInterval = 4 * time.Hour

type serviceFactory func(params app.App) services.Cron

var appServices = []serviceFactory{
	func(appConfig app.App) services.Cron {
		return media.NewMediaServiceCorn(appConfig)
	},
}

func ServicesCronJobs(ctx context.Context, appConfig app.App) {
	ticker := time.NewTicker(cronInterval)
	defer ticker.Stop()

	// initally trigger immediately
	triggerServicesCron(appConfig)

loop:
	for {
		select {
		case <-ctx.Done():
			{
				break loop
			}
		case <-ticker.C:
			{
				triggerServicesCron(appConfig)
			}
		}
	}

	log.Println("cron jobs ticker cancelled")
}

func triggerServicesCron(appConfig app.App) {
	log.Println("services cron jobs triggered")
	for _, factory := range appServices {
		appCron := factory(appConfig)
		go appCron.Trigger()
	}
}
