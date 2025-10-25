-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS global.organizations (
	id UUID PRIMARY KEY DEFAULT uuidv7 (),
	root_user UUID NOT NULL REFERENCES global.users (id),
	name TEXT NOT NULL,
	description TEXT,
	logo UUID REFERENCES global.media (id),
	cover_media UUID REFERENCES global.media (id),
	created_at TIMESTAMPTZ NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE global.organizations;

-- +goose StatementEnd
