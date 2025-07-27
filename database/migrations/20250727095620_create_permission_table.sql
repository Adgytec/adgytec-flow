-- +goose Up
-- +goose StatementBegin

create type global.permission_resource_type as enum ('project', 'logcial-partition', 'service-item');

create table management.permissions (
    key text primary key,
    service_name text not null references global.services(name),
    description text not null,
    required_resources global.permission_resource_type[] not null,
    created_at timestamptz not null,
    updated_at timestamptz not null
);

create or replace trigger on_insert_set_created_at 
before insert on management.permissions
for each row execute procedure global.set_created_at();

create or replace trigger on_insert_set_updated_at 
before insert on management.permissions
for each row execute procedure global.set_updated_at();

create or replace trigger on_update_set_updated_at 
before update on management.permissions
for each row execute procedure global.set_updated_at();

create or replace trigger on_update_prevent_created_at_update
before update on management.permissions
for each row execute procedure global.created_at_update();


create table application.permissions (
    key text primary key,
    service_name text not null references global.services(name),
    description text not null,
    required_resources global.permission_resource_type[] not null,
    created_at timestamptz not null,
    updated_at timestamptz not null
);

create or replace trigger on_insert_set_created_at 
before insert on application.permissions
for each row execute procedure global.set_created_at();

create or replace trigger on_insert_set_updated_at 
before insert on application.permissions
for each row execute procedure global.set_updated_at();

create or replace trigger on_update_set_updated_at 
before update on application.permissions
for each row execute procedure global.set_updated_at();

create or replace trigger on_update_prevent_created_at_update
before update on application.permissions
for each row execute procedure global.created_at_update();


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop trigger if exists on_update_set_updated_at on application.permissions;

drop trigger if exists on_insert_set_updated_at on application.permissons;

drop trigger if exists on_insert_set_created_at on application.permissons;

drop trigger if exists on_update_prevent_created_at_update on application.permissions;

drop table application.permissions;


drop trigger if exists on_update_set_updated_at on management.permissions;

drop trigger if exists on_insert_set_updated_at on management.permissons;

drop trigger if exists on_insert_set_created_at on management.permissons;

drop trigger if exists on_update_prevent_created_at_update on management.permissions;

drop table management.permissions;


drop type if exists global.permission_resource_type;

-- +goose StatementEnd
