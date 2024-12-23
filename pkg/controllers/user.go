package controllers

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type UserPostgres struct {
	postgres *sqlx.DB
}

func NewUserPostgres(postgres *sqlx.DB) *UserPostgres {
	return &UserPostgres{postgres: postgres}
}

func (postgres *UserPostgres) Create(user entities.User, password entities.HashedPassword) (uint64, error) {
	row := postgres.postgres.QueryRow("SELECT new_user($1, $2) AS user_id;", user.Name, password.Value)

	var id uint64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (postgres *UserPostgres) GetAll() ([]entities.User, error) {
	query := fmt.Sprintf(
		"SELECT id, name FROM %s;", usersTable,
	)
	rows, err := postgres.postgres.Query(query)

	if err != nil {
		return nil, err
	}

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (postgres *UserPostgres) GetById(id uint64) (entities.User, error) {
	query := fmt.Sprintf(
		"SELECT name FROM %s WHERE id = $1;", usersTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	var user entities.User
	user.ID = id
	if err := row.Scan(&user.Name); err != nil {
		return user, err
	}

	return user, nil
}

func (postgres *UserPostgres) GetByName(user entities.User) (*entities.User, entities.HashedPassword, error) {
	var hashedPassword entities.HashedPassword
	query := fmt.Sprintf(
		"SELECT id, name, password FROM %s WHERE name = $1;", usersTable,
	)
	row := postgres.postgres.QueryRow(query, user.Name)
	err := row.Scan(&user.ID, &user.Name, &hashedPassword.Value)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entities.HashedPassword{}, nil
		}
		return nil, entities.HashedPassword{}, err
	}

	return &user, hashedPassword, nil
}

func (postgres *UserPostgres) ExistsWithName(userName string) (bool, error) {
	query := fmt.Sprintf(
		"SELECT COUNT(*) FROM %s WHERE name = $1;", usersTable,
	)

	var count int
	row := postgres.postgres.QueryRow(query, userName)
	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

// func (postgres *UserPostgres) UpdateById(id uint64, userUpdate entities.UserUpdate) error {
// 	// setValues := make([]string, 0)
// 	// args := make([]interface{}, 0)
// 	// argId := 1

// 	// if userUpdate.Nickname != nil {
// 	// 	setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
// 	// 	args = append(args, *userUpdate.Nickname)
// 	// 	argId++
// 	// }

// 	// if userUpdate.Latitude != nil && userUpdate.Longitude != nil {
// 	// 	setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeUser($%d, $%d), %d)", argId, argId+1, WGSSRID))
// 	// 	args = append(args, *userUpdate.Longitude)
// 	// 	args = append(args, *userUpdate.Latitude)
// 	// 	argId += 2
// 	// } else {
// 	// 	if userUpdate.Latitude != nil {
// 	// 		setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeUser(ST_X(geometry), $%d), %d)", argId, WGSSRID))
// 	// 		args = append(args, *userUpdate.Latitude)
// 	// 		argId++
// 	// 	} else {
// 	// 		setValues = append(setValues, fmt.Sprintf("geometry=ST_SetSRID(ST_MakeUser($%d, ST_Y(geometry)), %d)", argId, WGSSRID))
// 	// 		args = append(args, *userUpdate.Longitude)
// 	// 		argId++

// 	// 	}
// 	// }

// 	// setQuery := strings.Join(setValues, ", ")

// 	// query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", usersTable, setQuery, argId)
// 	// args = append(args, id)

// 	// _, err := postgres.postgres.Exec(query, args...)

// 	// return err
// 	return nil
// }

func (postgres *UserPostgres) DeleteById(id uint64) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1;", usersTable,
	)
	row := postgres.postgres.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}
