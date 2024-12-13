package services

import (
	"errors"

	"github.com/vitosotdihaet/map-pinner/pkg/controllers"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

type RegionService struct {
	regionDB controllers.Region
	roleDB   controllers.Role
}

func NewRegionService(regionDB controllers.Region, roleDB controllers.Role) *RegionService {
	return &RegionService{
		regionDB: regionDB,
		roleDB:   roleDB,
	}
}

func (service *RegionService) GetAll(groupId uint64, userId uint64) ([]entities.Region, error) {
	ok, err := service.roleDB.HasAtLeastRoleInGroup(groupId, userId, "viewer")
	if err != nil {
		return []entities.Region{}, err
	}

	if !ok {
		return []entities.Region{}, errors.New("not enough rights")
	}

	return service.regionDB.GetAll(groupId)
}

func (service *RegionService) Create(region entities.Region, groupId uint64, authorId uint64) (uint64, error) {
	ok, err := service.roleDB.HasAtLeastRoleInGroup(groupId, authorId, "editor")
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("not enough rights")
	}

	return service.regionDB.Create(region, groupId)
}

// func (service *RegionService) GetById(id uint64) (entities.Region, error) {
// 	return service.database.GetById(id)
// }

// func (service *RegionService) UpdateById(id uint64, regionUpdate entities.RegionUpdate) error {
// 	return service.database.UpdateById(id, regionUpdate)
// }

func (service *RegionService) DeleteById(id uint64, userId uint64) error {
	ok, err := service.roleDB.HasAtLeastRoleInRegion(id, userId, "editor")
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("not enough rights")
	}

	return service.regionDB.DeleteById(id)
}
