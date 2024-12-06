package services

import (
	"errors"

	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type GroupService struct {
	groupDB controllers.Group
	roleDB  controllers.Role
}

func NewGroupService(groupController controllers.Group, roleController controllers.Role) *GroupService {
	return &GroupService{groupDB: groupController, roleDB: roleController}
}

func (service *GroupService) Create(group entities.Group, authorId uint64) (uint64, error) {
	return service.groupDB.Create(group, authorId)
}

func (service *GroupService) GetAll(userId uint64) ([]entities.Group, error) {
	return service.groupDB.GetAll(userId)
}

func (service *GroupService) GetById(groupId uint64, userId uint64) (*entities.Group, error) {
	ok, err := service.roleDB.HasAtLeastRoleInGroup(groupId, userId, "viewer")
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("not enough rights")
	}

	return service.groupDB.GetById(groupId)
}

// func (service *GroupService) UpdateById(id uint64, groupUpdate entities.GroupUpdate) error {
// 	return service.database.UpdateById(id, groupUpdate)
// }

func (service *GroupService) DeleteById(id uint64) error {
	return service.groupDB.DeleteById(id)
}

func (service *GroupService) AddUserToGroup(groupId uint64, authorId uint64, userName string, roleId uint64) error {
	ok, err := service.roleDB.ThereIsARoleWithId(roleId)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("invalid role")
	}

	ok, err = service.roleDB.HasAtLeastRoleInGroup(groupId, authorId, "admin")
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("not enough rights")
	}

	return service.groupDB.AddUserToGroup(groupId, userName, roleId)
}
