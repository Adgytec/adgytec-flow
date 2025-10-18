CREATE TEMPORARY TABLE services_staging (
	LIKE global.services including ALL
) ON
COMMIT
DROP;
