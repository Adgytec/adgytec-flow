package user

import db_actions "github.com/Adgytec/adgytec-flow/database/actions"

var userServiceDetails = db_actions.AddServiceParams{
	Name:             "user",
	Assignable:       false,
	LogicalPartition: db_actions.GlobalServiceLogicalPartitionTypeNone,
}
