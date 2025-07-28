-- name: AddManagementPermission :exec
INSERT INTO
	management.permissions (
		key,
		service_name,
		name,
		description,
		required_resources
	)
VALUES
	($1, $2, $3, $4, $5);

-- name: AddApplicationPermission :exec
INSERT INTO
	application.permissions (
		key,
		service_name,
		name,
		description,
		required_resources
	)
VALUES
	($1, $2, $3, $4, $5);
