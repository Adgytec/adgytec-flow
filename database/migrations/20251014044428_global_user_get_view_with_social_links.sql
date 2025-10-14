-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE VIEW global.user_details_with_social_links AS
SELECT
	ud.*,
	jsonb_agg_strict (
		jsonb_build_object(
			'id',
			links.id,
			'platformName',
			links.platform_name,
			'profileLink',
			links.profile_link,
			'createdAt',
			links.created_at,
			'updatedAt',
			links.updated_at
		)
	)
FROM
	global.user_details AS users
	LEFT JOIN global.user_social_links links ON users.id = links.user_id;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP VIEW if EXISTS global.user_details_with_social_links;

-- +goose StatementEnd
