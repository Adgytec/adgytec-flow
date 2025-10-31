-- name: NewUserGroupUser :exec
INSERT INTO
	management.user_group_users (
		user_group_id,
		user_id
	)
VALUES
	($1, $2)
ON CONFLICT (
	user_group_id,
	user_id
) DO NOTHING;

-- name: RemoveUserGroupUser :exec
DELETE FROM management.user_group_users
WHERE
	user_group_id = $1
	AND user_id = $2;

-- name: GetUserGroupUsersLatestFirst :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.user_group_users ugu ON ud.id = ugu.user_id
WHERE
	ugu.user_group_id = sqlc.arg ('userGroupID')::UUID
ORDER BY
	ud.created_at DESC
LIMIT
	sqlc.arg ('limit');

-- name: GetUserGroupUsersOldestFirst :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.user_group_users ugu ON ud.id = ugu.user_id
WHERE
	ugu.user_group_id = sqlc.arg ('userGroupID')::UUID
ORDER BY
	ud.created_at ASC
LIMIT
	sqlc.arg ('limit');

-- name: GetUserGroupUsersLatestFirstGreaterThanCursor :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.user_group_users ugu ON ud.id = ugu.user_id
WHERE
	ugu.user_group_id = sqlc.arg ('userGroupID')::UUID
	AND ud.created_at > sqlc.arg ('cursor')::TIMESTAMPTZ
ORDER BY
	ud.created_at DESC
LIMIT
	sqlc.arg ('limit');

-- name: GetUserGroupUsersOldestFirstGreaterThanCursor :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.user_group_users ugu ON ud.id = ugu.user_id
WHERE
	ugu.user_group_id = sqlc.arg ('userGroupID')::UUID
	AND ud.created_at > sqlc.arg ('cursor')::TIMESTAMPTZ
ORDER BY
	ud.created_at ASC
LIMIT
	sqlc.arg ('limit');

-- name: GetUserGroupUsersLatestFirstLesserThanCursor :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.user_group_users ugu ON ud.id = ugu.user_id
WHERE
	ugu.user_group_id = sqlc.arg ('userGroupID')::UUID
	AND ud.created_at < sqlc.arg ('cursor')::TIMESTAMPTZ
ORDER BY
	ud.created_at DESC
LIMIT
	sqlc.arg ('limit');

-- name: GetUserGroupUsersOldestFirstLesserThanCursor :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.user_group_users ugu ON ud.id = ugu.user_id
WHERE
	ugu.user_group_id = sqlc.arg ('userGroupID')::UUID
	AND ud.created_at < sqlc.arg ('cursor')::TIMESTAMPTZ
ORDER BY
	ud.created_at ASC
LIMIT
	sqlc.arg ('limit');

-- name: GetUserGroupUsersByQuery :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.user_group_users ugu ON ud.id = ugu.user_id
WHERE
	ugu.user_group_id = sqlc.arg ('userGroupID')::UUID
	AND (
		ud.normalized_name LIKE sqlc.arg ('query')::TEXT || '%'
		OR ud.normalized_email LIKE sqlc.arg ('query')::TEXT || '%'
	)
ORDER BY
	ud.created_at DESC
LIMIT
	sqlc.arg ('limit');
