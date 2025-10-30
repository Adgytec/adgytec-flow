-- +goose Up
-- +goose StatementBegin
CREATE INDEX if NOT EXISTS idx_users_normalized_name ON global.users (
	normalized_name text_pattern_ops
);

CREATE INDEX if NOT EXISTS idx_users_normalized_email ON global.users (
	normalized_email text_pattern_ops
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX if EXISTS global.idx_users_normalized_email;

DROP INDEX if EXISTS global.idx_users_normalized_name;

-- +goose StatementEnd
