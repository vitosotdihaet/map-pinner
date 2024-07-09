package services

import (
	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)


type GraphService struct {
	database controllers.Graph
}

func NewGraphService(database controllers.Graph) *GraphService {
	return &GraphService {
		database: database,
	}
}

func (service *GraphService) Create(graph entities.Graph) (uint64, error) {
	return service.database.Create(graph)
}

func (service *GraphService) GetAll() ([]entities.Graph, error) {
	return service.database.GetAll()
}

func (service *GraphService) GetById(id uint64) (entities.Graph, error) {
	return service.database.GetById(id)
}

func (service *GraphService) UpdateById(newGraph entities.Graph) error {
	return service.database.UpdateById(newGraph)
}

func (service *GraphService) DeleteById(id uint64) error {
	return service.database.DeleteById(id)
}
