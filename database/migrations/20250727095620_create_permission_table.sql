-- +goose Up
-- +goose StatementBegin
CREATE TYPE global.permission_resource_type AS ENUM(
	'project',
	'logcial-partition',
	'service-item'
);

CREATE TABLE management.permissions (
	key TEXT PRIMARY KEY,
	service_name TEXT NOT NULL REFERENCES global.services (name) ON DELETE CASCADE ON UPDATE CASCADE,
	description TEXT NOT NULL,
	required_resources global.permission_resource_type[] NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL
);

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON management.permissions FOR each ROW
EXECUTE procedure global.set_created_at ();

CREATE OR REPLACE TRIGGER on_insert_set_updated_at before insert ON management.permissions FOR each ROW
EXECUTE procedure global.set_updated_at ();

CREATE OR REPLACE TRIGGER on_update_set_updated_at before
UPDATE ON management.permissions FOR each ROW
EXECUTE procedure global.set_updated_at ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_at_update before
UPDATE ON management.permissions FOR each ROW
EXECUTE procedure global.created_at_update ();

CREATE OR REPLACE TRIGGER permissions_archive before delete ON management.permissions FOR each ROW
EXECUTE function archive.archive_before_delete ();

CREATE TABLE application.permissions (
	key TEXT PRIMARY KEY,
	service_name TEXT NOT NULL REFERENCES global.services (name) ON DELETE CASCADE ON UPDATE CASCADE,
	description TEXT NOT NULL,
	required_resources global.permission_resource_type[] NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL
);

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON application.permissions FOR each ROW
EXECUTE procedure global.set_created_at ();

CREATE OR REPLACE TRIGGER on_insert_set_updated_at before insert ON application.permissions FOR each ROW
EXECUTE procedure global.set_updated_at ();

CREATE OR REPLACE TRIGGER on_update_set_updated_at before
UPDATE ON application.permissions FOR each ROW
EXECUTE procedure global.set_updated_at ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_at_update before
UPDATE ON application.permissions FOR each ROW
EXECUTE procedure global.created_at_update ();

CREATE OR REPLACE TRIGGER permissions_archive before delete ON application.permissions FOR each ROW
EXECUTE function archive.archive_before_delete ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS permission_archive ON application.permissions;

DROP TRIGGER if EXISTS on_update_set_updated_at ON application.permissions;

DROP TRIGGER if EXISTS on_insert_set_updated_at ON application.permissons;

DROP TRIGGER if EXISTS on_insert_set_created_at ON application.permissons;

DROP TRIGGER if EXISTS on_update_prevent_created_at_update ON application.permissions;

DROP TABLE application.permissions;

DROP TRIGGER if EXISTS permission_archive ON mangement.permissions;

DROP TRIGGER if EXISTS on_update_set_updated_at ON management.permissions;

DROP TRIGGER if EXISTS on_insert_set_updated_at ON management.permissons;

DROP TRIGGER if EXISTS on_insert_set_created_at ON management.permissons;

DROP TRIGGER if EXISTS on_update_prevent_created_at_update ON management.permissions;

DROP TABLE management.permissions;

DROP TYPE if EXISTS global.permission_resource_type;

-- +goose StatementEnd
