package controllers

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type RolePostgres struct {
	postgres *sqlx.DB
}

func NewRolePostgres(postgres *sqlx.DB) *RolePostgres {
	return &RolePostgres{postgres: postgres}
}

func (postgres *RolePostgres) GetAllRoles() (map[uint64]string, error) {
	query := fmt.Sprintf("SELECT id, name FROM %s", rolesTable)

	rows, err := postgres.postgres.Query(query)
	if err != nil {
		return nil, err
	}

	roles := make(map[uint64]string)
	for rows.Next() {
		var id uint64
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		roles[id] = name
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (postgres *RolePostgres) HasAtLeastRoleInGroup(groupId uint64, userId uint64, role string) (bool, error) {
	query := fmt.Sprintf(
		`
		SELECT COUNT(*)
		FROM %s
		WHERE group_id = $1 AND user_id = $2 AND user_role_id IN (
			SELECT id
			FROM %s
			WHERE id <= (
				SELECT id
				FROM %s
				WHERE name = $3
			)
		)
		`, usersGroupsRelationTable, rolesTable, rolesTable,
	)

	var count int
	row := postgres.postgres.QueryRow(query, groupId, userId, role)
	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (postgres *RolePostgres) HasAtLeastRoleInRegion(regionId uint64, userId uint64, role string) (bool, error) {
	query := fmt.Sprintf(
		`
		SELECT COUNT(*)
		FROM %s
		WHERE group_id = (
			SELECT group_id
			FROM %s
			WHERE id = $1
		) AND user_id = $2 AND user_role_id IN (
			SELECT id
			FROM %s
			WHERE id <= (
				SELECT id
				FROM %s
				WHERE name = $3
			)
		)
		`, usersGroupsRelationTable, regionsTable, rolesTable, rolesTable,
	)

	var count int
	row := postgres.postgres.QueryRow(query, regionId, userId, role)
	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (postgres *RolePostgres) HasAtLeastRoleForMarker(markerType entities.MarkerType, markerId uint64, userId uint64, role string) (bool, error) {
	var markerTable string
	switch markerType {
	case entities.PointType:
		markerTable = pointsTable
	case entities.PolygonType:
		markerTable = polygonsTable
	case entities.LineType:
		markerTable = linesTable
	}

	query := fmt.Sprintf(
		`
		SELECT COUNT(*)
		FROM %s
		WHERE group_id = (
			SELECT group_id
			FROM %s
			WHERE id IN (
				SELECT region_id
				FROM %s
				WHERE id = $1
			)
		) AND user_id = $2 AND user_role_id IN (
			SELECT id
			FROM %s
			WHERE id <= (
				SELECT id
				FROM %s
				WHERE name = $3
			)
		)
		`, usersGroupsRelationTable, regionsTable, markerTable, rolesTable, rolesTable,
	)

	var count int
	row := postgres.postgres.QueryRow(query, markerId, userId, role)
	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (postgres *RolePostgres) ThereIsARoleWithId(roleId uint64) (bool, error) {
	query := fmt.Sprintf(
		`
		SELECT COUNT(*)
		FROM %s
		WHERE id = $1
		`, rolesTable,
	)

	var count int
	row := postgres.postgres.QueryRow(query, roleId)
	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (postgres *RolePostgres) ThereIsASystemRoleWithId(roleId uint64) (bool, error) {
	query := fmt.Sprintf(
		`
		SELECT COUNT(*)
		FROM %s
		WHERE id = $1
		`, systemRolesTable,
	)

	var count int
	row := postgres.postgres.QueryRow(query, roleId)
	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (postgres *RolePostgres) HasAtLeastSystemRole(userId uint64, role string) (bool, error) {
	query := fmt.Sprintf(
		`
		SELECT COUNT(*)
		FROM %s
		WHERE id = $1 AND user_role_id IN (
			SELECT id
			FROM %s
			WHERE id <= (
				SELECT id
				FROM %s
				WHERE name = $2
			)
		)
		`, usersTable, systemRolesTable, systemRolesTable,
	)

	var count int
	row := postgres.postgres.QueryRow(query, userId, role)
	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (postgres *RolePostgres) GetRoleID(userId uint64, groupId uint64) (uint64, error) {
	query := fmt.Sprintf(
		"SELECT user_role_id FROM %s WHERE user_id = $1 AND group_id = $2;", usersGroupsRelationTable,
	)

	var roleId uint64
	row := postgres.postgres.QueryRow(query, userId, groupId)
	if err := row.Scan(&roleId); err != nil {
		return 0, err
	}

	return roleId, nil
}
