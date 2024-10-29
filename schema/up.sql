-- enable postgis extension
CREATE extension IF NOT EXISTS postgis;

CREATE TABLE points (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    geom GEOMETRY(POINT) NOT NULL
);


CREATE TABLE polygons (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    geom GEOMETRY(POLYGON) NOT NULL
);

create table directed_graphs (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    geom GEOMETRY(LINESTRING) NOT NULL
);
