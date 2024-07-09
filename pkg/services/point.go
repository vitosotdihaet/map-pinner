package services

import (
	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)


type PointService struct {
	database controllers.Point
}

func NewPointService(database controllers.Point) *PointService {
	return &PointService{database: database}
}

func (service *PointService) Create(point entities.Point) (uint64, error) {
	return service.database.Create(point)
}

func (service *PointService) GetAll() ([]entities.Point, error) {
	return service.database.GetAll()
}

func (service *PointService) GetById(id uint64) (entities.Point, error) {
	return service.database.GetById(id)
}

func (service *PointService) UpdateById(id uint64, pointUpdate entities.PointUpdate) error {
	return service.database.UpdateById(id, pointUpdate)
}

func (service *PointService) DeleteById(id uint64) error {
	return service.database.DeleteById(id)
}
