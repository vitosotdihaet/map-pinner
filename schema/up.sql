-- enable postgis extension
create extension if not exists postgis;

-- srid 4326 for wgs 84

create table points (
    id serial primary key not null unique,
    name varchar(255) not null,
    geom geometry(point, 4326) not null
);


create table polygons (
    id serial primary key not null unique,
    name varchar(255) not null,
    geom geometry(polygon, 4326) not null
);

create table directed_graphs (
    id serial primary key not null unique,
    name varchar(255) not null,
    geom geometry(linestring, 4326) not null
);
