package dbRepository

import (
	"errors"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"

	"github.com/jinzhu/gorm"
)

type ModulesRepo struct {
	db *gorm.DB
}

func ModulesRepositoryInit(db *gorm.DB) *ModulesRepo {
	return &ModulesRepo{db}
}

//ModulesRepo implements the repository.ModulesRepository interface
// var _ services.ModulesAppInterface = &ModulesRepo{}

//GetAll all data
func (r *ModulesRepo) GetAll() ([]entity.Modules, error) {
	var datas []entity.Modules
	var err error
	err = r.db.Debug().Order("created_at desc").Find(&datas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

//GetAll all data
func (r *ModulesRepo) GetAllModulesMerge() ([]dto.ModulesAndPermission, error) {
	var err error
	var datas []dto.ModulesAndPermission
	err = r.db.Debug().Table("modules").Order("created_at desc").Find(&datas).Error

	//TODO: nasil preload yapilir bakilacak
	// var datas []entity.Modules
	// err = r.db.Debug().Preload("Permission").Take(&datas).Error

	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

//GetAll all data
func (r *ModulesRepo) GetAllModulesMergePermission() ([]dto.ModulesAndPermissionRole, error) {
	var err error
	var datas []dto.ModulesAndPermissionRole
	err = r.db.Debug().Table("modules").Order("created_at desc").Find(&datas).Error

	//TODO: nasil preload yapilir bakilacak
	// var datas []entity.Modules
	// err = r.db.Debug().Preload("Permission").Take(&datas).Error

	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}
