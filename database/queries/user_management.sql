-- name: NewManagementUser :exec
INSERT INTO
	management.users (id)
VALUES
	($1)
ON CONFLICT (id) DO NOTHING;

-- name: RemoveManagementUser :exec
DELETE FROM management.users
WHERE
	id = $1;
