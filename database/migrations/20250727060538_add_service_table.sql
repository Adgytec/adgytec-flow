-- +goose Up
-- +goose StatementBegin
CREATE TYPE global.service_logical_partition_type AS ENUM('hierarchy', 'none');

CREATE TABLE global.services (
	name TEXT PRIMARY KEY,
	permission_only BOOL NOT NULL DEFAULT FALSE,
	logical_partition global.service_logical_partition_type NOT NULL DEFAULT 'none',
	created_at TIMESTAMPTZ NOT NULL
);

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON global.services FOR each ROW
EXECUTE procedure global.set_created_at ();

CREATE OR REPLACE TRIGGER services_archive before delete ON global.services FOR each ROW
EXECUTE function archive.archive_before_delete ();

CREATE TYPE global.service_hierarchy_type AS ENUM('level', 'tree');

CREATE TYPE global.service_hierarchy_result AS ENUM('hierarchy', 'item');

CREATE TABLE global.service_hierarchy_details (
	service_name TEXT NOT NULL REFERENCES global.services (name) ON DELETE CASCADE ON UPDATE CASCADE,
	hierarchy_name TEXT NOT NULL,
	hierarchy_type global.service_hierarchy_type NOT NULL DEFAULT 'tree',
	hierarchy_result global.service_hierarchy_result NOT NULL DEFAULT 'item'
);

CREATE OR REPLACE TRIGGER service_hierarchy_details_archive before delete ON global.service_hierarchy_details FOR each ROW
EXECUTE function archive.archive_before_delete ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS service_hierarchy_details_archive ON global.service_hierarchy_details;

DROP TABLE global.service_hierarchy_details;

DROP TYPE if EXISTS global.service_hierarchy_result;

DROP TYPE if EXISTS global.service_hierarchy_type;

DROP TRIGGER if EXISTS services_archive ON global.services;

DROP TRIGGER if EXISTS on_insert_set_created_at ON global.services;

DROP TABLE global.services;

DROP TYPE if EXISTS global.service_logical_partition_type;

-- +goose StatementEnd
