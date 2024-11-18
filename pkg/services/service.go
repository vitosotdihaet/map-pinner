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
	UpdateById(id uint64, polygonUpdate entities.PolygonUpdate) error
	DeleteById(id uint64) error
}

type Line interface {
	GetAll() ([]entities.Line, error)
	Create(line entities.Line) (uint64, error)
	GetById(id uint64) (entities.Line, error)
	UpdateById(id uint64, lineUpdate entities.LineUpdate) error
	DeleteById(id uint64) error
}

// type Marker interface {
// 	GetAll() ([]entities.Line, error)
// 	Create(line interface {entities.Marker}) (uint64, error)
// 	GetById(id uint64) (entities.Line, error)
// 	UpdateById(id uint64, lineUpdate entities.LineUpdate) error
// 	DeleteById(id uint64) error
// }

type User interface {
	GetAll() ([]entities.User, error)
	Create(user entities.User, password entities.HashedPassword) (uint64, error)
	GetById(id uint64) (entities.User, error)
	GetByNamePassword(user entities.User, password entities.HashedPassword) (*entities.User, error)
	// UpdateById(id uint64, lineUpdate entities.GroupUpdate) error
	DeleteById(id uint64) error
}

type Group interface {
	GetAll() ([]entities.Group, error)
	Create(group entities.Group) (uint64, error)
	GetById(id uint64) (entities.Group, error)
	UpdateById(id uint64, groupUpdate entities.GroupUpdate) error
	DeleteById(id uint64) error
}

type Region interface {
	GetAll() ([]entities.Region, error)
	Create(region entities.Region) (uint64, error)
	GetById(id uint64) (entities.Region, error)
	UpdateById(id uint64, regionUpdate entities.RegionUpdate) error
	DeleteById(id uint64) error
}

type Service struct {
	Point
	Polygon
	Line
	User
	Group
	Region
}

func NewService(database *controllers.Database) *Service {
	return &Service{
		Point:   NewPointService(database.Point),
		Polygon: NewPolygonService(database.Polygon),
		Line:    NewLineService(database.Line),
		User:    NewUserService(database.User),
		Group:   NewGroupService(database.Group),
		Region:  NewRegionService(database.Region),
	}
}
