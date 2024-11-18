package controllers

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type RegionPostgres struct {
	postgres *sqlx.DB
}

func NewRegionPostgres(postgres *sqlx.DB) *RegionPostgres {
	return &RegionPostgres{postgres: postgres}
}

func (postgres *RegionPostgres) Create(region entities.Region) (uint64, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (name) VALUES ($1) RETURNING id;",
		regionsTable,
	)
	row := postgres.postgres.QueryRow(query, region.Name)

	var id uint64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (postgres *RegionPostgres) GetAll() ([]entities.Region, error) {
	query := fmt.Sprintf(
		"SELECT id, name FROM %s;", regionsTable,
	)
	rows, err := postgres.postgres.Query(query)

	if err != nil {
		return nil, err
	}

	var regions []entities.Region
	for rows.Next() {
		var region entities.Region
		if err := rows.Scan(&region.ID, &region.Name); err != nil {
			return nil, err
		}
		regions = append(regions, region)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return regions, nil
}

func (postgres *RegionPostgres) GetById(id uint64) (entities.Region, error) {
	query := fmt.Sprintf(
		"SELECT name FROM %s WHERE id = $1;", regionsTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	var region entities.Region
	region.ID = id
	if err := row.Scan(&region.Name); err != nil {
		return region, err
	}

	return region, nil
}

func (postgres *RegionPostgres) UpdateById(id uint64, regionUpdate entities.RegionUpdate) error {
	// setValues := make([]string, 0)
	// args := make([]interface{}, 0)
	// argId := 1

	// if regionUpdate.Name != nil {
	// 	setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
	// 	args = append(args, *regionUpdate.Name)
	// 	argId++
	// }

	// if regionUpdate.Latitude != nil && regionUpdate.Longitude != nil {
	// 	setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeRegion($%d, $%d), %d)", argId, argId+1, WGSSRID))
	// 	args = append(args, *regionUpdate.Longitude)
	// 	args = append(args, *regionUpdate.Latitude)
	// 	argId += 2
	// } else {
	// 	if regionUpdate.Latitude != nil {
	// 		setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeRegion(ST_X(geometry), $%d), %d)", argId, WGSSRID))
	// 		args = append(args, *regionUpdate.Latitude)
	// 		argId++
	// 	} else {
	// 		setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeRegion($%d, ST_Y(geometry)), %d)", argId, WGSSRID))
	// 		args = append(args, *regionUpdate.Longitude)
	// 		argId++

	// 	}
	// }

	// setQuery := strings.Join(setValues, ", ")

	// query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", regionsTable, setQuery, argId)
	// args = append(args, id)

	// _, err := postgres.postgres.Exec(query, args...)

	// return err
	return nil
}

func (postgres *RegionPostgres) DeleteById(id uint64) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1;", regionsTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}
