-- +goose Up
-- +goose StatementBegin
-- created by are always actor type user
CREATE OR REPLACE FUNCTION global.set_created_by () returns trigger AS $$
declare
    actor text;
begin
    actor := current_setting('global.actor_id', true);

    -- Fail if user is not set
    IF actor IS NULL THEN
        RAISE EXCEPTION 'global.actor_id session variable must be set before INSERT/UPDATE';
    END IF;

    new.created_by = actor::uuid;
    return new;
end;
$$ language plpgsql;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP FUNCTION if EXISTS global.set_created_by;

-- +goose StatementEnd
