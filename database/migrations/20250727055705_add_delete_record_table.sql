-- +goose Up
-- +goose StatementBegin

create table archive.deleted_records (
    table_name text not null,
    record jsonb not null,
    deleted_at timestamptz not null default clock_timestamp()
);

create or replace function archive.archive_before_delete()
returns trigger as $$
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

drop function if exists archive.archive_before_delete; 

drop table archive.deleted_records;

-- +goose StatementEnd
