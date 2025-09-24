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

-- name: AddMediaItemsToOutbox :copyfrom
INSERT INTO
	global.media_outbox (media_id)
VALUES
	($1);

-- name: UpdateMediaItemStatus :exec
UPDATE global.media
SET
	status = $1
WHERE
	id = $2;

-- name: UpdateMediaItemsStatus :exec
UPDATE global.media
SET
	status = $1
WHERE
	id = ANY (
		sqlc.arg (media_ids)::UUID[]
	);
