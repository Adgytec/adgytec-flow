-- name: NewMediaItems :copyfrom
INSERT INTO
	global.media (
		id,
		bucket_path,
		mime_type,
		upload_type,
		upload_id
	)
VALUES
	($1, $2, $3, $4, $5);

-- name: GetMediaDetails :many
SELECT
	*
FROM
	global.media
WHERE
	id = ANY (
		sqlc.arg (media_ids)::UUID[]
	);
