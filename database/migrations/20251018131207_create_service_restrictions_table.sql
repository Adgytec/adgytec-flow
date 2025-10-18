-- +goose Up
-- +goose StatementBegin
/*
value type defines the data type for the restriction value
this table also helps during organization create for defining organization plan limits
*/
CREATE TABLE IF NOT EXISTS global.service_restrictions (
	id UUID PRIMARY KEY,
	service_id UUID NOT NULL REFERENCES global.services (id) ON DELETE CASCADE,
	name TEXT NOT NULL,
	description TEXT,
	value_type TEXT NOT NULL,
	UNIQUE (service_id, name)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE global.service_restrictions;

-- +goose StatementEnd
