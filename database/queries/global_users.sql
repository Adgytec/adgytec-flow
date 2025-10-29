-- name: GetUserById :one
SELECT
	*
FROM
	global.user_details
WHERE
	id = sqlc.arg (user_id)::UUID;

-- name: GetUserSocialLinks :many
SELECT
	id,
	platform_name,
	profile_link,
	created_at,
	updated_at
FROM
	global.user_social_links
WHERE
	user_id = sqlc.arg (user_id)::UUID;

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
	normalized_name LIKE unaccent (
		sqlc.arg ('query')::TEXT
	) || '%'
	OR normalized_email LIKE unaccent (
		sqlc.arg ('query')::TEXT
	) || '%'
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: UpdateGlobalUserStatus :one
UPDATE global.users
SET
	status = $1
WHERE
	id = $2
RETURNING
	email;

-- name: CreateGlobalUser :exec
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

-- name: NewUserSocialLink :one
INSERT INTO
	global.user_social_links (
		platform_name,
		profile_link,
		user_id
	)
VALUES
	($1, $2, $3)
RETURNING
	*;

-- name: RemoveUserSocialLink :execrows
DELETE FROM global.user_social_links
WHERE
	id = $1
	AND user_id = $2;

-- name: UpdateUserSocialLink :one
UPDATE global.user_social_links
SET
	profile_link = $1
WHERE
	id = $2
	AND user_id = $3
RETURNING
	*;
