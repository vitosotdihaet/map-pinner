-- enable postgis extension
CREATE extension IF NOT EXISTS postgis;

CREATE TABLE points (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    geom GEOMETRY(POINT) NOT NULL
);


CREATE TABLE polygons (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    geom GEOMETRY(POLYGON) NOT NULL
);

CREATE TABLE directed_graphs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    geom GEOMETRY(LINESTRING) NOT NULL
);
