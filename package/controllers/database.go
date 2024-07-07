package controllers

import (
	"github.com/jmoiron/sqlx"
	"github.com/vitosotdihaet/map-pinner/package/entities"
)

type Point interface {
	GetAll() ([]entities.Point, error)
	Create(point entities.Point) (int, error)
	// GetById(id int) (entities.Point, error)
	// UpdateById(id int, newPoint entities.Point) error
	// DeleteById(id int) error
}

type Polygon interface{}

type Graph interface{}

type Database struct {
	Point
	Polygon
	Graph
}

func NewDatabase(postgres *sqlx.DB) *Database {
	return &Database{
		Point: NewPointPostgres(postgres),
	}
}