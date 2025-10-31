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

-- name: RemoveUserGroupUser :exec
DELETE FROM management.user_group_users
WHERE
	user_group_id = $1
	AND user_id = $2;
