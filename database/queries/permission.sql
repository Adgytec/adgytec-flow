-- name: NewManagementPermissionStagingTable :exec
CREATE TEMPORARY TABLE management_permission_staging (
	LIKE management.permissions including ALL
) ON
COMMIT
DROP;

-- name: AddManagementPermissionsIntoStaging :copyfrom
INSERT INTO
	management_permission_staging (
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
	);

-- name: UpsertManagementPermissionsFromStaging :exec
INSERT INTO
	management.permissions AS p (
		id,
		service_id,
		key,
		name,
		description,
		required_resources,
		assignable_actor
	)
SELECT
	id,
	service_id,
	key,
	name,
	description,
	required_resources,
	assignable_actor
FROM
	management_permission_staging
ON CONFLICT (id) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description,
	assignable_actor = excluded.assignable_actor,
	required_resources = excluded.required_resources
WHERE
	p.name IS DISTINCT FROM excluded.name
	OR p.description IS DISTINCT FROM excluded.description
	OR p.assignable_actor IS DISTINCT FROM excluded.assignable_actor
	OR p.required_resources IS DISTINCT FROM excluded.required_resources;

-- name: NewApplicationPermissionStagingTable :exec
CREATE TEMPORARY TABLE application_permission_staging (
	LIKE application.permissions including ALL
) ON
COMMIT
DROP;

-- name: AddApplicationPermissionsIntoStaging :copyfrom
INSERT INTO
	application_permission_staging (
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
	);

-- name: UpsertApplicationPermissionsFromStaging :exec
INSERT INTO
	application.permissions AS p (
		id,
		service_id,
		key,
		name,
		description,
		required_resources,
		assignable_actor
	)
SELECT
	id,
	service_id,
	key,
	name,
	description,
	required_resources,
	assignable_actor
FROM
	application_permission_staging
ON CONFLICT (id) DO UPDATE
SET
	name = excluded.name,
	description = excluded.description,
	assignable_actor = excluded.assignable_actor,
	required_resources = excluded.required_resources
WHERE
	p.name IS DISTINCT FROM excluded.name
	OR p.description IS DISTINCT FROM excluded.description
	OR p.assignable_actor IS DISTINCT FROM excluded.assignable_actor
	OR p.required_resources IS DISTINCT FROM excluded.required_resources;
