package dbRepository

import (
	"errors"
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

// GetAll all data
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
