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

func (postgres *GroupPostgres) GetAll(userId uint64) ([]entities.Group, error) {
	query := fmt.Sprintf(
		`
		SELECT id, name
		FROM %s
		WHERE id IN (
			SELECT group_id
			FROM %s
			WHERE user_id = $1
		);
		`, groupsTable, usersGroupsRelationTable,
	)

	rows, err := postgres.postgres.Query(query, userId)
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

func (postgres *GroupPostgres) GetById(id uint64) (*entities.Group, error) {
	query := fmt.Sprintf(
		"SELECT name FROM %s WHERE id = $1;", groupsTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	var group entities.Group
	if err := row.Scan(&group.Name); err != nil {
		return nil, err
	}

	group.ID = id

	return &group, nil
}

func (postgres *GroupPostgres) GetAllUsers(id uint64) ([]entities.User, []string, error) {
	query := fmt.Sprintf(
		`
		SELECT 
			u.id AS id,
			u.name AS name,
			r.name AS role
		FROM 
			%s ugr
		JOIN
			%s u ON ugr.user_id = u.id
		JOIN
			%s r ON ugr.user_role_id = r.id
		JOIN
			%s g ON ugr.group_id = g.id
		WHERE 
			g.id = $1;
		`, usersGroupsRelationTable, usersTable, rolesTable, groupsTable,
	)

	rows, err := postgres.postgres.Query(query, id)
	if err != nil {
		return nil, nil, err
	}

	var users []entities.User
	var roles []string
	for rows.Next() {
		var user entities.User
		var role string
		if err := rows.Scan(&user.ID, &user.Name, &role); err != nil {
			return nil, nil, err
		}
		users = append(users, user)
		roles = append(roles, role)
	}

	return users, roles, nil
}

// func (postgres *GroupPostgres) UpdateById(id uint64, groupUpdate entities.GroupUpdate) error {
// 	// setValues := make([]string, 0)
// 	// args := make([]interface{}, 0)
// 	// argId := 1

// 	// if groupUpdate.Name != nil {
// 	// 	setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
// 	// 	args = append(args, *groupUpdate.Name)
// 	// 	argId++
// 	// }

// 	// if groupUpdate.Latitude != nil && groupUpdate.Longitude != nil {
// 	// 	setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeGroup($%d, $%d), %d)", argId, argId+1, WGSSRID))
// 	// 	args = append(args, *groupUpdate.Longitude)
// 	// 	args = append(args, *groupUpdate.Latitude)
// 	// 	argId += 2
// 	// } else {
// 	// 	if groupUpdate.Latitude != nil {
// 	// 		setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeGroup(ST_X(geometry), $%d), %d)", argId, WGSSRID))
// 	// 		args = append(args, *groupUpdate.Latitude)
// 	// 		argId++
// 	// 	} else {
// 	// 		setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeGroup($%d, ST_Y(geometry)), %d)", argId, WGSSRID))
// 	// 		args = append(args, *groupUpdate.Longitude)
// 	// 		argId++

// 	// 	}
// 	// }

// 	// setQuery := strings.Join(setValues, ", ")

// 	// query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", groupsTable, setQuery, argId)
// 	// args = append(args, id)

// 	// _, err := postgres.postgres.Exec(query, args...)

// 	// return err
// 	return nil
// }

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

func (postgres *GroupPostgres) AddUserToGroup(id uint64, userName string, roleId uint64) error {
	query := fmt.Sprintf(
		`
		INSERT INTO %s (group_id, user_id, user_role_id)
		VALUES ($1, (SELECT id FROM %s WHERE name = $2), $3)
		ON CONFLICT (group_id, user_id)
		DO UPDATE SET user_role_id = EXCLUDED.user_role_id;
		`, usersGroupsRelationTable, usersTable,
	)

	_, err := postgres.postgres.Exec(query, id, userName, roleId)
	if err != nil {
		return err
	}

	return nil
}
