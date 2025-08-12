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
	ROW (
		media.bucket_path,
		media.size,
		image.status,
		image.variants
	)::global.image_query_type AS profile_picture
FROM
	global.users AS users
	LEFT JOIN global.media AS media ON media.id = users.profile_picture_id
	LEFT JOIN global.media_image AS image ON media.id = image.media_id;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP VIEW if EXISTS global.user_details;

-- +goose StatementEnd
