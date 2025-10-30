-- name: NewManagementUser :exec
INSERT INTO
	management.users (id)
VALUES
	($1)
ON CONFLICT (id) DO NOTHING;

-- name: RemoveManagementUser :exec
DELETE FROM management.users
WHERE
	id = $1;

-- name: GetManagementUsersLatestFirst :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.users mu ON ud.id = mu.id
ORDER BY
	ud.created_at DESC
LIMIT
	$1;

-- name: GetManagementUsersOldestFirst :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.users mu ON ud.id = mu.id
ORDER BY
	ud.created_at ASC
LIMIT
	$1;

-- name: GetManagementUsersLatestFirstGreaterThanCursor :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.users mu ON ud.id = mu.id
WHERE
	ud.created_at > sqlc.arg (cursor)::TIMESTAMPTZ
ORDER BY
	ud.created_at DESC
LIMIT
	$1;

-- name: GetManagementUsersOldestFirstGreaterThanCursor :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.users mu ON ud.id = mu.id
WHERE
	ud.created_at > sqlc.arg (cursor)::TIMESTAMPTZ
ORDER BY
	ud.created_at ASC
LIMIT
	$1;

-- name: GetManagementUsersLatestFirstLesserThanCursor :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.users mu ON ud.id = mu.id
WHERE
	ud.created_at < sqlc.arg (cursor)::TIMESTAMPTZ
ORDER BY
	ud.created_at DESC
LIMIT
	$1;

-- name: GetManagementUsersOldestFirstLesserThanCursor :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.users mu ON ud.id = mu.id
WHERE
	ud.created_at < sqlc.arg (cursor)::TIMESTAMPTZ
ORDER BY
	ud.created_at ASC
LIMIT
	$1;

-- name: GetManagementUsersByQuery :many
SELECT
	ud.*
FROM
	global.user_details ud
	JOIN management.users mu ON ud.id = mu.id
WHERE
	ud.normalized_name LIKE sqlc.arg ('query')::TEXT || '%'
	OR ud.normalized_email LIKE sqlc.arg ('query')::TEXT || '%'
ORDER BY
	ud.created_at DESC
LIMIT
	$1;

-- name: ManagementUserExists :one
SELECT
	EXISTS (
		SELECT
			1
		FROM
			management.users
		WHERE
			id = $1
	);
