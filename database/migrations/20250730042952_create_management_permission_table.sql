-- +goose Up
-- +goose StatementBegin
CREATE TYPE global.assignable_actor_type AS ENUM(
	'api-key',
	'user',
	'all'
);

CREATE TYPE management.permission_resource_type AS ENUM('organization');

CREATE TABLE IF NOT EXISTS management.permissions (
	id UUID PRIMARY KEY,
	service_id UUID NOT NULL REFERENCES global.services (id) ON DELETE CASCADE,
	key TEXT UNIQUE NOT NULL,
	assignable_actor global.assignable_actor_type NOT NULL,
	required_resources management.permission_resource_type[] NOT NULL,
	name TEXT NOT NULL,
	description TEXT,
	created_at TIMESTAMPTZ NOT NULL clock_timestamp(),
	updated_at TIMESTAMPTZ
);

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON management.permissions FOR each ROW
EXECUTE function global.set_created_at ();

CREATE OR REPLACE TRIGGER on_update_set_updated_at before
UPDATE ON management.permissions FOR each ROW
EXECUTE function global.set_updated_at ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_at_update before
UPDATE ON management.permissions FOR each ROW WHEN (
	old.created_at IS DISTINCT FROM new.created_at
)
EXECUTE function global.created_at_update ();

CREATE OR REPLACE TRIGGER permissions_archive before delete ON management.permissions FOR each ROW
EXECUTE function archive.archive_before_delete ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS permission_archive ON mangement.permissions;

DROP TRIGGER if EXISTS on_update_set_updated_at ON management.permissions;

DROP TRIGGER if EXISTS on_insert_set_created_at ON management.permissions;

DROP TRIGGER if EXISTS on_update_prevent_created_at_update ON management.permissions;

DROP TABLE management.permissions;

DROP TYPE if EXISTS management.permission_resource_type;

DROP TYPE if EXISTS global.assignable_actor_type;

-- +goose StatementEnd
