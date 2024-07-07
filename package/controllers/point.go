package controllers

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/vitosotdihaet/map-pinner/package/entities"
)

type PointPostgres struct {
	postgres *sqlx.DB
}

func NewPointPostgres(postgres *sqlx.DB) *PointPostgres {
	return &PointPostgres{postgres: postgres}
}

func (postgres *PointPostgres) Create(point entities.Point) (int, error) {
	var id int

	query := fmt.Sprintf(
		"INSERT INTO %s (name, geom) VALUES ($1, ST_SetSRID(ST_MakePoint($2, $3), %v)) RETURNING id;",
		pointsTable, WGSSRID,
	)
	row := postgres.postgres.QueryRow(query, ""/* TODO: Add point names */, point.Longtitude, point.Lattitude)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (postgres *PointPostgres) GetAll() ([]entities.Point, error) {
	var points []entities.Point

	query := fmt.Sprintf(
		"SELECT id, ST_X(geom) AS longtitude, ST_Y(geom) AS lattitude FROM %s;", pointsTable,
	)
	rows, err := postgres.postgres.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var point entities.Point
		if err := rows.Scan(&point.ID, &point.Longtitude, &point.Lattitude); err != nil {
			return nil, err
		}
		logrus.Tracef("%v\n", point)
		points = append(points, point)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	rows.Close()

	return points, nil
}