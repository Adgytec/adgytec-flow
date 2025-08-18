package user

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

var serviceName = "user"

var userServiceDetails = db_actions.AddServiceParams{
	ID:               helpers.GetIDFromString(serviceName),
	Name:             serviceName,
	Assignable:       false,
	LogicalPartition: db_actions.GlobalServiceLogicalPartitionTypeNone,
}
