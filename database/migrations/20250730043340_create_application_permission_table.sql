-- +goose Up
-- +goose StatementBegin
CREATE TYPE application.permission_resource_type AS ENUM(
	'project',
	'logical-partition',
	'service-item'
);

CREATE TABLE IF NOT EXISTS application.permissions (
	id UUID PRIMARY KEY,
	service_id UUID NOT NULL REFERENCES global.services (id) ON DELETE CASCADE,
	key TEXT UNIQUE NOT NULL,
	assignable_actor global.assignable_actor_type NOT NULL,
	required_resources application.permission_resource_type[] NOT NULL,
	name TEXT NOT NULL,
	description TEXT,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ
);

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON application.permissions FOR each ROW
EXECUTE function global.set_created_at ();

CREATE OR REPLACE TRIGGER on_update_set_updated_at before
UPDATE ON application.permissions FOR each ROW
EXECUTE function global.set_updated_at ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_at_update before
UPDATE ON application.permissions FOR each ROW WHEN (
	old.created_at IS DISTINCT FROM new.created_at
)
EXECUTE function global.created_at_update ();

CREATE OR REPLACE TRIGGER permissions_archive before delete ON application.permissions FOR each ROW
EXECUTE function archive.archive_before_delete ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS permission_archive ON application.permissions;

DROP TRIGGER if EXISTS on_update_set_updated_at ON application.permissions;

DROP TRIGGER if EXISTS on_insert_set_created_at ON application.permissions;

DROP TRIGGER if EXISTS on_update_prevent_created_at_update ON application.permissions;

DROP TABLE application.permissions;

DROP TYPE if EXISTS application.permission_resource_type;

-- +goose StatementEnd
