-- +goose Up
-- +goose StatementBegin
CREATE TABLE archive.deleted_records (
	table_name TEXT NOT NULL,
	record JSONB NOT NULL,
	deleted_at TIMESTAMPTZ NOT NULL DEFAULT CLOCK_TIMESTAMP()
);

CREATE OR REPLACE FUNCTION archive.archive_before_delete () returns trigger AS $$
declare
    table_full_name text;
begin
    table_full_name := TG_TABLE_SCHEMA || '.' || TG_TABLE_NAME;

    insert into archive.deleted_records (table_name, record)
    values (
        table_full_name,
        jsonb_object(old)
    );
    return old;
end;
$$ language plpgsql;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP FUNCTION if EXISTS archive.archive_before_delete;

DROP TABLE archive.deleted_records;

-- +goose StatementEnd
