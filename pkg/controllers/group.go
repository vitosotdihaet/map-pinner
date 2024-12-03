package controllers

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type GroupPostgres struct {
	postgres *sqlx.DB
}

func NewGroupPostgres(postgres *sqlx.DB) *GroupPostgres {
	return &GroupPostgres{postgres: postgres}
}

func (postgres *GroupPostgres) Create(group entities.Group, authorId uint64) (uint64, error) {
	row := postgres.postgres.QueryRow("SELECT new_group($1, $2) AS group_id;", group.Name, authorId)

	var id uint64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (postgres *GroupPostgres) GetAll() ([]entities.Group, error) {
	query := fmt.Sprintf(
		"SELECT id, name FROM %s;", groupsTable,
	)
	rows, err := postgres.postgres.Query(query)

	if err != nil {
		return nil, err
	}

	var groups []entities.Group
	for rows.Next() {
		var group entities.Group
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (postgres *GroupPostgres) GetById(id uint64) (entities.Group, error) {
	query := fmt.Sprintf(
		"SELECT name FROM %s WHERE id = $1;", groupsTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	var group entities.Group
	group.ID = id
	if err := row.Scan(&group.Name); err != nil {
		return group, err
	}

	return group, nil
}

func (postgres *GroupPostgres) UpdateById(id uint64, groupUpdate entities.GroupUpdate) error {
	// setValues := make([]string, 0)
	// args := make([]interface{}, 0)
	// argId := 1

	// if groupUpdate.Name != nil {
	// 	setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
	// 	args = append(args, *groupUpdate.Name)
	// 	argId++
	// }

	// if groupUpdate.Latitude != nil && groupUpdate.Longitude != nil {
	// 	setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeGroup($%d, $%d), %d)", argId, argId+1, WGSSRID))
	// 	args = append(args, *groupUpdate.Longitude)
	// 	args = append(args, *groupUpdate.Latitude)
	// 	argId += 2
	// } else {
	// 	if groupUpdate.Latitude != nil {
	// 		setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeGroup(ST_X(geometry), $%d), %d)", argId, WGSSRID))
	// 		args = append(args, *groupUpdate.Latitude)
	// 		argId++
	// 	} else {
	// 		setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeGroup($%d, ST_Y(geometry)), %d)", argId, WGSSRID))
	// 		args = append(args, *groupUpdate.Longitude)
	// 		argId++

	// 	}
	// }

	// setQuery := strings.Join(setValues, ", ")

	// query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", groupsTable, setQuery, argId)
	// args = append(args, id)

	// _, err := postgres.postgres.Exec(query, args...)

	// return err
	return nil
}

func (postgres *GroupPostgres) DeleteById(id uint64) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1;", groupsTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}
