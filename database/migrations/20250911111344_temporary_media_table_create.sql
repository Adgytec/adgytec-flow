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
	content_type TEXT,
	expires_at TIMESTAMPTZ NOT NULL
);

CREATE OR REPLACE TRIGGER on_insert_set_expires_at before insert ON global.temporary_media FOR each ROW
EXECUTE function global.set_expires_at ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TRIGGER if EXISTS on_insert_set_expires_at ON global.temporary_media;

DROP TABLE IF EXISTS global.temporary_media;

DROP TYPE if EXISTS global.media_upload_type;

-- +goose StatementEnd
