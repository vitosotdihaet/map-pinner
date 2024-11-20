package services

import (
	"fmt"

	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type MarkerService struct {
	pointDB   controllers.Point
	polygonDB controllers.Polygon
	lineDB    controllers.Line
}

func NewMarkerService(pointDB controllers.Point, polygonDB controllers.Polygon, lineDB controllers.Line) *MarkerService {
	return &MarkerService{
		pointDB:   pointDB,
		polygonDB: polygonDB,
		lineDB:    lineDB,
	}
}

func (service *MarkerService) Create(regionId uint64, marker entities.Marker) (uint64, error) {
	switch marker.GetType() {
	case entities.PointType:
		return service.pointDB.Create(regionId, *marker.(*entities.Point))
	case entities.PolygonType:
		return service.polygonDB.Create(regionId, *marker.(*entities.Polygon))
	case entities.LineType:
		return service.lineDB.Create(regionId, *marker.(*entities.Line))
	}
	return 0, fmt.Errorf("service: invalid marker type %s", marker.GetType())
}

func (service *MarkerService) GetAll(regionId uint64) ([]entities.Marker, error) {
	points, err := service.pointDB.GetAll(regionId)
	if err != nil {
		return []entities.Marker{}, err
	}
	polygons, err := service.polygonDB.GetAll(regionId)
	if err != nil {
		return []entities.Marker{}, err
	}
	lines, err := service.lineDB.GetAll(regionId)
	if err != nil {
		return []entities.Marker{}, err
	}

	var all []entities.Marker

	// Append points
	for _, point := range points {
		all = append(all, entities.PointType)
		all = append(all, &point)
	}

	// Append polygons
	for _, polygon := range polygons {
		all = append(all, entities.PolygonType)
		all = append(all, &polygon)
	}

	// Append lines
	for _, line := range lines {
		all = append(all, entities.LineType)
		all = append(all, &line)
	}

	return all, nil
}

func (service *MarkerService) GetById(markerType entities.MarkerType, id uint64) (entities.Marker, error) {
	switch markerType {
	case entities.PointType:
		out, err := service.pointDB.GetById(id)
		return &out, err
	case entities.PolygonType:
		out, err := service.polygonDB.GetById(id)
		return &out, err
	case entities.LineType:
		out, err := service.lineDB.GetById(id)
		return &out, err
	}
	return &entities.None{}, fmt.Errorf("service: invalid marker type %s", markerType)
}

func (service *MarkerService) UpdateById(id uint64, markerUpdate entities.Marker) error {
	switch markerUpdate.GetType() {
	case entities.PointType:
		return service.pointDB.UpdateById(id, *markerUpdate.(*entities.PointUpdate))
	case entities.PolygonType:
		return service.polygonDB.UpdateById(id, *markerUpdate.(*entities.PolygonUpdate))
	case entities.LineType:
		return service.lineDB.UpdateById(id, *markerUpdate.(*entities.LineUpdate))
	}
	return fmt.Errorf("service: invalid marker type %s", markerUpdate.GetType())
}

func (service *MarkerService) DeleteById(markerType entities.MarkerType, id uint64) error {
	switch markerType {
	case entities.PointType:
		return service.pointDB.DeleteById(id)
	case entities.PolygonType:
		return service.polygonDB.DeleteById(id)
	case entities.LineType:
		return service.lineDB.DeleteById(id)
	}
	return fmt.Errorf("service: invalid marker type %s", markerType)
}
