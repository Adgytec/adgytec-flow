package user

import (
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

var serviceName = "user"

var userServiceDetails = db.AddServicesIntoStagingParams{
	ID:   core.GetIDFromPayload([]byte(serviceName)),
	Name: serviceName,
	Type: db.GlobalServiceTypeCore,
}
