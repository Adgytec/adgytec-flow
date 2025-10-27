-- +goose Up
-- +goose StatementBegin
/*
core: services that are part of all organizations and global managmenet scope
optional: org scoped services that can be added to any organization
platform: global scope services that are not org scoped and used for global management
*/
CREATE TYPE global.service_type AS ENUM(
	'core',
	'optional',
	'platform'
);

CREATE TABLE IF NOT EXISTS global.services (
	id UUID PRIMARY KEY,
	name TEXT NOT NULL UNIQUE,
	description TEXT,
	type global.service_type NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON global.services FOR each ROW
EXECUTE function global.set_created_at ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_at_update before
UPDATE ON global.services FOR each ROW WHEN (
	old.created_at IS DISTINCT FROM new.created_at
)
EXECUTE function global.created_at_update ();

CREATE OR REPLACE TRIGGER services_archive before delete ON global.services FOR each ROW
EXECUTE function archive.archive_before_delete ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS services_archive ON global.services;

DROP TRIGGER if EXISTS on_update_pervent_created_at_update ON global.services;

DROP TRIGGER if EXISTS on_insert_set_created_at ON global.services;

DROP TABLE global.services;

DROP TYPE if EXISTS global.service_type;

-- +goose StatementEnd
