-- name: NewServiceStagingTable :exec
CREATE TEMPORARY TABLE services_staging (
	LIKE global.services including ALL
) ON
COMMIT
DROP;

-- name: AddServicesIntoStaging :copyfrom
INSERT INTO
	services_staging (
		id,
		name,
		description,
		type
	)
VALUES
	($1, $2, $3, $4);

-- name: UpsertServicesFromStaging :exec
INSERT INTO
	global.services AS s (
		id,
		name,
		description,
		type
	)
SELECT
	id,
	name,
	description,
	type
FROM
	services_staging
ON CONFLICT (id) DO UPDATE
SET
	description = excluded.description,
	type = excluded.type
WHERE
	s.description IS DISTINCT FROM excluded.description
	OR s.type IS DISTINCT FROM excluded.type;
