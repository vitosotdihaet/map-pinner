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
	return service.database.GetAll()
}
