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
	),
	expanded_permissions AS (
		SELECT
			perm ->> 'key' AS key,
			perm ->> 'service_name' AS service_name,
			perm ->> 'name' AS name,
			perm ->> 'description' AS description,
			ARRAY(
				SELECT
					jsonb_array_elements_text(
						perm -> 'required_resources'
					)::global.permission_resource_type
			) AS required_resources
		FROM
			input_permissions
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
	key,
	service_name,
	name,
	description,
	required_resources
FROM
	expanded_permissions
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
	),
	expanded_permissions AS (
		SELECT
			perm ->> 'key' AS key,
			perm ->> 'service_name' AS service_name,
			perm ->> 'name' AS name,
			perm ->> 'description' AS description,
			ARRAY(
				SELECT
					jsonb_array_elements_text(
						perm -> 'required_resources'
					)::global.permission_resource_type
			) AS required_resources
		FROM
			input_permissions
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
	key,
	service_name,
	name,
	description,
	required_resources
FROM
	expanded_permissions
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description,
	required_resources = excluded.required_resources;
