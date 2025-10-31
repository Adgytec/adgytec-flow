-- name: NewUserGroup :one
INSERT INTO
	management.user_groups (name, description)
VALUES
	($1, $2)
RETURNING
	id,
	name,
	description,
	created_at;

-- name: UpdateUserGroup :one
UPDATE management.user_groups
SET
	name = $1,
	description = $2
WHERE
	id = $3
RETURNING
	id,
	name,
	description,
	created_at;

-- name: GetUserGroupByIDForUpdate :one
SELECT
	id,
	name,
	description,
	created_at
FROM
	management.user_groups
WHERE
	id = $1
FOR UPDATE;

-- name: GetUserGroupsLatestFirst :many
SELECT
	*
FROM
	management.user_group_details
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: GetUserGroupsOldestFirst :many
SELECT
	*
FROM
	management.user_group_details
ORDER BY
	created_at ASC
LIMIT
	$1;

-- name: GetUserGroupsLatestFirstGreaterThanCursor :many
SELECT
	*
FROM
	management.user_group_details
WHERE
	created_at > sqlc.arg (cursor)::TIMESTAMPTZ
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: GetUserGroupsOldestFirstGreaterThanCursor :many
SELECT
	*
FROM
	management.user_group_details
WHERE
	created_at > sqlc.arg (cursor)::TIMESTAMPTZ
ORDER BY
	created_at ASC
LIMIT
	$1;

-- name: GetUserGroupsLatestFirstLesserThanCursor :many
SELECT
	*
FROM
	management.user_group_details
WHERE
	created_at < sqlc.arg (cursor)::TIMESTAMPTZ
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: GetUserGroupsOldestFirstLesserThanCursor :many
SELECT
	*
FROM
	management.user_group_details
WHERE
	created_at < sqlc.arg (cursor)::TIMESTAMPTZ
ORDER BY
	created_at ASC
LIMIT
	$1;

-- name: GetUserGroupsByQuery :many
SELECT
	*
FROM
	management.user_group_details
WHERE
	lower(name) LIKE lower(
		sqlc.arg ('query')::TEXT
	) || '%'
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: DeleteUserGroup :exec
DELETE FROM management.user_groups
WHERE
	id = $1;
