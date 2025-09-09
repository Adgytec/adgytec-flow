package media

import "github.com/Adgytec/adgytec-flow/database/db"

var serviceName = "media"

var mediaServiceDetails = db.AddServiceParams{
ID:               core.GetIDFromPayload([]byte(serviceName)),
	Name:             serviceName,
	Assignable:       false,
	LogicalPartition: db.GlobalServiceLogicalPartitionTypeNone,

}
