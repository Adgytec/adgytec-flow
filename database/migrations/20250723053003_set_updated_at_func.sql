-- +goose Up
-- +goose StatementBegin

create or replace function set_updated_columns()
returns trigger as $$
begin
  new.updated_at = clock_timestamp();
  return new;
end;
$$ language plpgsql;
  

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop function if exists set_updated_columns;

-- +goose StatementEnd
