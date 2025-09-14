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
		return media.NewMediaServiceCron(appConfig)
	},
}

func ServicesCronJobs(ctx context.Context, appConfig app.App) {
	cronServices := make([]services.Cron, len(appServices))
	for i, factory := range appServices {
		cronServices[i] = factory(appConfig)
	}

	ticker := time.NewTicker(cronInterval)
	defer ticker.Stop()

	// initally trigger immediately
	triggerServicesCron(cronServices)

loop:
	for {
		select {
		case <-ctx.Done():
			{
				break loop
			}
		case <-ticker.C:
			{
				triggerServicesCron(cronServices)
			}
		}
	}

	log.Println("cron jobs ticker cancelled")
}

func triggerServicesCron(cronServices []services.Cron) {
	log.Println("services cron jobs triggered")
	for _, cron := range cronServices {
		// all cron jobs do is some basic db calls and update the field
		// this will be done in short amount of time so no need to use sync mechanisms
		go cron.Trigger()
	}
}
