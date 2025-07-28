-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION global.created_by_update () returns trigger AS $$
begin
    new.created_by := old.created_by;
    return new;
end;
$$ language plpgsql;

CREATE OR REPLACE FUNCTION global.created_at_update () returns trigger AS $$
begin
    new.created_at := old.created_at;
    return new;
end;
$$ language plpgsql;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP FUNCTION if EXISTS global.created_by_update ();

DROP FUNCTION if EXISTS global.created_at_update ();

-- +goose StatementEnd
