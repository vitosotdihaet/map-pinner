package controllers

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)



type PolygonPostgres struct {
	postgres *sqlx.DB
}

func NewPolygonPostgres(postgres *sqlx.DB) *PolygonPostgres {
	return &PolygonPostgres{postgres: postgres}
}

func (postgres *PolygonPostgres) Create(polygon entities.Polygon) (uint64, error) {
	postgisPoints := pointsToWKT(polygon.Points)
	if len(postgisPoints) != 0 {
		postgisPoints = append(postgisPoints, postgisPoints[0])
	}
	logrus.Tracef("wkt polygon: %s", strings.Join(postgisPoints, ", "))


	query := fmt.Sprintf(
		"INSERT INTO %s (name, geom) VALUES ($1, ST_SetSRID(ST_MakePolygon(ST_MakeLine(ARRAY[%s])), %v)) RETURNING id;", 
		polygonsTable, strings.Join(postgisPoints, ", "), WGSSRID,
	)
	row := postgres.postgres.QueryRow(query, polygon.Name)

	var id uint64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (postgres *PolygonPostgres) GetAll() ([]entities.Polygon, error) {
	query := fmt.Sprintf(
		"SELECT id, name, ST_AsText(geom) AS geom FROM %s;", polygonsTable,
	)

	rows, err := postgres.postgres.Query(query)
	if err != nil {
		return nil, err
	}

	var polygons []entities.Polygon
	for rows.Next() {
		var polygon entities.Polygon
		var wkt string
		if err := rows.Scan(&polygon.ID, &polygon.Name, &wkt); err != nil {
			return nil, err
		}
		polygon.Points = parseWKTPolygon(wkt)
		polygons = append(polygons, polygon)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return polygons, nil
}

func (postgres *PolygonPostgres) GetById(id uint64) (entities.Polygon, error) {
	query := fmt.Sprintf(
		"SELECT name, ST_AsText(geom) AS geom FROM %s WHERE id = $1;", polygonsTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	var polygon entities.Polygon
	polygon.ID = id

	var wkt string

	if err := row.Scan(&polygon.Name, &wkt); err != nil {
		return polygon, err
	}

	polygon.Points = parseWKTPolygon(wkt)

	return polygon, nil
}

func (postgres *PolygonPostgres) UpdateById(id uint64, polygonUpdate entities.PolygonUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if polygonUpdate.Points != nil {
		wkt := pointsToWKT(*polygonUpdate.Points)
		if len(wkt) != 0 {
			wkt = append(wkt, wkt[0])
		}
		setValues = append(setValues, fmt.Sprintf("geom=ST_SetSRID(ST_MakePolygon(ST_MakeLine(ARRAY[%s])), %d)", strings.Join(wkt, ", "), WGSSRID))
	}

	if polygonUpdate.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *polygonUpdate.Name)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", polygonsTable, setQuery, argId)
	args = append(args, id)

	_, err := postgres.postgres.Exec(query, args...)

	return err
}

func (postgres *PolygonPostgres) DeleteById(id uint64) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1;", polygonsTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}