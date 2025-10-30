-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION if NOT EXISTS unaccent;

CREATE TYPE global.user_status AS ENUM(
	'disabled',
	'enabled'
);

CREATE TABLE IF NOT EXISTS global.users (
	id UUID PRIMARY KEY,
	email TEXT NOT NULL,
	normalized_email TEXT NOT NULL,
	name TEXT,
	normalized_name TEXT,
	profile_picture_id UUID REFERENCES global.media (id),
	about TEXT,
	date_of_birth date,
	status global.user_status NOT NULL DEFAULT 'enabled',
	created_at TIMESTAMPTZ NOT NULL,
	CHECK (
		name IS NULL
		OR normalized_name IS NOT NULL
	),
	CHECK (
		char_length(name) BETWEEN 3 AND 100
	),
	CHECK (
		char_length(about) BETWEEN 8 AND 1024
	)
);

CREATE UNIQUE INDEX global_users_email_unique_idx ON global.users (lower(email));

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON global.users FOR each ROW
EXECUTE function global.set_created_at ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_at_update before
UPDATE ON global.users FOR each ROW WHEN (
	old.created_at IS DISTINCT FROM new.created_at
)
EXECUTE function global.created_at_update ();

CREATE OR REPLACE FUNCTION global.normalize_user_fields () returns trigger AS $$ 
begin
    new.normalized_name := lower(unaccent(new.name)); 
    new.normalized_email := lower(unaccent(new.email)); 
    return new;
end; 
$$ language plpgsql;

CREATE OR REPLACE TRIGGER add_normalized_fields before insert
OR
UPDATE of name,
email ON global.users FOR each ROW
EXECUTE function global.normalize_user_fields ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS add_normalized_fields ON global.users;

DROP FUNCTION if EXISTS global.normalize_user_fields ();

DROP TRIGGER if EXISTS on_update_prevent_created_at_update ON global.users;

DROP TRIGGER if EXISTS on_insert_set_created_at ON global.users;

DROP INDEX if EXISTS global.global_users_email_unique_idx;

DROP TABLE IF EXISTS global.users;

DROP TYPE if EXISTS global.user_status;

DROP EXTENSION if EXISTS unaccent;

-- +goose StatementEnd
