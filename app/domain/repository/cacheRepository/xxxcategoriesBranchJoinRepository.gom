package repository

import (
	"errors"
	"stncCms/app/domain/entity"
	"strings"

	"github.com/jinzhu/gorm"
)

//CategoriesBranchJoinRepo struct
type CategoriesBranchJoinRepo struct {
	db *gorm.DB
}

////SELECT * FROM kiosk_slider p1 WHERE id IN (SELECT kiosk_id FROM  categories_kiosk_join  where categories_kiosk_join.category_id=1)

//CategoriesBranchJoinRepositoryInit initial
func CategoriesBranchJoinRepositoryInit(db *gorm.DB) *CategoriesBranchJoinRepo {
	return &CategoriesBranchJoinRepo{db}
}

//PostRepo implements the repository.PostRepository interface
// var _ interfaces.CategoriesBranchJoinAppInterface = &CategoriesBranchJoinRepo{}

//Save data
func (r *CategoriesBranchJoinRepo) Save(cat *entity.CategoriesBranch) (*entity.CategoriesBranch, map[string]string) {
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

//GetByID get data
func (r *CategoriesBranchJoinRepo) GetByID(id uint64) (*entity.CategoriesBranch, error) {
	var cat entity.CategoriesBranch
	err := r.db.Debug().Where("id = ?", id).Take(&cat).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")

	// }
	return &cat, nil
}

//GetAllforCatID get data
func (r *CategoriesBranchJoinRepo) GetAllforBranchID(branchID uint64) ([]entity.CategoriesBranch, error) {
	var cat []entity.CategoriesBranch
	err := r.db.Debug().Limit(100).Where("branch_id = ?", branchID).Order("created_at desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}

	// if err.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")
	// }
	return cat, nil
}

//GetAllforUserID all data
func (r *CategoriesBranchJoinRepo) GetAllforUserID(UserID uint64) ([]entity.CategoriesBranch, error) {
	var cat []entity.CategoriesBranch
	err := r.db.Debug().Limit(100).Where("user_id = ?", UserID).Order("created_at desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")
	// }
	return cat, nil
}

//GetAll all data
func (r *CategoriesBranchJoinRepo) GetAll() ([]entity.CategoriesBranch, error) {
	var cat []entity.CategoriesBranch
	err := r.db.Debug().Limit(100).Order("created_at desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")
	// }
	return cat, nil
}

//Update upate data
func (r *CategoriesBranchJoinRepo) Update(cat *entity.CategoriesBranch) (*entity.CategoriesBranch, map[string]string) {
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

//Delete delete data
func (r *CategoriesBranchJoinRepo) Delete(id uint64) error {
	var cat entity.CategoriesBranch
	err := r.db.Debug().Where("id = ?", id).Unscoped().Delete(&cat).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

//DeleteForKioskID delete data
func (r *CategoriesBranchJoinRepo) DeleteForUserID(userId uint64) error {
	var cat entity.CategoriesBranch
	err := r.db.Debug().Where("user_id = ?", userId).Unscoped().Delete(&cat).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

//DeleteForCatID delete data
func (r *CategoriesBranchJoinRepo) DeleteForBranchID(branchID uint64) error {
	var cat entity.CategoriesBranch
	err := r.db.Debug().Where("branch_id = ?", branchID).Unscoped().Delete(&cat).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}
