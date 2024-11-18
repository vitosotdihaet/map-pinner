-- https://dbdiagram.io/d/DB-course-project-6720daa8b4216d5a2898e693


-- enable postgis extension
CREATE extension IF NOT EXISTS postgis;


-- CREATE SCHEMA media;
CREATE SCHEMA rbac;
CREATE SCHEMA userspace;
CREATE SCHEMA markerspace;
CREATE TYPE markerspace.marker_type AS ENUM ('point', 'polygon', 'line');



CREATE TABLE userspace.users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL UNIQUE,
    password BYTEA NOT NULL
);



CREATE TABLE rbac.permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL UNIQUE
);

CREATE TABLE rbac.roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL UNIQUE
);

CREATE TABLE rbac.roles_permissions_relation (
    role_id INT NOT NULL REFERENCES rbac.roles(id),
    permission_id INT NOT NULL REFERENCES rbac.permissions(id),
    PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE rbac.system_roles_relation (
    user_id INT NOT NULL REFERENCES userspace.users(id),
    role_id INT NOT NULL REFERENCES rbac.roles(id),
    PRIMARY KEY (user_id, role_id)
);



CREATE TABLE userspace.groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE userspace.users_groups_relation (
    group_id INT NOT NULL REFERENCES userspace.groups(id),
    user_id INT NOT NULL REFERENCES userspace.users(id),
    user_role_id INT NOT NULL REFERENCES rbac.roles(id),
    PRIMARY KEY (group_id, user_id)
);

CREATE TABLE markerspace.regions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    group_id INT NOT NULL REFERENCES userspace.groups(id)
    -- top_left GEOMETRY(POINT, 4326) NOT NULL,
    -- bottom_right GEOMETRY(POINT, 4326) NOT NULL,
);



-- a markers table to link various marker types with regions
CREATE TABLE markerspace.markers (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    type markerspace.marker_type NOT NULL,
    markerable_id INT NOT NULL,
    region_id INT NOT NULL REFERENCES markerspace.regions(id)
);


-- a points table
CREATE TABLE markerspace.points (
    id SERIAL PRIMARY KEY,
    type markerspace.marker_type NOT NULL DEFAULT 'point' CHECK (type = 'point'),
    name VARCHAR(255) NOT NULL,
    geometry GEOMETRY(POINT, 4326) NOT NULL
    -- image_id INT REFERENCES media.images(id)
);

-- a polygons table
CREATE TABLE markerspace.polygons (
    id SERIAL PRIMARY KEY,
    type markerspace.marker_type NOT NULL DEFAULT 'polygon' CHECK (type = 'polygon'),
    name VARCHAR(255) NOT NULL,
    geometry GEOMETRY(POLYGON, 4326) NOT NULL
);

-- a lines table
CREATE TABLE markerspace.lines (
    id SERIAL PRIMARY KEY,
    type markerspace.marker_type NOT NULL DEFAULT 'line' CHECK (type = 'line'),
    name VARCHAR(255) NOT NULL,
    geometry GEOMETRY(LINESTRING, 4326) NOT NULL
);


-- foreign key constraints for polymorphic association in markers table
ALTER TABLE markerspace.markers
    ADD CONSTRAINT fk_point_marker
    FOREIGN KEY (markerable_id)
    REFERENCES markerspace.points(id) ON DELETE CASCADE;

ALTER TABLE markerspace.markers
    ADD CONSTRAINT fk_polygon_marker
    FOREIGN KEY (markerable_id)
    REFERENCES markerspace.polygons(id) ON DELETE CASCADE;

ALTER TABLE markerspace.markers
    ADD CONSTRAINT fk_line_marker
    FOREIGN KEY (markerable_id)
    REFERENCES markerspace.lines(id) ON DELETE CASCADE;

