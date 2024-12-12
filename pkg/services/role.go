package services

import "github.com/vitosotdihaet/map-pinner/pkg/controllers"

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
