package services

import (
	"errors"

	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
)

type RoleService struct {
	database controllers.Role
}

func NewRoleService(database controllers.Role) *RoleService {
	return &RoleService{
		database: database,
	}
}

func (service *RoleService) GetAll() (map[uint64]string, error) {
	return service.database.GetAllRoles()
}

func (service *RoleService) HasAtLeastSystemRole(userId uint64, role string) (bool, error) {
	return service.database.HasAtLeastSystemRole(userId, role)
}

func (service *RoleService) GetRoleID(userId uint64, groupId uint64) (uint64, error) {
	ok, err := service.database.HasAtLeastRoleInGroup(groupId, userId, "viewer")
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("not enough rights")
	}

	return service.database.GetRoleID(userId, groupId)
}
