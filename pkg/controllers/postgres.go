package controllers

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	WGSSRID       = 4326
	pointsTable   = "markerspace.points"
	polygonsTable = "markerspace.polygons"
	linesTable    = "markerspace.lines"
	groupsTable   = "userspace.groups"
	usersTable    = "userspace.users"
	regionsTable  = "markerspace.regions"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode,
		),
	)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
