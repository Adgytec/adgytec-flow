-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE VIEW management.user_group_details AS
SELECT
	ug.id,
	ug.name,
	ug.description,
	ug.created_at,
	count(ugu.user_id) AS user_count
FROM
	management.user_groups ug
	LEFT JOIN management.user_group_users ugu ON ug.id = ugu.user_group_id
GROUP BY
	ug.id;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP VIEW if EXISTS management.user_group_details;

-- +goose StatementEnd
