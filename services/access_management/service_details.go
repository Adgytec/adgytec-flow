package access_management

import (
	db_actions "github.com/Adgytec/adgytec-flow/database/actions"
	"github.com/Adgytec/adgytec-flow/utils/helpers"
)

var serviceName = "access-management"

var accessManagementDetails = db_actions.AddServiceParams{
	ID:               helpers.GetIDFromPayload([]byte(serviceName)),
	Name:             serviceName,
	Assignable:       false,
	LogicalPartition: db_actions.GlobalServiceLogicalPartitionTypeNone,
}
