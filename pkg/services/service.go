package services

import (
	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type Marker interface {
	GetAll(regionId uint64) ([]entities.Marker, error)
	Create(regionId uint64, marker entities.Marker) (uint64, error)
	GetById(markerType entities.MarkerType, id uint64) (entities.Marker, error)
	UpdateById(id uint64, markerUpdate entities.Marker) error
	DeleteById(markerType entities.MarkerType, id uint64) error
}

type User interface {
	GetAll() ([]entities.User, error)
	Create(user entities.User, password entities.HashedPassword) (uint64, error)
	GetById(id uint64) (entities.User, error)
	GetByNamePassword(user entities.User, password entities.HashedPassword) (*entities.User, error)
	// UpdateById(id uint64, lineUpdate entities.GroupUpdate) error
	DeleteById(id uint64) error
}

type Group interface {
	GetAll(userId uint64) ([]entities.Group, error)
	Create(group entities.Group, authorId uint64) (uint64, error)
	GetById(id uint64) (entities.Group, error)
	UpdateById(id uint64, groupUpdate entities.GroupUpdate) error
	DeleteById(id uint64) error
}

type Region interface {
	GetAll(groupId uint64) ([]entities.Region, error)
	Create(region entities.Region, groupId uint64) (uint64, error)
	GetById(id uint64) (entities.Region, error)
	UpdateById(id uint64, regionUpdate entities.RegionUpdate) error
	DeleteById(id uint64) error
}

type Service struct {
	Marker
	User
	Group
	Region
}

func NewService(database *controllers.Database) *Service {
	return &Service{
		Marker: NewMarkerService(database.Point, database.Polygon, database.Line),
		User:   NewUserService(database.User),
		Group:  NewGroupService(database.Group),
		Region: NewRegionService(database.Region),
	}
}
