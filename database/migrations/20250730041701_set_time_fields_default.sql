-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION global.set_updated_at () returns trigger AS $$
begin
  new.updated_at = clock_timestamp();
  return new;
end;
$$ language plpgsql;

CREATE OR REPLACE FUNCTION global.set_created_at () returns trigger AS $$
begin
    new.created_at = clock_timestamp();
    return new;
end;
$$ language plpgsql;

CREATE OR REPLACE FUNCTION global.set_expires_at () returns trigger AS $$
begin
    IF TG_NARGS > 0 THEN
        new.expires_at = clock_timestamp() + (TG_ARGV[0])::interval;
    ELSE
        new.expires_at = clock_timestamp() + interval '24 hours';
    END IF;
    return new;
end;
$$ language plpgsql;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP FUNCTION if EXISTS global.set_expires_at;

DROP FUNCTION if EXISTS global.set_created_at;

DROP FUNCTION if EXISTS global.set_updated_at;

-- +goose StatementEnd
