-- name: NewUserGroupUser :exec
INSERT INTO
	management.user_group_users (
		user_group_id,
		user_id
	)
VALUES
	($1, $2)
ON CONFLICT (
	user_group_id,
	user_id
) DO NOTHING;
