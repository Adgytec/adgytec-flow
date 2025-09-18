-- name: NewTemporaryMedia :exec
INSERT INTO
	global.temporary_media (
		id,
		bucket_path,
		upload_type,
		media_type,
		upload_id
	)
VALUES
	($1, $2, $3, $4, $5);
