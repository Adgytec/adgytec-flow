-- name: NewServiceStagingTable :exec
CREATE TEMP TABLE services_staging (
	LIKE global.services including ALL
) ON
COMMIT delete rows;

-- name: AddServicesIntoStaging :copyfrom
INSERT INTO
	services_staging (id, name, type)
VALUES
	($1, $2, $3);

-- name: UpsertServicesFromStaging :exec
INSERT INTO
	global.services AS t (id, name, type)
SELECT
	id,
	name,
	type
FROM
	services_staging
ON CONFLICT (id) DO UPDATE
SET
	type = excluded.type
    where t.type is distinct from excluded.type;
