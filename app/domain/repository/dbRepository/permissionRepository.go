package dbRepository

import (
	"errors"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
	"stncCms/app/services"

	"github.com/jinzhu/gorm"
)

type PermissionRepo struct {
	db *gorm.DB
}

func PermissionRepositoryInit(db *gorm.DB) *PermissionRepo {
	return &PermissionRepo{db}
}

//PermissionRepo implements the repository.PermissionRepository interface
var _ services.PermissionAppInterface = &PermissionRepo{}

//GetAll all data
func (r *PermissionRepo) GetAll() ([]entity.Permission, error) {
	var datas []entity.Permission
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

func (r *PermissionRepo) GetAllPaginationermissionForModulID(modulId int) ([]entity.Permission, error) {
	var datas []entity.Permission
	var err error
	err = r.db.Debug().Where("modul_id = ?", modulId).Order("id desc").Find(&datas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (r *PermissionRepo) GetUserPermission(roleID int) ([]dto.RbcaCheck, error) {
	var data []dto.RbcaCheck
	var err error
	query := r.db.Table(entity.RolePermissonTableName + " AS role_permission")
	query = query.Select(`permission.modul_id AS modul_id,
	         role_permission.role_id,role_permission.permission_id,
		     role_permission.active AS role_permission_active ,
			 permission.title AS permission_Title,
			 permission.controller AS controller,
			 permission.func_name AS func ,
		   CONCAT (permission.controller,'-',permission.func_name) As permission_name`)
	query = query.Joins(" INNER JOIN rbca_permission AS  permission ON permission.id=role_permission.permission_id   ")
	query = query.Where("role_permission.role_id=? ", roleID)
	err = query.Find(&data).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return data, nil
}
