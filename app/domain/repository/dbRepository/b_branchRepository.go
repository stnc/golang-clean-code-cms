package dbRepository

import (
	"errors"
	"stncCms/app/domain/entity"
	"strings"

	"github.com/jinzhu/gorm"
)

//BranchRepo struct
type BranchRepo struct {
	db *gorm.DB
}

//BranchRepositoryInit initial
func BranchRepositoryInit(db *gorm.DB) *BranchRepo {
	return &BranchRepo{db}
}

//PostRepo implements the repository.PostRepository interface
// var _ interfaces.CatAppInterface = &BranchRepo{}

//Save data
func (r *BranchRepo) Save(cat *entity.Branches) (*entity.Branches, map[string]string) {
	dbErr := map[string]string{}

	err := r.db.Debug().Create(&cat).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "post title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return cat, nil
}

//Update upate data
func (r *BranchRepo) Update(cat *entity.Branches) (*entity.Branches, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Save(&cat).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return cat, nil
}

//GetByID get data
func (r *BranchRepo) GetByID(id uint64) (*entity.Branches, error) {
	var cat entity.Branches
	err := r.db.Debug().Where("id = ?", id).Take(&cat).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")

	}
	return &cat, nil
}

//GetAll all data
func (r *BranchRepo) GetAll() ([]entity.Branches, error) {
	var cat []entity.Branches
	err := r.db.Debug().Order("id desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return cat, nil
}

//Delete delete data
func (r *BranchRepo) Delete(id uint64) error {
	var cat entity.Branches
	err := r.db.Debug().Where("id = ?", id).Delete(&cat).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}
