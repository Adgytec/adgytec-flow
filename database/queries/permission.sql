-- name: AddManagementPermission :exec
INSERT INTO
	management.permissions (
		key,
		service_id,
		name,
		description,
		required_resources
	)
VALUES
	($1, $2, $3, $4, $5)
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description;

-- name: AddApplicationPermission :exec
INSERT INTO
	application.permissions (
		key,
		service_id,
		name,
		description,
		required_resources
	)
VALUES
	($1, $2, $3, $4, $5)
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description;
