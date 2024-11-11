-- enable postgis extension
CREATE extension IF NOT EXISTS postgis;

-- https://dbdiagram.io/d/DB-course-project-6720daa8b4216d5a2898e693


CREATE SCHEMA markerspace;

CREATE TYPE markerspace.marker_type AS ENUM ('point', 'polygon', 'line');

-- a markers table to link various marker types
CREATE TABLE markerspace.markers (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    type markerspace.marker_type NOT NULL,
    markerable_id INT NOT NULL
    -- map_id UUID NOT NULL REFERENCES markerspace.maps(id) ON DELETE CASCADE
);


-- a points table
CREATE TABLE markerspace.points (
    id SERIAL PRIMARY KEY,
    type markerspace.marker_type NOT NULL DEFAULT 'point' CHECK (type = 'point'),
    name VARCHAR(255) NOT NULL,
    geometry GEOMETRY(POINT) NOT NULL
    -- image_id INT REFERENCES media.images(id)
);

-- a polygons table
CREATE TABLE markerspace.polygons (
    id SERIAL PRIMARY KEY,
    type markerspace.marker_type NOT NULL DEFAULT 'polygon' CHECK (type = 'polygon'),
    name VARCHAR(255) NOT NULL,
    geometry GEOMETRY(POLYGON) NOT NULL
);

-- a lines table
CREATE TABLE markerspace.lines (
    id SERIAL PRIMARY KEY,
    type markerspace.marker_type NOT NULL DEFAULT 'line' CHECK (type = 'line'),
    name VARCHAR(255) NOT NULL,
    geometry GEOMETRY(LINESTRING) NOT NULL
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

