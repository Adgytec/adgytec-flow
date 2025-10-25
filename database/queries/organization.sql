-- name: NewOrganization :one
INSERT INTO
	global.organizations (
		root_user,
		name,
		description,
		logo,
		cover_image
	)
VALUES
	($1, $2, $3, $4, $5)
RETURNING
	id;

-- name: AddOrganizationRestrictions :copyfrom
INSERT INTO
	management.organization_service_restrictions (
		org_id,
		restriction_id,
		value,
		info
	)
VALUES
	($1, $2, $3, $4);
