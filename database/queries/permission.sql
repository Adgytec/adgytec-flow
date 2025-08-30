-- name: AddManagementPermission :exec
INSERT INTO
	management.permissions (
		id,
		service_id,
		key,
		name,
		description,
		required_resources,
		assignable_actor
	)
VALUES
	(
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	)
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description,
	assignable_actor = excluded.assignable_actor;

-- name: AddApplicationPermission :exec
INSERT INTO
	application.permissions (
		id,
		service_id,
		key,
		name,
		description,
		required_resources,
		assignable_actor
	)
VALUES
	(
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	)
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description,
	assignable_actor = excluded.assignable_actor;
