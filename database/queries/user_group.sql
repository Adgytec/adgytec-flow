-- name: NewUserGroup :one
INSERT INTO
	management.user_groups (name, description)
VALUES
	($1, $2)
RETURNING
	id,
	name,
	description,
	created_at;
