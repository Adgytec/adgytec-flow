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
	($1, $2, $3, $4, $5)
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description,
	required_resources = excluded.required_resources;

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
	($1, $2, $3, $4, $5)
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description,
	required_resources = excluded.required_resources;

-- name: BatchAddManagementPermission :exec
WITH
	input_permissions AS (
		SELECT
			jsonb_array_elements(
				sqlc.arg (permissions)::JSONB
			) AS perm
	)
INSERT INTO
	management.permissions (
		key,
		service_name,
		name,
		description,
		required_resources
	)
SELECT
	perm ->> 'key',
	perm ->> 'service_name',
	perm ->> 'name',
	perm ->> 'description',
	(
		perm -> 'required_resources'
	)::global.permission_resource_type[]
FROM
	input_permissions
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description,
	required_resources = excluded.required_resources;

-- name: BatchAddApplicationPermission :exec
WITH
	input_permissions AS (
		SELECT
			jsonb_array_elements(
				sqlc.arg (permissions)::JSONB
			) AS perm
	)
INSERT INTO
	application.permissions (
		key,
		service_name,
		name,
		description,
		required_resources
	)
SELECT
	perm ->> 'key',
	perm ->> 'service_name',
	perm ->> 'name',
	perm ->> 'description',
	(
		perm -> 'required_resources'
	)::global.permission_resource_type[]
FROM
	input_permissions
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description,
	required_resources = excluded.required_resources;
