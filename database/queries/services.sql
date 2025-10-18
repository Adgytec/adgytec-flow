-- name: AddServiceDetails :exec
INSERT INTO
	global.services (id, name, type)
VALUES
	($1, $2, $3)
ON CONFLICT (id) DO UPDATE
SET
	type = excluded.type;
