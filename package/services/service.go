package services

import (
	"github.com/vitosotdihaet/map-pinner/package/controllers"
	"github.com/vitosotdihaet/map-pinner/package/entities"
)

type Point interface {
	GetAll() ([]entities.Point, error)
	Create(point entities.Point) (int, error)
	GetById(id uint64) (entities.Point, error)
	UpdateById(newPoint entities.Point) error
	DeleteById(id uint64) error
}

type Polygon interface{
	GetAll() ([]entities.Polygon, error)
	Create(point entities.Polygon) (int, error)
	GetById(id uint64) (entities.Polygon, error)
	UpdateById(newPoint entities.Polygon) error
	DeleteById(id uint64) error
}

type Graph interface{}

type Service struct {
	Point
	Polygon
	Graph
}

func NewService(database *controllers.Database) *Service {
	return &Service{
		Point: NewPointService(database.Point),
		Polygon: NewPolygonService(database.Polygon),
	}
}