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

-- name: UpdateUserGroup :one
UPDATE management.user_groups
SET
	name = $1,
	description = $2
WHERE
	id = $3
RETURNING
	id,
	name,
	description,
	created_at;

-- name: GetUserGroupByIDForUpdate :one
SELECT
	id,
	name,
	description,
	created_at
FROM
	management.user_groups
WHERE
	id = $1
FOR UPDATE;
