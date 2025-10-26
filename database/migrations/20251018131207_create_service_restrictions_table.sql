-- +goose Up
-- +goose StatementBegin
/*
this table also helps during organization create for defining organization plan limits
*/
CREATE TABLE IF NOT EXISTS global.service_restrictions (
	id UUID PRIMARY KEY,
	service_id UUID NOT NULL REFERENCES global.services (id) ON DELETE CASCADE,
	name TEXT NOT NULL,
	description TEXT,
	UNIQUE (service_id, name)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE global.service_restrictions;

-- +goose StatementEnd
