-- +goose Up
-- +goose StatementBegin

create or replace function global.created_by_update()
returns trigger as $$
begin
    new.created_by := old.created_by;
    return new;
end;
$$
language plpgsql;


create or replace function global.created_at_update()
returns trigger as $$
begin
    new.created_at := old.created_at;
    return new;
end;
$$
language plpgsql;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop function if exists global.created_by_update();

drop function if exists global.created_at_update();

-- +goose StatementEnd
