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

create table polygon_points (
    polygon_id int references polygons(id) on delete cascade not null,
    point_id int references points(id) on delete cascade not null,
    primary key (polygon_id, point_id)
);


create table edges (
    id serial primary key not null unique,
    start_point_id int references points(id) on delete cascade not null,
    end_point_id int references points(id) on delete cascade not null,
    geom geometry(linestring, 4326)
);

create table directed_graphs (
    id serial primary key not null unique,
    name varchar(255)
);

create table graph_edges (
    graph_id int references directed_graphs(id) on delete cascade not null,
    edge_id int references edges(id) on delete cascade not null,
    primary key (graph_id, edge_id)
);
