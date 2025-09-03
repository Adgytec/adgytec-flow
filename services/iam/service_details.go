package iam

import (
	"github.com/Adgytec/adgytec-flow/database/db"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

var serviceName = "access-management"

var accessManagementDetails = db.AddServiceParams{
	ID:               helpers.GetIDFromPayload([]byte(serviceName)),
	Name:             serviceName,
	Assignable:       false,
	LogicalPartition: db.GlobalServiceLogicalPartitionTypeNone,
}
