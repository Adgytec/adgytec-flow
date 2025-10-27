-- name: NewManagementUser :exec
INSERT INTO
	management.users (id)
VALUES
	($1)
ON CONFLICT (id) DO NOTHING;
