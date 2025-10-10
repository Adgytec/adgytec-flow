-- name: AddMediaItems :copyfrom
INSERT INTO
	global.media (
		id,
		bucket_path,
		required_mime_type,
		upload_type,
		upload_id
	)
VALUES
	($1, $2, $3, $4, $5);

-- name: GetMediaItemDetails :one
SELECT
	bucket_path AS key,
	upload_id,
	upload_type
FROM
	global.media
WHERE
	id = $1;
