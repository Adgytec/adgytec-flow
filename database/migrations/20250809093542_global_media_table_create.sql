-- +goose Up
-- +goose StatementBegin
CREATE TYPE global.media_type AS ENUM(
	'image',
	'video',
	'other'
);

CREATE TYPE global.media_status AS ENUM(
	'processing',
	'failed',
	'completed'
);

CREATE TABLE IF NOT EXISTS global.media (
	id UUID PRIMARY KEY DEFAULT global.uuid_generate_v7 (),
	bucket_path TEXT NOT NULL UNIQUE,
	size BIGINT NOT NULL CHECK (size > 0),
	media_type global.media_type NOT NULL,
	content_type TEXT,
	status global.media_status NOT NULL DEFAULT 'processing',
	created_at TIMESTAMPTZ NOT NULL
);

CREATE OR REPLACE TRIGGER on_insert_set_created_at before insert ON global.media FOR each ROW
EXECUTE function global.set_created_at ();

CREATE OR REPLACE TRIGGER on_update_prevent_created_at_update before
UPDATE ON global.media FOR each ROW WHEN (
	old.created_at IS DISTINCT FROM new.created_at
)
EXECUTE function global.created_at_update ();

CREATE TABLE IF NOT EXISTS global.media_video (
	media_id UUID NOT NULL REFERENCES global.media (id) ON DELETE CASCADE,
	thumbnail TEXT,
	adaptive_manifest TEXT,
	preview TEXT
);

CREATE TABLE IF NOT EXISTS global.media_image (
	media_id UUID NOT NULL REFERENCES global.media (id) ON DELETE CASCADE,
	thumbnail TEXT,
	small TEXT,
	medium TEXT,
	large TEXT,
	extra_large TEXT
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS global.media_image;

DROP TABLE IF EXISTS global.media_video;

DROP TRIGGER if EXISTS on_update_prevent_created_at_update ON global.media;

DROP TRIGGER if EXISTS on_insert_set_created_at ON global.media;

DROP TABLE IF EXISTS global.media;

DROP TYPE if EXISTS global.media_status;

DROP TYPE if EXISTS global.media_type;

-- +goose StatementEnd
