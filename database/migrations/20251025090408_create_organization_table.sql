-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS global.organizations (
	id UUID PRIMARY KEY DEFAULT uuidv7 (),
	root_user UUID NOT NULL REFERENCES global.users (id),
	name TEXT NOT NULL,
	description TEXT,
	logo UUID REFERENCES global.media (id),
	cover_media UUID REFERENCES global.media (id),
	created_at TIMESTAMPTZ NOT NULL,
	created_by UUID NOT NULL REFERENCES global.users (id)
);

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON global.organizations FOR each ROW
EXECUTE function global.set_created_at ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_at_update before
UPDATE ON global.organizations FOR each ROW WHEN (
	old.created_at IS DISTINCT FROM new.created_at
)
EXECUTE function global.created_at_update ();

CREATE OR REPLACE TRIGGER on_insert_set_created_by before insert ON global.organizations FOR each ROW
EXECUTE function global.set_created_by ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_by_update before
UPDATE ON global.organizations FOR each ROW WHEN (
	old.created_by IS DISTINCT FROM new.created_by
)
EXECUTE function global.created_by_update ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS on_update_prevent_created_by_update ON global.organizations;

DROP TRIGGER if EXISTS on_insert_set_created_by ON global.organizations;

DROP TRIGGER if EXISTS on_update_prevent_created_at_update ON global.organizations;

DROP TRIGGER if EXISTS on_insert_set_created_at ON global.organizations;

DROP TABLE global.organizations;

-- +goose StatementEnd
