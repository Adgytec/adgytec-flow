package iam

import (
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/core"
)

var serviceName = "iam"

var iamServiceDetails = db.AddServiceDetailsParams{
	ID:               core.GetIDFromPayload([]byte(serviceName)),
	Name:             serviceName,
	Assignable:       false,
	LogicalPartition: db.GlobalServiceLogicalPartitionTypeNone,
}
