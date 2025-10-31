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
	ug.id,
	ug.name,
	ug.description,
	ug.created_at,
	count(ugu.user_id) AS user_count
FROM
	management.user_groups ug
	LEFT JOIN management.user_group_users ugu ON ug.id = ugu.user_group_id
GROUP BY
	ug.id
ORDER BY
	ug.created_at DESC
LIMIT
	$1;

-- name: GetUserGroupsOldestFirst :many
SELECT
	ug.id,
	ug.name,
	ug.description,
	ug.created_at,
	count(ugu.user_id) AS user_count
FROM
	management.user_groups ug
	LEFT JOIN management.user_group_users ugu ON ug.id = ugu.user_group_id
GROUP BY
	ug.id
ORDER BY
	ug.created_at ASC
LIMIT
	$1;

-- name: GetUserGroupsLatestFirstGreaterThanCursor :many
SELECT
	ug.id,
	ug.name,
	ug.description,
	ug.created_at,
	count(ugu.user_id) AS user_count
FROM
	management.user_groups ug
	LEFT JOIN management.user_group_users ugu ON ug.id = ugu.user_group_id
WHERE
	ug.created_at > sqlc.arg (cursor)::TIMESTAMPTZ
GROUP BY
	ug.id
ORDER BY
	ug.created_at DESC
LIMIT
	$1;

-- name: GetUserGroupsOldestFirstGreaterThanCursor :many
SELECT
	ug.id,
	ug.name,
	ug.description,
	ug.created_at,
	count(ugu.user_id) AS user_count
FROM
	management.user_groups ug
	LEFT JOIN management.user_group_users ugu ON ug.id = ugu.user_group_id
WHERE
	ug.created_at > sqlc.arg (cursor)::TIMESTAMPTZ
GROUP BY
	ug.id
ORDER BY
	ug.created_at ASC
LIMIT
	$1;

-- name: GetUserGroupsLatestFirstLesserThanCursor :many
SELECT
	ug.id,
	ug.name,
	ug.description,
	ug.created_at,
	count(ugu.user_id) AS user_count
FROM
	management.user_groups ug
	LEFT JOIN management.user_group_users ugu ON ug.id = ugu.user_group_id
WHERE
	ug.created_at < sqlc.arg (cursor)::TIMESTAMPTZ
GROUP BY
	ug.id
ORDER BY
	ug.created_at DESC
LIMIT
	$1;

-- name: GetUserGroupsOldestFirstLesserThanCursor :many
SELECT
	ug.id,
	ug.name,
	ug.description,
	ug.created_at,
	count(ugu.user_id) AS user_count
FROM
	management.user_groups ug
	LEFT JOIN management.user_group_users ugu ON ug.id = ugu.user_group_id
WHERE
	ug.created_at < sqlc.arg (cursor)::TIMESTAMPTZ
GROUP BY
	ug.id
ORDER BY
	ug.created_at ASC
LIMIT
	$1;

-- name: GetUserGroupsByQuery :many
SELECT
	ug.id,
	ug.name,
	ug.description,
	ug.created_at,
	count(ugu.user_id) AS user_count
FROM
	management.user_groups ug
	LEFT JOIN management.user_group_users ugu ON ug.id = ugu.user_group_id
WHERE
	lower(ug.name) LIKE lower(
		sqlc.arg ('query')::TEXT
	) || '%'
GROUP BY
	ug.id
ORDER BY
	ug.created_at DESC
LIMIT
	$1;
