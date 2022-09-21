package dbRepository

import (
	"errors"
	"stncCms/app/domain/entity"
	"stncCms/app/services"
	"strings"

	"github.com/jinzhu/gorm"
)

type RolePermissionRepo struct {
	db *gorm.DB
}

func RolePermissionRepositoryInit(db *gorm.DB) *RolePermissionRepo {
	return &RolePermissionRepo{db}
}

//PermissionRepo implements the repository.PermissionRepository interface
var _ services.RolePermissionAppInterface = &RolePermissionRepo{}

//GetAll all data
func (r *RolePermissionRepo) GetAll() ([]entity.RolePermisson, error) {
	var datas []entity.RolePermisson
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
func (r *RolePermissionRepo) Save(data *entity.RolePermisson) (*entity.RolePermisson, map[string]string) {
	dbErr := map[string]string{}
	var err error
	err = r.db.Debug().Create(&data).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "data title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return data, nil
}

//Update upate data
func (r *RolePermissionRepo) Update(post *entity.RolePermisson) (*entity.RolePermisson, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Save(&post).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return post, nil
}

func (r *RolePermissionRepo) UpdateActiveStatus(roleId int, permissionId int, active int) {
	r.db.Debug().Model(&entity.RolePermisson{}).Where("role_id = ? and permission_id=?", roleId, permissionId).Update("active", active)
}
