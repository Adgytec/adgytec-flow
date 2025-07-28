-- name: AddService :exec
INSERT INTO
	global.services (
		name,
		assignable,
		logical_partition
	)
VALUES
	($1, $2, $3);

-- name: AddServiceHierarchyDetails :exec
INSERT INTO
	global.service_hierarchy_details (
		service_name,
		hierarchy_name,
		hierarchy_type,
		hierarchy_result
	)
VALUES
	($1, $2, $3, $4);
