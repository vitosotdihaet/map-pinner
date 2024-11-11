package controllers

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type LinePostgres struct {
	postgres *sqlx.DB
}

func NewLinePostgres(postgres *sqlx.DB) *LinePostgres {
	return &LinePostgres{postgres: postgres}
}

func (postgres *LinePostgres) Create(line entities.Line) (uint64, error) {
	postgisPoints := pointsToWKT(line.Points)

	query := fmt.Sprintf(
		"INSERT INTO %s (name, geometry) VALUES ($1, ST_SetSRID(ST_MakeLine(ARRAY[%s]), %v)) RETURNING id;",
		linesTable, strings.Join(postgisPoints, ", "), WGSSRID,
	)
	row := postgres.postgres.QueryRow(query, line.Name)

	var id uint64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (postgres *LinePostgres) GetAll() ([]entities.Line, error) {
	query := fmt.Sprintf(
		"SELECT id, name, ST_AsText(geometry) AS geometry FROM %s;", linesTable,
	)

	rows, err := postgres.postgres.Query(query)
	if err != nil {
		return nil, err
	}

	var lines []entities.Line
	for rows.Next() {
		var line entities.Line
		var wkt string
		if err := rows.Scan(&line.ID, &line.Name, &wkt); err != nil {
			return nil, err
		}
		line.Points = parseWKTLine(wkt)
		lines = append(lines, line)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func (postgres *LinePostgres) GetById(id uint64) (entities.Line, error) {
	query := fmt.Sprintf(
		"SELECT name, ST_AsText(geometry) AS geometry FROM %s WHERE id = $1;", linesTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	var line entities.Line
	line.ID = id

	var wkt string

	if err := row.Scan(&line.Name, &wkt); err != nil {
		return line, err
	}

	line.Points = parseWKTLine(wkt)

	return line, nil
}

func (postgres *LinePostgres) UpdateById(id uint64, lineUpdate entities.LineUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if lineUpdate.Points != nil {
		setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeLine(ARRAY[%s]), %d)", strings.Join(pointsToWKT(*lineUpdate.Points), ", "), WGSSRID))
	}

	if lineUpdate.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *lineUpdate.Name)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", linesTable, setQuery, argId)
	args = append(args, id)

	_, err := postgres.postgres.Exec(query, args...)

	return err
}

func (postgres *LinePostgres) DeleteById(id uint64) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1;", linesTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}
