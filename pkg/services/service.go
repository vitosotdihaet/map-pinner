package services

import (
	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)


type Point interface {
	GetAll() ([]entities.Point, error)
	Create(point entities.Point) (uint64, error)
	GetById(id uint64) (entities.Point, error)
	UpdateById(id uint64, pointUpdate entities.PointUpdate) error
	DeleteById(id uint64) error
}

type Polygon interface {
	GetAll() ([]entities.Polygon, error)
	Create(polygon entities.Polygon) (uint64, error)
	GetById(id uint64) (entities.Polygon, error)
	UpdateById(id uint64, newPolygon entities.PolygonUpdate) error
	DeleteById(id uint64) error
}

type Graph interface {
	GetAll() ([]entities.Graph, error)
	Create(graph entities.Graph) (uint64, error)
	GetById(id uint64) (entities.Graph, error)
	UpdateById(id uint64, graphUpdate entities.GraphUpdate) error
	DeleteById(id uint64) error
}

type Service struct {
	Point
	Polygon
	Graph
}

func NewService(database *controllers.Database) *Service {
	return &Service{
		Point: NewPointService(database.Point),
		Polygon: NewPolygonService(database.Polygon),
		Graph: NewGraphService(database.Graph),
	}
}