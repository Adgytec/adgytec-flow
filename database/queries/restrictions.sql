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
		description,
		value_type
	)
VALUES
	($1, $2, $3, $4, $5);

-- name: UpsertServiceRestrictionsFromStaging :exec
INSERT INTO
	global.service_restrictions AS s (
		id,
		service_id,
		name,
		description,
		value_type
	)
SELECT
	id,
	service_id,
	name,
	description,
	value_type
FROM
	service_restrictions_staging
ON CONFLICT (id) DO UPDATE
SET
	description = excluded.description,
	value_type = excluded.value_type
WHERE
	s.description IS DISTINCT FROM excluded.description
	OR s.value_type IS DISTINCT FROM excluded.value_type;
