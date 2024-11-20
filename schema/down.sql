DROP TABLE IF EXISTS markerspace.lines;
DROP TABLE IF EXISTS markerspace.polygons;
DROP TABLE IF EXISTS markerspace.points;
DROP TABLE IF EXISTS markerspace.regions;

DROP TABLE IF EXISTS userspace.users_groups_relation;
DROP TABLE IF EXISTS userspace.groups;

DROP TABLE IF EXISTS rbac.system_roles_relation;
DROP TABLE IF EXISTS rbac.roles_permissions_relation;
DROP TABLE IF EXISTS rbac.roles;
DROP TABLE IF EXISTS rbac.permissions;

DROP TABLE IF EXISTS userspace.users;

DROP TYPE IF EXISTS markerspace.marker_type;

DROP SCHEMA IF EXISTS markerspace CASCADE;
DROP SCHEMA IF EXISTS userspace CASCADE;
DROP SCHEMA IF EXISTS rbac CASCADE;

DROP EXTENSION IF EXISTS postgis;
