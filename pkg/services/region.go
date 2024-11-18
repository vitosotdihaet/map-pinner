package services

import (
	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type RegionService struct {
	database controllers.Region
}

func NewRegionService(database controllers.Region) *RegionService {
	return &RegionService{
		database: database,
	}
}

func (service *RegionService) Create(region entities.Region) (uint64, error) {
	return service.database.Create(region)
}

func (service *RegionService) GetAll() ([]entities.Region, error) {
	return service.database.GetAll()
}

func (service *RegionService) GetById(id uint64) (entities.Region, error) {
	return service.database.GetById(id)
}

func (service *RegionService) UpdateById(id uint64, regionUpdate entities.RegionUpdate) error {
	return service.database.UpdateById(id, regionUpdate)
}

func (service *RegionService) DeleteById(id uint64) error {
	return service.database.DeleteById(id)
}
