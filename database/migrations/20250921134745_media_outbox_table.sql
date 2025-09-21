-- +goose Up
-- +goose StatementBegin
CREATE TYPE global.media_outbox_status AS ENUM(
	'pending',
	'completed'
);

CREATE TABLE global.media_outbox (
	media_id UUID NOT NULL REFERENCES global.media (id),
	status global.media_outbox_status NOT NULL DEFAULT 'pending'
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS global.media_outbox;

DROP TRIGGER if EXISTS global.media_outbox_status;

-- +goose StatementEnd
