-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS global.user_social_links (
	id UUID PRIMARY KEY DEFAULT uuidv7 (),
	platform_name TEXT NOT NULL,
	profile_link TEXT NOT NULL,
	user_id UUID NOT NULL REFERENCES global.users (id) ON DELETE CASCADE,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ,
	UNIQUE (
		user_id,
		platform_name
	)
);

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON global.user_social_links FOR each ROW
EXECUTE function global.set_created_at ();

CREATE OR REPLACE TRIGGER on_update_set_updated_at before
UPDATE ON global.user_social_links FOR each ROW
EXECUTE function global.set_updated_at ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_at_update before
UPDATE ON global.user_social_links FOR each ROW WHEN (
	old.created_at IS DISTINCT FROM new.created_at
)
EXECUTE function global.created_at_update ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS on_update_set_updated_at ON global.user_social_links;

DROP TRIGGER if EXISTS on_insert_set_created_at ON global.user_social_links;

DROP TRIGGER if EXISTS on_update_prevent_created_at_update ON global.user_social_links;

DROP TABLE IF EXISTS global.user_social_links;

-- +goose StatementEnd
