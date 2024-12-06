package services

import (
	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type UserService struct {
	database controllers.User
}

func NewUserService(database controllers.User) *UserService {
	return &UserService{
		database: database,
	}
}

func (service *UserService) Create(user entities.User, password entities.HashedPassword) (uint64, error) {
	return service.database.Create(user, password)
}

func (service *UserService) GetAll() ([]entities.User, error) {
	return service.database.GetAll()
}

func (service *UserService) GetById(id uint64) (entities.User, error) {
	return service.database.GetById(id)
}

func (service *UserService) GetByNamePassword(user entities.User, password entities.HashedPassword) (*entities.User, error) {
	return service.database.GetByNamePassword(user, password)
}

// func (service *UserService) UpdateById(id uint64, userUpdate entities.UserUpdate) error {
// 	return service.database.UpdateById(id, userUpdate)
// }

func (service *UserService) DeleteById(id uint64) error {
	return service.database.DeleteById(id)
}
