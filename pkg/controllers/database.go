package controllers

import (
	"github.com/jmoiron/sqlx"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)


type Point interface {
	GetAll() ([]entities.Point, error)
	Create(point entities.Point) (uint64, error)
	GetById(id uint64) (entities.Point, error)
	UpdateById(id uint64, pointUpdate entities.PointUpdate) error
	DeleteById(id uint64) error
}

type Polygon interface{
	GetAll() ([]entities.Polygon, error)
	Create(polygon entities.Polygon) (uint64, error)
	GetById(id uint64) (entities.Polygon, error)
	UpdateById(newPolygon entities.Polygon) error
	DeleteById(id uint64) error
}

type Graph Polygon

type Database struct {
	Point
	Polygon
	Graph
}

func NewDatabase(postgres *sqlx.DB) *Database {
	return &Database{
		Point: NewPointPostgres(postgres),
		Polygon: NewPolygonPostgres(postgres),
		Graph: NewPolygonPostgres(postgres),
	}
}