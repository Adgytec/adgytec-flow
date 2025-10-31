-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS management.user_groups (
	id UUID PRIMARY KEY DEFAULT uuidv7 (),
	name TEXT NOT NULL,
	description TEXT,
	created_by UUID NOT NULL REFERENCES global.users (id) ON DELETE RESTRICT,
	created_at TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX user_group_name_unique_idx ON management.user_groups (lower(name));

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON management.user_groups FOR each ROW
EXECUTE function global.set_created_at ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_at_update before
UPDATE ON management.user_groups FOR each ROW WHEN (
	old.created_at IS DISTINCT FROM new.created_at
)
EXECUTE function global.created_at_update ();

CREATE OR REPLACE TRIGGER on_insert_set_created_by before insert ON management.user_groups FOR each ROW
EXECUTE function global.set_created_by ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_by_update before
UPDATE ON management.user_groups FOR each ROW WHEN (
	old.created_by IS DISTINCT FROM new.created_by
)
EXECUTE function global.created_by_update ();

CREATE OR REPLACE TRIGGER user_group_delete_archive before delete ON management.user_groups FOR each ROW
EXECUTE function archive.archive_before_delete ();

CREATE OR REPLACE TRIGGER user_group_update_archive before
UPDATE ON management.user_groups FOR each ROW WHEN (
	old.* IS DISTINCT FROM new.*
)
EXECUTE function archive.archive_before_update ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS user_group_delete_archive ON management.user_groups;

DROP TRIGGER if EXISTS user_group_update_archive ON management.user_groups;

DROP TRIGGER if EXISTS on_update_prevent_created_by_update ON management.user_groups;

DROP TRIGGER if EXISTS on_insert_set_created_by ON management.user_groups;

DROP TRIGGER if EXISTS on_update_prevent_created_at_update ON management.user_groups;

DROP TRIGGER if EXISTS on_insert_set_created_at ON management.user_groups;

DROP INDEX if EXISTS management.user_group_name_unique_idx;

DROP TABLE management.user_groups;

-- +goose StatementEnd
