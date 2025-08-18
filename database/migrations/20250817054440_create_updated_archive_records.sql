-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS archive.updated_records (
	id UUID PRIMARY KEY DEFAULT global.uuid_generate_v7 (),
	table_name TEXT NOT NULL,
	old JSONB NOT NULL,
	new JSONB NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
	updated_by UUID NOT NULL REFERENCES global.users (id)
);

CREATE OR REPLACE FUNCTION archive.archive_after_update () returns trigger AS $$
declare
    actor text;
    table_full_name text;
begin
    actor := current_setting('global.user_id', true);
    table_full_name := TG_TABLE_SCHEMA || '.' || TG_TABLE_NAME;

    -- Fail if user is not set
    IF actor IS NULL THEN
        RAISE EXCEPTION 'global.user_id session variable must be set before INSERT/UPDATE';
    END IF;

    insert into archive.updated_records (table_name, old, new, updated_by)
    values (
        table_full_name,
        jsonb_object(old),
        jsonb_object(new),
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

-- +goose StatementEnd
