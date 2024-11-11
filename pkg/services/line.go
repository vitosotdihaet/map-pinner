package services

import (
	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type LineService struct {
	database controllers.Line
}

func NewLineService(database controllers.Line) *LineService {
	return &LineService{
		database: database,
	}
}

func (service *LineService) Create(line entities.Line) (uint64, error) {
	return service.database.Create(line)
}

func (service *LineService) GetAll() ([]entities.Line, error) {
	return service.database.GetAll()
}

func (service *LineService) GetById(id uint64) (entities.Line, error) {
	return service.database.GetById(id)
}

func (service *LineService) UpdateById(id uint64, lineUpdate entities.LineUpdate) error {
	return service.database.UpdateById(id, lineUpdate)
}

func (service *LineService) DeleteById(id uint64) error {
	return service.database.DeleteById(id)
}
