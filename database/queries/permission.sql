-- name: AddManagementPermission :exec
INSERT INTO
	management.permissions (
		key,
		service_id,
		name,
		description,
		required_resources,
		api_key_assignable
	)
VALUES
	(
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	)
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description,
	api_key_assignable = excluded.api_key_assignable;

-- name: AddApplicationPermission :exec
INSERT INTO
	application.permissions (
		key,
		service_id,
		name,
		description,
		required_resources,
		api_key_assignable
	)
VALUES
	(
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	)
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description,
	api_key_assignable = excluded.api_key_assignable;
