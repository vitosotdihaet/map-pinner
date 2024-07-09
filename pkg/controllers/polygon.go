package controllers

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)


func polygonToWKT(polygon entities.Polygon) []string {
	postgisPoints := make([]string, polygon.Length + 1)
	for i := range polygon.Length {
		point := polygon.Points[i]
		postgisPoints[i] = fmt.Sprintf("ST_MakePoint(%f, %f)", point.Latitude, point.Longitude)
	}
	postgisPoints[polygon.Length] = fmt.Sprintf("ST_MakePoint(%f, %f)", polygon.Points[0].Latitude, polygon.Points[0].Longitude)

	return postgisPoints
}

func parseWKT(wkt string) [20]*entities.Point {
	// Remove the "POLYGON((" prefix and "))" suffix
	wkt = strings.TrimPrefix(wkt, "POLYGON((")
	wkt = strings.TrimSuffix(wkt, "))")

	// Split the coordinates
	coordPairs := strings.Split(wkt, ",")

	var points [20]*entities.Point
	for i, pair := range coordPairs {
		coords := strings.Fields(pair)

		if len(coords) != 2 {
			logrus.Errorf("invalid wkt polygon parsing: %s", wkt)
			break
		}

		var point entities.Point
		fmt.Sscanf(coords[0], "%f", &point.Longitude)
		fmt.Sscanf(coords[1], "%f", &point.Latitude)
		points[i] = &point
	}

	return points
}


type PolygonPostgres struct {
	postgres *sqlx.DB
}

func NewPolygonPostgres(postgres *sqlx.DB) *PolygonPostgres {
	return &PolygonPostgres{postgres: postgres}
}

func (postgres *PolygonPostgres) Create(polygon entities.Polygon) (uint64, error) {
	postgisPoints := polygonToWKT(polygon)

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
		polygon.Points = parseWKT(wkt)
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

	polygon.Points = parseWKT(wkt)

	return polygon, nil
}

func (postgres *PolygonPostgres) UpdateById(newPolygon entities.Polygon) error {
	postgisPoints := polygonToWKT(newPolygon)

	query := fmt.Sprintf(
		"UPDATE %s SET name = $1, geom = ST_SetSRID(ST_MakePolygon(ST_MakeLine(ARRAY[%s])), %v) WHERE id = $4;",
		polygonsTable, strings.Join(postgisPoints, ", "), WGSSRID,
	)
	row := postgres.postgres.QueryRow(query, newPolygon.Name)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
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