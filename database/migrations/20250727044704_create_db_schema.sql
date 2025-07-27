-- +goose Up
-- +goose StatementBegin

create schema if not exists global;

create schema if not exists management;

create schema if not exists application;

create schema if not exists archive;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop schema if exists archive;

drop schema if exists application;

drop schema if exists management;

drop schema if exists global;

-- +goose StatementEnd
