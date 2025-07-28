-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA if NOT EXISTS global;

CREATE SCHEMA if NOT EXISTS management;

CREATE SCHEMA if NOT EXISTS application;

CREATE SCHEMA if NOT EXISTS archive;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP SCHEMA if EXISTS archive;

DROP SCHEMA if EXISTS application;

DROP SCHEMA if EXISTS management;

DROP SCHEMA if EXISTS global;

-- +goose StatementEnd
