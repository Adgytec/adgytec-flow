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
			perm ->> 'serviceId' AS service_id,
			perm ->> 'name' AS name,
			perm ->> 'description' AS description,
			ARRAY(
				SELECT
					jsonb_array_elements_text(
						perm -> 'requiredResources'
					)::management.permission_resource_type
			) AS required_resources
		FROM
			input_permissions
	)
INSERT INTO
	management.permissions (
		key,
		service_id,
		name,
		description,
		required_resources
	)
SELECT
	key,
	service_id,
	name,
	description,
	required_resources
FROM
	expanded_permissions
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description;

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
			perm ->> 'serviceId' AS service_id,
			perm ->> 'name' AS name,
			perm ->> 'description' AS description,
			ARRAY(
				SELECT
					jsonb_array_elements_text(
						perm -> 'requiredResources'
					)::application.permission_resource_type
			) AS required_resources
		FROM
			input_permissions
	)
INSERT INTO
	application.permissions (
		key,
		service_id,
		name,
		description,
		required_resources
	)
SELECT
	key,
	service_id,
	name,
	description,
	required_resources
FROM
	expanded_permissions
ON CONFLICT (key) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description;
