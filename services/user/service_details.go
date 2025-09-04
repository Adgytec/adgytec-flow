package user

import (
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

var serviceName = "user"

var userServiceDetails = db.AddServiceParams{
	ID:               core.GetIDFromPayload([]byte(serviceName)),
	Name:             serviceName,
	Assignable:       false,
	LogicalPartition: db.GlobalServiceLogicalPartitionTypeNone,
}
