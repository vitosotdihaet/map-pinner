package controllers

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type PointPostgres struct {
	postgres *sqlx.DB
}

func NewPointPostgres(postgres *sqlx.DB) *PointPostgres {
	return &PointPostgres{postgres: postgres}
}

func (postgres *PointPostgres) Create(point entities.Point) (uint64, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (name, geometry) VALUES ($1, ST_SetSRID(ST_MakePoint($2, $3), %v)) RETURNING id;",
		pointsTable, WGSSRID,
	)
	row := postgres.postgres.QueryRow(query, point.Name, point.Longitude, point.Latitude)

	var id uint64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (postgres *PointPostgres) GetAll() ([]entities.Point, error) {
	query := fmt.Sprintf(
		"SELECT id, name, ST_X(geometry) AS longtitude, ST_Y(geometry) AS lattitude FROM %s;", pointsTable,
	)
	rows, err := postgres.postgres.Query(query)

	if err != nil {
		return nil, err
	}

	var points []entities.Point
	for rows.Next() {
		var point entities.Point
		if err := rows.Scan(&point.ID, &point.Name, &point.Longitude, &point.Latitude); err != nil {
			return nil, err
		}
		points = append(points, point)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return points, nil
}

func (postgres *PointPostgres) GetById(id uint64) (entities.Point, error) {
	query := fmt.Sprintf(
		"SELECT name, ST_X(geometry) AS longtitude, ST_Y(geometry) AS lattitude FROM %s WHERE id = $1;", pointsTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	var point entities.Point
	point.ID = id
	if err := row.Scan(&point.Name, &point.Longitude, &point.Latitude); err != nil {
		return point, err
	}

	return point, nil
}

func (postgres *PointPostgres) UpdateById(id uint64, pointUpdate entities.PointUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if pointUpdate.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *pointUpdate.Name)
		argId++
	}

	if pointUpdate.Latitude != nil && pointUpdate.Longitude != nil {
		setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakePoint($%d, $%d), %d)", argId, argId+1, WGSSRID))
		args = append(args, *pointUpdate.Longitude)
		args = append(args, *pointUpdate.Latitude)
		argId += 2
	} else {
		if pointUpdate.Latitude != nil {
			setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakePoint(ST_X(geometry), $%d), %d)", argId, WGSSRID))
			args = append(args, *pointUpdate.Latitude)
			argId++
		} else {
			setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakePoint($%d, ST_Y(geometry)), %d)", argId, WGSSRID))
			args = append(args, *pointUpdate.Longitude)
			argId++

		}
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", pointsTable, setQuery, argId)
	args = append(args, id)

	_, err := postgres.postgres.Exec(query, args...)

	return err
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
