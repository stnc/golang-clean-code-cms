package cacheRepository

import (
	"encoding/json"
	"fmt"
	"stncCms/app/domain/cache"
	"stncCms/app/domain/entity"
	repository "stncCms/app/domain/repository/dbRepository"
	"stncCms/app/services"
	"time"

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
	var data []entity.RolePermisson
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllRolePermission(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetAllPaginationermission"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllRolePermission(r.db)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			if err != nil {
				fmt.Println("hata baş")
			}
			return data, nil
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data, nil
}
func getAllRolePermission(db *gorm.DB) ([]entity.RolePermisson, error) {
	repo := repository.RolePermissionRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

//Save data
func (r *RolePermissionRepo) Save(data *entity.RolePermisson) (*entity.RolePermisson, map[string]string) {
	repo := repository.RolePermissionRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

//Update upate data
func (r *RolePermissionRepo) Update(data *entity.RolePermisson) (*entity.RolePermisson, map[string]string) {
	repo := repository.RolePermissionRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err
}

func (r *RolePermissionRepo) UpdateActiveStatus(roleId int, permissionId int, active int) {
	repo := repository.RolePermissionRepositoryInit(r.db)
	repo.UpdateActiveStatus(roleId, permissionId, active)
}
