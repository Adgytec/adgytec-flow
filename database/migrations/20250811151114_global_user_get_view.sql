-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE VIEW global.user_details AS
SELECT
	users.id,
	users.email,
	users.name,
	users.about,
	users.date_of_birth,
	users.created_at,
	users.updated_at,
	users.profile_picture_id,
	media.bucket_path AS uncompressed_profile_picture,
	media.size AS profile_picture_size,
	image.status,
	image.thumbnail,
	image.small,
	image.medium,
	image.large,
	image.extra_large
FROM
	global.users AS users
	LEFT JOIN global.media AS media ON media.id = users.profile_picture_id
	LEFT JOIN global.media_image AS image ON media.id = image.media_id;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP VIEW if EXISTS global.user_details;

-- +goose StatementEnd
