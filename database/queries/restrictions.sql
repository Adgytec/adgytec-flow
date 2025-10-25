-- name: NewServiceRestrictionsStagingTable :exec
CREATE TEMPORARY TABLE service_restrictions_staging (
	LIKE global.service_restrictions including ALL
) ON
COMMIT
DROP;

-- name: AddServiceRestrictionIntoStaging :copyfrom
INSERT INTO
	service_restrictions_staging (
		id,
		service_id,
		name,
		description
	)
VALUES
	($1, $2, $3, $4);

-- name: UpsertServiceRestrictionsFromStaging :exec
INSERT INTO
	global.service_restrictions AS s (
		id,
		service_id,
		name,
		description
	)
SELECT
	id,
	service_id,
	name,
	description
FROM
	service_restrictions_staging
ON CONFLICT (id) DO UPDATE
SET
	description = excluded.description
WHERE
	s.description IS DISTINCT FROM excluded.description;

-- name: GetCoreServiceRestrictions :many
SELECT
	r.id,
	r.name,
	s.name AS service_name
FROM
	global.service_restrictions r
	JOIN global.services s ON r.service_id = s.id
WHERE
	s.type = 'core';
