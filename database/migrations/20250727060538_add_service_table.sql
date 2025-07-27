-- +goose Up
-- +goose StatementBegin

create type global.service_logical_partition_type as enum ('hierarchy', 'none');

create table global.services (
    name text primary key,
    permission_only bool not null default false,
    logical_partition global.service_logical_partition_type not null default 'none',
    created_at timestamptz not null
);

create or replace trigger on_insert_set_created_at 
before insert on global.services
for each row execute procedure global.set_created_at();

create type global.service_hierarchy_type as enum ('level', 'tree');

create type global.service_hierarchy_result as enum ('hierarchy', 'item');

create table global.service_hierarchy_details (
    service_name text not null references global.services(name),
    hierarchy_name text not null,
    hierarchy_type global.service_hierarchy_type not null default 'tree',
    hierarchy_result global.service_hierarchy_result not null default 'item'
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table global.service_hierarchy_details;

drop type if exists global.service_hierarchy_result;

drop type if exists global.service_hierarchy_type;

drop trigger if exists on_insert_set_created_at on global.services;

drop table global.services;

drop type if exists global.service_logical_partition_type;

-- +goose StatementEnd
