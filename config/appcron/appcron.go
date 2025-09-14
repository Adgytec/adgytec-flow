package appcron

import (
	"context"
	"log"
	"os"

	"github.com/Adgytec/adgytec-flow/config/app"
	"github.com/Adgytec/adgytec-flow/services/media"
	"github.com/Adgytec/adgytec-flow/utils/services"
	"github.com/robfig/cron/v3"
)

type serviceFactory func(params app.App) services.Cron

var appServices = []serviceFactory{
	func(appConfig app.App) services.Cron {
		return media.NewMediaServiceCron(appConfig)
	},
}

func ServicesCronJobs(ctx context.Context, appConfig app.App) {
	cronExpr := os.Getenv("CRON_EXPR")
	if cronExpr == "" {
		// default every 4 hours
		cronExpr = "@every 4h"
	}

	c := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cron.DefaultLogger),
		cron.Recover(cron.DefaultLogger),
	))

	// build services
	cronServices := make([]services.Cron, len(appServices))
	for i, factory := range appServices {
		cronServices[i] = factory(appConfig)
	}

	// schedule jobs
	_, err := c.AddFunc(cronExpr, func() {
		triggerServicesCron(cronServices)
	})
	if err != nil {
		log.Printf("failed to add cron job: %v", err)
	}

	// run once immediately
	triggerServicesCron(cronServices)

	// start scheduler
	c.Start()

	// wait for cancellation
	<-ctx.Done()

	// stop scheduler gracefully
	c.Stop()
	log.Println("cron jobs stopped")
}

func triggerServicesCron(cronServices []services.Cron) {
	log.Println("services cron jobs triggered")
	for _, cron := range cronServices {
		cron.Trigger()
	}
}
