package cacheRepository

import (
	"encoding/json"
	"fmt"
	"stncCms/app/domain/cache"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stnccollection"
	repository "stncCms/app/domain/repository/dbRepository"
	"stncCms/app/services"
	"time"

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
	var data []entity.Permission
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = GetAllPaginationermission(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetAllPaginationermission"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = GetAllPaginationermission(r.db)
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
func GetAllPaginationermission(db *gorm.DB) ([]entity.Permission, error) {
	repo := repository.PermissionRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

//getAllPaginationermissionForModulID all data
func (r *PermissionRepo) GetAllPaginationermissionForModulID(modulId int) ([]entity.Permission, error) {
	var data []entity.Permission
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllPaginationermissionForModulID(modulId, r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllPaginationermissionForModulID" + stnccollection.IntToString(modulId)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllPaginationermissionForModulID(modulId, r.db)
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
func getAllPaginationermissionForModulID(modulId int, db *gorm.DB) ([]entity.Permission, error) {
	repo := repository.PermissionRepositoryInit(db)
	data, _ := repo.GetAllPaginationermissionForModulID(modulId)
	return data, nil
}

//getAllPaginationermissionForModulID all data
func (r *PermissionRepo) GetUserPermission(roleID int) ([]dto.RbcaCheck, error) {
	var data []dto.RbcaCheck
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getUserPermission(roleID, r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetUserPermission" + stnccollection.IntToString(roleID)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getUserPermission(roleID, r.db)
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
func getUserPermission(roleID int, db *gorm.DB) ([]dto.RbcaCheck, error) {
	repo := repository.PermissionRepositoryInit(db)
	data, _ := repo.GetUserPermission(roleID)
	return data, nil
}
