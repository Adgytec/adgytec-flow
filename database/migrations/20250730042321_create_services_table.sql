-- +goose Up
-- +goose StatementBegin
CREATE TYPE global.service_logical_partition_type AS ENUM('hierarchy', 'none');

/*
assignble tells if this permission can be assignable to organization
if the value is false it means it is normal service which can be part of organization or not 
*/
CREATE TABLE IF NOT EXISTS global.services (
	id UUID PRIMARY KEY,
	name TEXT NOT NULL UNIQUE,
	assignable BOOL NOT NULL DEFAULT FALSE,
	logical_partition global.service_logical_partition_type NOT NULL DEFAULT 'none',
	created_at TIMESTAMPTZ NOT NULL
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

DROP TYPE if EXISTS global.service_logical_partition_type;

-- +goose StatementEnd
