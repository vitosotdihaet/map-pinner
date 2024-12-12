package services

import (
	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type Marker interface {
	GetAll(regionId uint64, userId uint64) ([]entities.Marker, error)
	Create(regionId uint64, userId uint64, marker entities.Marker) (uint64, error)
	GetById(markerType entities.MarkerType, id uint64, userId uint64) (entities.Marker, error)
	UpdateById(id uint64, markerUpdate entities.Marker, userId uint64) error
	DeleteById(markerType entities.MarkerType, id uint64, userId uint64) error
}

type User interface {
	GetAll() ([]entities.User, error)
	Create(user entities.User, password entities.HashedPassword) (uint64, error)
	GetById(id uint64) (entities.User, error)
	GetByName(user entities.User) (*entities.User, entities.HashedPassword, error)
	// UpdateById(id uint64, lineUpdate entities.GroupUpdate) error
	DeleteById(id uint64) error
}

type Group interface {
	GetAll(userId uint64) ([]entities.Group, error)
	Create(group entities.Group, authorId uint64) (uint64, error)
	GetById(groupId uint64, userId uint64) (*entities.Group, error)
	// UpdateById(id uint64, groupUpdate entities.GroupUpdate) error
	DeleteById(id uint64) error
	AddUserToGroup(groupId uint64, authorId uint64, userName string, roleId uint64) error
}

type Region interface {
	GetAll(groupId uint64) ([]entities.Region, error)
	// TODO: check for roles
	Create(region entities.Region, groupId uint64) (uint64, error)
	GetById(id uint64) (entities.Region, error)
	UpdateById(id uint64, regionUpdate entities.RegionUpdate) error
	DeleteById(id uint64) error
}

type Role interface {
	GetAll() (map[uint64]string, error)
	HasAtLeastSystemRole(userId uint64, role string) (bool, error)
}

type Service struct {
	Marker
	User
	Group
	Region
	Role
}

func NewService(database *controllers.Database) *Service {
	return &Service{
		Marker: NewMarkerService(database.Point, database.Polygon, database.Line, database.Role),
		User:   NewUserService(database.User),
		Group:  NewGroupService(database.Group, database.Role),
		Region: NewRegionService(database.Region),
		Role:   NewRoleService(database.Role),
	}
}
