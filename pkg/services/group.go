package services

import (
	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type GroupService struct {
	database controllers.Group
}

func NewGroupService(database controllers.Group) *GroupService {
	return &GroupService{database: database}
}

func (service *GroupService) Create(group entities.Group) (uint64, error) {
	return service.database.Create(group)
}

func (service *GroupService) GetAll() ([]entities.Group, error) {
	return service.database.GetAll()
}

func (service *GroupService) GetById(id uint64) (entities.Group, error) {
	return service.database.GetById(id)
}

func (service *GroupService) UpdateById(id uint64, groupUpdate entities.GroupUpdate) error {
	return service.database.UpdateById(id, groupUpdate)
}

func (service *GroupService) DeleteById(id uint64) error {
	return service.database.DeleteById(id)
}