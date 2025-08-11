-- name: GetUserById :one
SELECT
	users.id,
	users.email,
	users.name,
	users.about,
	users.date_of_birth,
	jsonb_build_object(
		'originalProfile',
		media.bucket_path,
		'status',
		image.status,
		'variants',
		image.variants
	) AS profile_picture
FROM
	global.users AS users
	LEFT JOIN global.media AS media ON media.id = users.profile_picture_id
	LEFT JOIN global.media_image AS image ON media.id = image.media_id
WHERE
	users.id = sqlc.arg (user_id)::UUID;

-- name: GetGlobalUsersFromNextCursor :many
SELECT
	users.id,
	users.email,
	users.name,
	users.about,
	users.date_of_birth,
	users.created_at,
	jsonb_build_object(
		'originalProfile',
		media.bucket_path,
		'status',
		image.status,
		'variants',
		image.variants
	) AS profile_picture
FROM
	global.users AS users
	LEFT JOIN global.media AS media ON media.id = users.profile_picture_id
	LEFT JOIN global.media_image AS image ON media.id = image.media_id
WHERE
	users.created_at < sqlc.arg (next_cursor)::TIMESTAMPTZ
ORDER BY
	users.created_at DESC
LIMIT
	$1;

-- name: GetGlobalUsersFromPrevCursor :many
SELECT
	users.id,
	users.email,
	users.name,
	users.about,
	users.date_of_birth,
	users.created_at,
	jsonb_build_object(
		'originalProfile',
		media.bucket_path,
		'status',
		image.status,
		'variants',
		image.variants
	) AS profile_picture
FROM
	global.users AS users
	LEFT JOIN global.media AS media ON media.id = users.profile_picture_id
	LEFT JOIN global.media_image AS image ON media.id = image.media_id
WHERE
	users.created_at > sqlc.arg (prev_cursor)::TIMESTAMPTZ
ORDER BY
	users.created_at DESC
LIMIT
	$1;

-- name: GetGlobalUsersFromNextCursorOldestFirst :many
SELECT
	users.id,
	users.email,
	users.name,
	users.about,
	users.date_of_birth,
	users.created_at,
	jsonb_build_object(
		'originalProfile',
		media.bucket_path,
		'status',
		image.status,
		'variants',
		image.variants
	) AS profile_picture
FROM
	global.users AS users
	LEFT JOIN global.media AS media ON media.id = users.profile_picture_id
	LEFT JOIN global.media_image AS image ON media.id = image.media_id
WHERE
	users.created_at > sqlc.arg (next_cursor)::TIMESTAMPTZ
ORDER BY
	users.created_at ASC
LIMIT
	$1;

-- name: GetGlobalUsersFromPrevCursorOldestFirst :many
SELECT
	users.id,
	users.email,
	users.name,
	users.about,
	users.date_of_birth,
	users.created_at,
	jsonb_build_object(
		'originalProfile',
		media.bucket_path,
		'status',
		image.status,
		'variants',
		image.variants
	) AS profile_picture
FROM
	global.users AS users
	LEFT JOIN global.media AS media ON media.id = users.profile_picture_id
	LEFT JOIN global.media_image AS image ON media.id = image.media_id
WHERE
	users.created_at < sqlc.arg (prev_cursor)::TIMESTAMPTZ
ORDER BY
	users.created_at ASC
LIMIT
	$1;

-- name: GetGlobalUsersByQuery :many
SELECT
	users.id,
	users.email,
	users.name,
	users.about,
	users.date_of_birth,
	users.created_at,
	jsonb_build_object(
		'originalProfile',
		media.bucket_path,
		'status',
		image.status,
		'variants',
		image.variants
	) AS profile_picture
FROM
	global.users AS users
	LEFT JOIN global.media AS media ON media.id = users.profile_picture_id
	LEFT JOIN global.media_image AS image ON media.id = image.media_id
WHERE
	users.normalized_name ILIKE '%' || unaccent (
		sqlc.arg ('query')::TEXT
	) || '%'
	OR users.normalized_email ILIKE '%' || unaccent (
		sqlc.arg ('query')::TEXT
	) || '%'
ORDER BY
	users.created_at DESC
LIMIT
	$1;
