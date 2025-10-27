-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE TRIGGER archive_user_record before
UPDATE ON global.users FOR each ROW
EXECUTE function archive.archive_before_update ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS archive_user_record ON global.users;

-- +goose StatementEnd
