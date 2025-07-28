package types

import db_actions "github.com/Adgytec/adgytec-flow/database/actions"

type ServiceDetails struct {
	Name             string
	Assignable       bool
	LogicalPartition db_actions.GlobalServiceLogicalPartitionType
}

type ServiceHierarchyDetails struct {
	ServiceName     string
	HierarchyName   string
	HierarchyType   db_actions.GlobalServiceHierarchyType
	HierarchyResult db_actions.GlobalServiceHierarchyResult
}
