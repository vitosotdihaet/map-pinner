package controllers

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)



type GraphPostgres struct {
	postgres *sqlx.DB
}

func NewGraphPostgres(postgres *sqlx.DB) *GraphPostgres {
	return &GraphPostgres{postgres: postgres}
}

func (postgres *GraphPostgres) Create(graph entities.Graph) (uint64, error) {
	postgisPoints := pointsToWKT(graph.Points)

	query := fmt.Sprintf(
		"INSERT INTO %s (name, geom) VALUES ($1, ST_SetSRID(ST_MakePolygon(ST_MakeLine(ARRAY[%s])), %v)) RETURNING id;", 
		graphsTable, strings.Join(postgisPoints, ", "), WGSSRID,
	)
	row := postgres.postgres.QueryRow(query, graph.Name)

	var id uint64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (postgres *GraphPostgres) GetAll() ([]entities.Graph, error) {
	query := fmt.Sprintf(
		"SELECT id, name, ST_AsText(geom) AS geom FROM %s;", graphsTable,
	)

	rows, err := postgres.postgres.Query(query)
	if err != nil {
		return nil, err
	}

	var graphs []entities.Graph
	for rows.Next() {
		var graph entities.Graph
		var wkt string
		if err := rows.Scan(&graph.ID, &graph.Name, &wkt); err != nil {
			return nil, err
		}
		graph.Points = parseWKT(wkt)
		graphs = append(graphs, graph)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return graphs, nil
}

func (postgres *GraphPostgres) GetById(id uint64) (entities.Graph, error) {
	query := fmt.Sprintf(
		"SELECT name, ST_AsText(geom) AS geom FROM %s WHERE id = $1;", graphsTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	var graph entities.Graph
	graph.ID = id

	var wkt string

	if err := row.Scan(&graph.Name, &wkt); err != nil {
		return graph, err
	}

	graph.Points = parseWKT(wkt)

	return graph, nil
}

func (postgres *GraphPostgres) UpdateById(newGraph entities.Graph) error {
	postgisPoints := pointsToWKT(newGraph.Points)

	query := fmt.Sprintf(
		"UPDATE %s SET name = $1, geom = ST_SetSRID(ST_MakePolygon(ST_MakeLine(ARRAY[%s])), %v) WHERE id = $4;",
		graphsTable, strings.Join(postgisPoints, ", "), WGSSRID,
	)
	row := postgres.postgres.QueryRow(query, newGraph.Name)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}

func (postgres *GraphPostgres) DeleteById(id uint64) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1;", graphsTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}