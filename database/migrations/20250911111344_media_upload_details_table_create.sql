-- +goose Up
-- +goose StatementBegin
CREATE TYPE global.media_upload_type AS ENUM(
	'singlepart',
	'multipart'
);

CREATE TABLE IF NOT EXISTS global.media_upload_details (
	media_id UUID NOT NULL REFERENCES global.media (id),
	upload_type global.media_upload_type NOT NULL,
	upload_id TEXT,
	CHECK (
		(
			upload_type = 'multipart'
			AND upload_id IS NOT NULL
		)
		OR (
			upload_type <> 'multipart'
			AND upload_id IS NULL
		)
	)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS global.media_upload_details;

DROP TYPE if EXISTS global.media_upload_type;

-- +goose StatementEnd
