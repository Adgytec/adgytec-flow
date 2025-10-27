-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS management.users (
	id UUID PRIMARY KEY REFERENCES global.users (id) ON DELETE CASCADE,
	created_by UUID NOT NULL REFERENCES global.users (id) ON DELETE RESTRICT,
	created_at TIMESTAMPTZ NOT NULL
);

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON management.users FOR each ROW
EXECUTE function global.set_created_at ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_at_update before
UPDATE ON management.users FOR each ROW WHEN (
	old.created_at IS DISTINCT FROM new.created_at
)
EXECUTE function global.created_at_update ();

CREATE OR REPLACE TRIGGER on_insert_set_created_by before insert ON management.users FOR each ROW
EXECUTE function global.set_created_by ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_by_update before
UPDATE ON management.users FOR each ROW WHEN (
	old.created_by IS DISTINCT FROM new.created_by
)
EXECUTE function global.created_by_update ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS on_update_prevent_created_by_update ON management.users;

DROP TRIGGER if EXISTS on_insert_set_created_by ON management.users;

DROP TRIGGER if EXISTS on_update_prevent_created_at_update ON management.users;

DROP TRIGGER if EXISTS on_insert_set_created_at ON management.users;

DROP TABLE management.users;

-- +goose StatementEnd
