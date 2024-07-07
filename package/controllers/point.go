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
	query := fmt.Sprintf(
		"INSERT INTO %s (name, geom) VALUES ($1, ST_SetSRID(ST_MakePoint($2, $3), %v)) RETURNING id;",
		pointsTable, WGSSRID,
	)
	row := postgres.postgres.QueryRow(query, point.Name, point.Longitude, point.Lattitude)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (postgres *PointPostgres) GetAll() ([]entities.Point, error) {
	query := fmt.Sprintf(
		"SELECT id, name, ST_X(geom) AS longtitude, ST_Y(geom) AS lattitude FROM %s;", pointsTable,
	)
	rows, err := postgres.postgres.Query(query)

	if err != nil {
		return nil, err
	}

	var points []entities.Point
	for rows.Next() {
		var point entities.Point
		if err := rows.Scan(&point.ID, &point.Name, &point.Longitude, &point.Lattitude); err != nil {
			return nil, err
		}
		logrus.Tracef("%v\n", point)
		points = append(points, point)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return points, nil
}

func (postgres *PointPostgres) GetById(id uint64) (entities.Point, error) {
	query := fmt.Sprintf(
		"SELECT name, ST_X(geom) AS longtitude, ST_Y(geom) AS lattitude FROM %s WHERE id = $1;", pointsTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	var point entities.Point
	point.ID = id
	if err := row.Scan(&point.Name, &point.Longitude, &point.Lattitude); err != nil {
		return point, err
	}

	return point, nil
}

func (postgres *PointPostgres) UpdateById(newPoint entities.Point) error {
	query := fmt.Sprintf(
		"UPDATE %s SET name = $1, geom = ST_SetSRID(ST_MakePoint($2, $3), %v) WHERE id = $4;",
		pointsTable, WGSSRID,
	)
	row := postgres.postgres.QueryRow(query, newPoint.Name, newPoint.Longitude, newPoint.Lattitude, newPoint.ID)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}

func (postgres *PointPostgres) DeleteById(id uint64) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1;", pointsTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}