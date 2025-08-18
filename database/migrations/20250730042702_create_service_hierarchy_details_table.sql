-- +goose Up
-- +goose StatementBegin
CREATE TYPE global.service_hierarchy_type AS ENUM('level', 'tree');

CREATE TYPE global.service_hierarchy_result AS ENUM('hierarchy', 'item');

CREATE TABLE IF NOT EXISTS global.service_hierarchy_details (
	service_id UUID PRIMARY KEY NOT NULL REFERENCES global.services (id) ON DELETE CASCADE,
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

-- +goose StatementEnd
