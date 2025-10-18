CREATE TEMPORARY TABLE application_permission_staging (
	LIKE application.permissions including ALL
) ON
COMMIT
DROP;

CREATE TEMPORARY TABLE management_permission_staging (
	LIKE management.permissions including ALL
) ON
COMMIT
DROP;
