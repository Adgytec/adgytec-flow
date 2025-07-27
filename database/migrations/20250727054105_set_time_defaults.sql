-- +goose Up
-- +goose StatementBegin

create or replace function global.set_updated_at()
returns trigger as $$
begin
  new.updated_at = clock_timestamp();
  return new;
end;
$$ language plpgsql;

create or replace function global.set_created_at()
returns trigger as $$
begin
    new.created_at = clock_timestamp();
    return new;
end;
$$ language plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop function if exists global.set_created_at;

drop function if exists global.set_updated_at;

-- +goose StatementEnd
