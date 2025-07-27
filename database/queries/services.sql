-- name: AddService :exec
insert into global.services (name, permission_only, logical_partition)
values ($1, $2, $3);

-- name: AddServiceHierarchyDetails :exec
insert into global.service_hierarchy_details (service_name, hierarchy_name, hierarchy_type, hierarchy_result)
values ($1, $2, $3, $4);
