CREATE TEMPORARY TABLE service_restrictions_staging (
	LIKE global.service_restrictions including ALL
) ON
COMMIT
DROP;
