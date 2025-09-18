-- name: GetUserById :one
SELECT
	*
FROM
	global.user_details
WHERE
	id = sqlc.arg (user_id)::UUID;

-- name: GetGlobalUsersLatestFirst :many
SELECT
	*
FROM
	global.user_details
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: GetGlobalUsersOldestFirst :many
SELECT
	*
FROM
	global.user_details
ORDER BY
	created_at ASC
LIMIT
	$1;

-- name: GetGlobalUsersLatestFirstGreaterThanCursor :many
SELECT
	*
FROM
	global.user_details
WHERE
	created_at > sqlc.arg (cursor)::TIMESTAMPTZ
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: GetGlobalUsersOldestFirstGreaterThanCursor :many
SELECT
	*
FROM
	global.user_details
WHERE
	created_at > sqlc.arg (cursor)::TIMESTAMPTZ
ORDER BY
	created_at ASC
LIMIT
	$1;

-- name: GetGlobalUsersLatestFirstLesserThanCursor :many
SELECT
	*
FROM
	global.user_details
WHERE
	created_at < sqlc.arg (cursor)::TIMESTAMPTZ
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: GetGlobalUsersOldestFirstLesserThanCursor :many
SELECT
	*
FROM
	global.user_details
WHERE
	created_at < sqlc.arg (cursor)::TIMESTAMPTZ
ORDER BY
	created_at ASC
LIMIT
	$1;

-- name: GetGlobalUsersByQuery :many
SELECT
	*
FROM
	global.user_details
WHERE
	normalized_name ILIKE '%' || unaccent (
		sqlc.arg ('query')::TEXT
	) || '%'
	OR normalized_email ILIKE '%' || unaccent (
		sqlc.arg ('query')::TEXT
	) || '%'
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: UpdateGlobalUserStatus :exec
UPDATE global.users
SET
	status = $1
WHERE
	id = $2;

-- name: CreateGlobalUser :execrows
INSERT INTO
	global.users (id, email)
VALUES
	($1, $2)
ON CONFLICT (id) DO NOTHING;

-- name: UpdateGlobalUserProfile :one
WITH
	updated AS (
		UPDATE global.users u
		SET
			name = $1,
			about = $2,
			profile_picture_id = $3,
			date_of_birth = $4
		WHERE
			u.id = $5
		RETURNING
			u.id
	)
SELECT
	*
FROM
	global.user_details d
WHERE
	d.id = $5;
