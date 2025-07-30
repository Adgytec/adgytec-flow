-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION global.permission_required_resources_update () returns trigger AS $$
begin
    new.required_resources := old.required_resources;
    return new;
end;
$$ language plpgsql;

CREATE OR REPLACE TRIGGER on_update_required_resources before
UPDATE ON management.permissions FOR each ROW WHEN (
	old.required_resources IS DISTINCT FROM new.required_resources
)
EXECUTE function global.permission_required_resources_update ();

CREATE OR REPLACE TRIGGER on_update_required_resources before
UPDATE ON application.permissions FOR each ROW WHEN (
	old.required_resources IS DISTINCT FROM new.required_resources
)
EXECUTE function global.permission_required_resources_update ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS on_update_required_resources ON application.permissions;

DROP TRIGGER if EXISTS on_update_required_resources ON management.permissions;

DROP FUNCTION if EXISTS global.permission_required_resources_update ();

-- +goose StatementEnd
