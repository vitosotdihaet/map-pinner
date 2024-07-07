package services

import (
	"github.com/vitosotdihaet/map-pinner/package/controllers"
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

type Service struct {
	Point
	Polygon
	Graph
}

func NewService(database *controllers.Database) *Service {
	return &Service{
		Point: NewPointService(database.Point),
	}
}