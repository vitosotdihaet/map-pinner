package controllers

import (
	"github.com/jmoiron/sqlx"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type Point interface {
	GetAll(regionId uint64) ([]entities.Point, error)
	Create(regionId uint64, point entities.Point) (uint64, error)
	GetById(id uint64) (entities.Point, error)
	UpdateById(id uint64, pointUpdate entities.PointUpdate) error
	DeleteById(id uint64) error
}

type Polygon interface {
	GetAll(regionId uint64) ([]entities.Polygon, error)
	Create(regionId uint64, polygon entities.Polygon) (uint64, error)
	GetById(id uint64) (entities.Polygon, error)
	UpdateById(id uint64, newPolygon entities.PolygonUpdate) error
	DeleteById(id uint64) error
}

type Line interface {
	GetAll(regionId uint64) ([]entities.Line, error)
	Create(regionId uint64, line entities.Line) (uint64, error)
	GetById(id uint64) (entities.Line, error)
	UpdateById(id uint64, lineUpdate entities.LineUpdate) error
	DeleteById(id uint64) error
}

type User interface {
	GetAll() ([]entities.User, error)
	Create(user entities.User, password entities.HashedPassword) (uint64, error)
	GetById(id uint64) (entities.User, error)
	GetByName(user entities.User) (*entities.User, entities.HashedPassword, error)
	ExistsWithName(userName string) (bool, error)
	// UpdateById(id uint64, lineUpdate entities.GroupUpdate) error
	DeleteById(id uint64) error
}

type Group interface {
	GetAll(userId uint64) ([]entities.Group, error)
	Create(group entities.Group, authorId uint64) (uint64, error)
	GetById(id uint64) (*entities.Group, error)
	// UpdateById(id uint64, lineUpdate entities.GroupUpdate) error
	GetAllUsers(id uint64) ([]entities.User, []string, error)
	DeleteById(id uint64) error
	AddUserToGroup(id uint64, userName string, roleId uint64) error
}

type Region interface {
	GetAll(groupId uint64) ([]entities.Region, error)
	Create(region entities.Region, groupId uint64) (uint64, error)
	GetById(id uint64) (entities.Region, error)
	// UpdateById(id uint64, lineUpdate entities.RegionUpdate) error
	DeleteById(id uint64) error
}

type Role interface {
	GetAllRoles() (map[uint64]string, error)

	HasAtLeastRoleInGroup(groupId uint64, userId uint64, role string) (bool, error)
	HasAtLeastRoleInRegion(regionId uint64, userId uint64, role string) (bool, error)
	HasAtLeastRoleForMarker(markerType entities.MarkerType, markerId uint64, userId uint64, role string) (bool, error)
	ThereIsARoleWithId(roleId uint64) (bool, error)

	ThereIsASystemRoleWithId(roleId uint64) (bool, error)
	HasAtLeastSystemRole(userId uint64, role string) (bool, error)

	GetRoleID(userId uint64, groupId uint64) (uint64, error)
}

type Database struct {
	Point
	Polygon
	Line
	User
	Group
	Region
	Role
}

func NewDatabase(postgres *sqlx.DB) *Database {
	return &Database{
		Point:   NewPointPostgres(postgres),
		Polygon: NewPolygonPostgres(postgres),
		Line:    NewLinePostgres(postgres),
		User:    NewUserPostgres(postgres),
		Group:   NewGroupPostgres(postgres),
		Region:  NewRegionPostgres(postgres),
		Role:    NewRolePostgres(postgres),
	}
}
