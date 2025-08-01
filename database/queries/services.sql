-- name: AddService :exec
INSERT INTO
	global.services (
		name,
		assignable,
		logical_partition
	)
VALUES
	($1, $2, $3)
ON CONFLICT (name) DO UPDATE
SET
	assignable = excluded.assignable,
	logical_partition = excluded.logical_partition;

-- name: AddServiceHierarchyDetails :exec
INSERT INTO
	global.service_hierarchy_details (
		service_name,
		hierarchy_name,
		hierarchy_type,
		hierarchy_result
	)
VALUES
	($1, $2, $3, $4)
ON CONFLICT (service_name) DO UPDATE
SET
	hierarchy_name = excluded.hierarchy_name,
	hierarchy_type = excluded.hierarchy_type,
	hierarchy_result = excluded.hierarchy_result;
