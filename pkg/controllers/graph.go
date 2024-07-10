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
	if len(postgisPoints) > 0 {
		postgisPoints = postgisPoints[:len(postgisPoints) - 1]
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (name, geom) VALUES ($1, ST_SetSRID(ST_MakeLine(ARRAY[%s]), %v)) RETURNING id;", 
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

func (postgres *GraphPostgres) UpdateById(id uint64, graphUpdate entities.GraphUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if graphUpdate.Points != nil {
		wkt := pointsToWKT(*graphUpdate.Points)
		if len(wkt) > 0 {
			wkt = wkt[:len(wkt) - 1]
		}
		setValues = append(setValues, fmt.Sprintf("geom=ST_SetSRID(ST_MakePolygon(ST_MakeLine(ARRAY[%s])), %d)", strings.Join(wkt, ", "), WGSSRID))
	}

	if graphUpdate.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *graphUpdate.Name)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", graphsTable, setQuery, argId)
	args = append(args, id)

	_, err := postgres.postgres.Exec(query, args...)

	return err
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