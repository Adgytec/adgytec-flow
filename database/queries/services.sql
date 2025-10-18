-- name: AddServiceDetails :exec
INSERT INTO
	global.services (
		id,
		name,
		assignable,
		logical_partition
	)
VALUES
	($1, $2, $3, $4)
ON CONFLICT (id) DO UPDATE
SET
	assignable = excluded.assignable,
	logical_partition = excluded.logical_partition;
