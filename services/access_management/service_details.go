package access_management

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
)

var accessManagementDetails = db_actions.AddServiceParams{
	Name:             "access-management",
	Assignable:       false,
	LogicalPartition: db_actions.GlobalServiceLogicalPartitionTypeNone,
}
