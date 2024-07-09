package services

import (
	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)


type PolygonService struct {
	database controllers.Polygon
}

func NewPolygonService(database controllers.Polygon) *PolygonService {
	return &PolygonService {
		database: database,
	}
}

func (service *PolygonService) Create(polygon entities.Polygon) (uint64, error) {
	return service.database.Create(polygon)
}

func (service *PolygonService) GetAll() ([]entities.Polygon, error) {
	return service.database.GetAll()
}

func (service *PolygonService) GetById(id uint64) (entities.Polygon, error) {
	return service.database.GetById(id)
}

func (service *PolygonService) UpdateById(id uint64, polygonUpdate entities.PolygonUpdate) error {
	return service.database.UpdateById(id, polygonUpdate)
}

func (service *PolygonService) DeleteById(id uint64) error {
	return service.database.DeleteById(id)
}
