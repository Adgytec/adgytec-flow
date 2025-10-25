-- +goose Up
-- +goose StatementBegin
/*
value -1 defines no restriction limit
*/
CREATE TABLE IF NOT EXISTS management.organization_service_restrictions (
	org_id UUID NOT NULL REFERENCES global.organizations (id),
	restriction_id UUID NOT NULL REFERENCES global.service_restrictions (id),
	value SMALLINT NOT NULL,
	PRIMARY KEY (
		org_id,
		restriction_id
	),
	CHECK (value >= -1)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE management.organization_service_restrictions;

-- +goose StatementEnd
