-- +goose Up
-- +goose StatementBegin
CREATE TYPE global.actor_type AS ENUM(
	'api-key',
	'user',
	'signed',
	'system'
);

CREATE TABLE IF NOT EXISTS archive.updated_records (
	id UUID PRIMARY KEY DEFAULT uuidv7 (),
	table_name TEXT NOT NULL,
	old JSONB NOT NULL,
	new JSONB NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by_type global.actor_type NOT NULL,
	updated_by UUID NOT NULL
);

CREATE OR REPLACE FUNCTION archive.archive_after_update () returns trigger AS $$
declare
    actor text;
    actor_type text;
    table_full_name text;
begin
    actor := current_setting('global.actor_id', true);
    actor_type := current_setting('global.actor_type', true);
    table_full_name := TG_TABLE_SCHEMA || '.' || TG_TABLE_NAME;

    -- Fail if user is not set
    IF actor IS NULL OR actor_type IS NULL THEN
        RAISE EXCEPTION 'global.actor_id and global.actor_type session variable must be set before INSERT/UPDATE';
    END IF;

    insert into archive.updated_records (table_name, old, new, updated_by_type, updated_by)
    values (
        table_full_name,
        to_jsonb(old),
        to_jsonb(new),
        actor_type::global.actor_type,
        actor::uuid
    );
    return null;
end;
$$ language plpgsql;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP FUNCTION if EXISTS archive.archive_after_update;

DROP TABLE archive.updated_records;

DROP TYPE if EXISTS global.actor_type;

-- +goose StatementEnd
