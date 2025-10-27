-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS management.user_group_users (
	user_group_id UUID NOT NULL REFERENCES management.user_groups (id) ON DELETE CASCADE,
	user_id UUID NOT NULL REFERENCES management.users (id) ON DELETE CASCADE,
	created_by UUID NOT NULL REFERENCES management.users (id) ON DELETE RESTRICT,
	created_at TIMESTAMPTZ NOT NULL,
	PRIMARY KEY (
		user_group_id,
		user_id
	)
);

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON management.user_group_users FOR each ROW
EXECUTE function global.set_created_at ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_at_update before
UPDATE ON management.user_group_users FOR each ROW WHEN (
	old.created_at IS DISTINCT FROM new.created_at
)
EXECUTE function global.created_at_update ();

CREATE OR REPLACE TRIGGER on_insert_set_created_by before insert ON management.user_group_users FOR each ROW
EXECUTE function global.set_created_by ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_by_update before
UPDATE ON management.user_group_users FOR each ROW WHEN (
	old.created_by IS DISTINCT FROM new.created_by
)
EXECUTE function global.created_by_update ();

CREATE OR REPLACE TRIGGER user_group_users_delete_archive before delete ON management.user_group_users FOR each ROW
EXECUTE function archive.archive_before_delete ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS user_group_users_delete_archive ON management.user_group_users;

DROP TRIGGER if EXISTS on_update_prevent_created_by_update ON management.user_group_users;

DROP TRIGGER if EXISTS on_insert_set_created_by ON management.user_group_users;

DROP TRIGGER if EXISTS on_update_prevent_created_at_update ON management.user_group_users;

DROP TRIGGER if EXISTS on_insert_set_created_at ON management.user_group_users;

DROP TABLE management.user_group_users;

-- +goose StatementEnd
