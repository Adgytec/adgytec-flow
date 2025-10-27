package org

import (
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

var serviceName = "organization"

var orgServiceDetails = db.AddServicesIntoStagingParams{
	ID:   core.GetIDFromPayload([]byte(serviceName)),
	Name: serviceName,
	Type: db.GlobalServiceTypePlatform,
}
