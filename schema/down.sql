ALTER TABLE userspace.users DISABLE TRIGGER default_role_users;
ALTER TABLE userspace.users_groups_relation DISABLE TRIGGER default_role_groups;

DROP TRIGGER IF EXISTS default_role_users ON userspace.users;
DROP TRIGGER IF EXISTS default_role_groups ON userspace.users_groups_relation;

DROP FUNCTION IF EXISTS assign_default_roles_users;
DROP FUNCTION IF EXISTS assign_default_roles_groups;
DROP PROCEDURE IF EXISTS make_owner;
DROP FUNCTION IF EXISTS new_group;
DROP FUNCTION IF EXISTS new_user;

DROP TABLE IF EXISTS markerspace.lines;
DROP TABLE IF EXISTS markerspace.polygons;
DROP TABLE IF EXISTS markerspace.points;
DROP TABLE IF EXISTS userspace.regions;
DROP TABLE IF EXISTS userspace.users_groups_relation;
DROP TABLE IF EXISTS userspace.groups;
DROP TABLE IF EXISTS userspace.users;
DROP TABLE IF EXISTS rbac.system_roles;
DROP TABLE IF EXISTS rbac.roles;

DROP TYPE IF EXISTS markerspace.marker_type;

DROP SCHEMA IF EXISTS markerspace CASCADE;
DROP SCHEMA IF EXISTS userspace CASCADE;
DROP SCHEMA IF EXISTS rbac CASCADE;

DROP EXTENSION IF EXISTS postgis;
