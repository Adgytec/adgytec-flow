-- name: GetUserById :one
SELECT
	*
FROM
	global.user_details
WHERE
	id = sqlc.arg (user_id)::UUID;

-- name: GetGlobalUsersFromNextCursor :many
SELECT
	*
FROM
	global.user_details
WHERE
	created_at < sqlc.arg (next_cursor)::TIMESTAMPTZ
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: GetGlobalUsersFromPrevCursor :many
SELECT
	*
FROM
	global.user_details
WHERE
	created_at > sqlc.arg (prev_cursor)::TIMESTAMPTZ
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: GetGlobalUsersFromNextCursorOldestFirst :many
SELECT
	*
FROM
	global.user_details
WHERE
	created_at > sqlc.arg (next_cursor)::TIMESTAMPTZ
ORDER BY
	created_at ASC
LIMIT
	$1;

-- name: GetGlobalUsersFromPrevCursorOldestFirst :many
SELECT
	*
FROM
	global.user_details
WHERE
	created_at < sqlc.arg (prev_cursor)::TIMESTAMPTZ
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
