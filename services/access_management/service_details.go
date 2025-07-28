package access_management

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/types"
)

var accessManagementDetails = types.ServiceDetails{
	Name:             "access-management",
	Assignable:       true,
	LogicalPartition: db_actions.GlobalServiceLogicalPartitionTypeNone,
}
