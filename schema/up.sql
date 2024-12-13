-- enable postgis extension
CREATE extension IF NOT EXISTS postgis;


CREATE SCHEMA rbac;
CREATE SCHEMA userspace;
CREATE SCHEMA markerspace;

CREATE TYPE markerspace.marker_type AS ENUM ('point', 'polygon', 'line');



CREATE TABLE rbac.roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL UNIQUE
);

CREATE TABLE rbac.system_roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL UNIQUE
);



CREATE TABLE userspace.users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    system_role_id INT NOT NULL REFERENCES rbac.system_roles(id)
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

CREATE TABLE userspace.regions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    group_id INT NOT NULL REFERENCES userspace.groups(id)
);



-- a points table
CREATE TABLE markerspace.points (
    id SERIAL PRIMARY KEY,
    type markerspace.marker_type NOT NULL DEFAULT 'point' CHECK (type = 'point'),
    name VARCHAR(255) NOT NULL,
    geometry GEOMETRY(POINT, 4326) NOT NULL,
    region_id INT NOT NULL REFERENCES userspace.regions(id)
);

-- a polygons table
CREATE TABLE markerspace.polygons (
    id SERIAL PRIMARY KEY,
    type markerspace.marker_type NOT NULL DEFAULT 'polygon' CHECK (type = 'polygon'),
    name VARCHAR(255) NOT NULL,
    geometry GEOMETRY(POLYGON, 4326) NOT NULL,
    region_id INT NOT NULL REFERENCES userspace.regions(id)
);

-- a lines table
CREATE TABLE markerspace.lines (
    id SERIAL PRIMARY KEY,
    type markerspace.marker_type NOT NULL DEFAULT 'line' CHECK (type = 'line'),
    name VARCHAR(255) NOT NULL,
    geometry GEOMETRY(LINESTRING, 4326) NOT NULL,
    region_id INT NOT NULL REFERENCES userspace.regions(id)
);



CREATE FUNCTION new_user(user_name VARCHAR(32), password_hash VARCHAR)
RETURNS INT AS $$
DECLARE
    user_id INT;
BEGIN
    INSERT INTO userspace.users (name, password)
    VALUES (user_name, password_hash)
    RETURNING id INTO user_id;
    RETURN user_id;
END;
$$ LANGUAGE plpgsql;

-- create a new group
CREATE FUNCTION new_group(group_name VARCHAR(255), user_id INT)
RETURNS INT AS $$
DECLARE
    group_id INT;
BEGIN
    INSERT INTO userspace.groups (name)
    VALUES (group_name)
    RETURNING id INTO group_id;
    
    INSERT INTO userspace.users_groups_relation (group_id, user_id)
    VALUES (group_id, user_id);

    RETURN group_id;
END;
$$ LANGUAGE plpgsql;


-- make user an owner
CREATE PROCEDURE make_owner(user_id INT)
AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM userspace.users WHERE id = user_id) THEN
        UPDATE userspace.users
        SET system_role_id = (SELECT id FROM rbac.system_roles WHERE name = 'owner')
        WHERE id = user_id;
    END IF;
END;
$$ LANGUAGE plpgsql;


CREATE FUNCTION assign_default_roles_users()
RETURNS TRIGGER AS $$
BEGIN
    NEW.system_role_id := (SELECT id FROM rbac.system_roles WHERE name = 'user');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION assign_default_roles_groups()
RETURNS TRIGGER AS $$
BEGIN
    NEW.user_role_id := (SELECT id FROM rbac.roles WHERE name = 'admin');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER default_role_users
BEFORE INSERT ON userspace.users
FOR EACH ROW
EXECUTE FUNCTION assign_default_roles_users();

CREATE TRIGGER default_role_groups
BEFORE INSERT ON userspace.users_groups_relation
FOR EACH ROW
EXECUTE FUNCTION assign_default_roles_groups();


INSERT INTO rbac.roles (name) VALUES ('admin');
INSERT INTO rbac.roles (name) VALUES ('editor');
INSERT INTO rbac.roles (name) VALUES ('viewer');

INSERT INTO rbac.system_roles (name) VALUES ('owner');
INSERT INTO rbac.system_roles (name) VALUES ('user');
