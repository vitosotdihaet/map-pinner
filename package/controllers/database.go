package controllers

import (
	"github.com/jmoiron/sqlx"
	"github.com/vitosotdihaet/map-pinner/package/entities"
)

type Point interface {
	GetAll() ([]entities.Point, error)
	Create(point entities.Point) (int, error)
	GetById(id uint64) (entities.Point, error)
	UpdateById(newPoint entities.Point) error
	DeleteById(id uint64) error
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