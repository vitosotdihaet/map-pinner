package services

import (
	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)


type PolygonService struct {
	databasePoint controllers.Point
	databasePolygon controllers.Polygon
}

func NewPolygonService(databasePolygon controllers.Polygon, databasePoint controllers.Point) *PolygonService {
	return &PolygonService {
		databasePolygon: databasePolygon,
		databasePoint: databasePoint,
	}
}

func (service *PolygonService) Create(polygon entities.Polygon) (uint64, error) {
	return service.databasePolygon.Create(polygon)
}

func (service *PolygonService) GetAll() ([]entities.Polygon, error) {
	return service.databasePolygon.GetAll()
}

func (service *PolygonService) GetById(id uint64) (entities.Polygon, error) {
	return service.databasePolygon.GetById(id)
}

func (service *PolygonService) UpdateById(newPolygon entities.Polygon) error {
	return service.databasePolygon.UpdateById(newPolygon)
}

func (service *PolygonService) DeleteById(id uint64) error {
	return service.databasePolygon.DeleteById(id)
}
