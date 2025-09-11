-- +goose Up
-- +goose StatementBegin
CREATE TYPE global.media_upload_type AS ENUM(
	'singlepart',
	'multipart'
);

CREATE TABLE IF NOT EXISTS global.temporary_media (
	id UUID PRIMARY KEY DEFAULT global.uuid_generate_v7 (),
	bucket_path TEXT NOT NULL,
	size BIGINT NOT NULL CHECK (size > 0),
	upload_type global.media_upload_type NOT NULL,
	upload_id TEXT,
	content_type TEXT
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS global.temporary_media;

DROP TYPE if EXISTS global.media_upload_type;

-- +goose StatementEnd
