package media

import "github.com/Adgytec/adgytec-flow/utils/services"

type mediaServiceCron struct {
	service *mediaService
}

func (c *mediaServiceCron) Trigger() {
	go c.service.cleanInvalidMediaItems()
}

func NewMediaServiceCorn(params mediaServiceParams) services.Cron {
	return &mediaServiceCron{
		service: newMediaService(params),
	}
}
