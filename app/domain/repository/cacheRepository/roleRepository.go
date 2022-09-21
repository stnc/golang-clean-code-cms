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

type RoleRepo struct {
	db *gorm.DB
}

func RoleRepositoryInit(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db}
}

//RoleRepo implements the repository.RoleRepository interface
var _ services.RoleAppInterface = &RoleRepo{}

//GetAll all data
func (r *RoleRepo) GetAll() ([]entity.Role, error) {
	var data []entity.Role
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllRole(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllRole"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllRole(r.db)
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
func getAllRole(db *gorm.DB) ([]entity.Role, error) {
	repo := repository.RoleRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

//Save data
func (r *RoleRepo) Save(data *entity.Role) (*entity.Role, map[string]string) {
	repo := repository.RoleRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

func (r *RoleRepo) EditList(modulID int, roleID int) ([]dto.RoleEditList, error) {
	repo := repository.RoleRepositoryInit(r.db)
	datas, err := repo.EditList(modulID, roleID)
	return datas, err
}

//Count
func (r *RoleRepo) Count(totalCount *int64) {
	var count int64
	repo := repository.RoleRepositoryInit(r.db)
	repo.Count(&count)
	*totalCount = count
}

//Delete data
func (r *RoleRepo) Delete(id uint64) error {
	repo := repository.RoleRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}

//GetAllPagination pagination all data
func GetAllPaginationrole(db *gorm.DB, perPage int, offset int) ([]entity.Role, error) {
	repo := repository.RoleRepositoryInit(db)
	data, _ := repo.GetAllPagination(perPage, offset)
	return data, nil
}

//GetAllPagination pagination all data
func (r *RoleRepo) GetAllPagination(perPage int, offset int) ([]entity.Role, error) {
	var data []entity.Role
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = GetAllPaginationrole(r.db, perPage, offset)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetAllPaginationpost_" + stnccollection.IntToString(perPage) + stnccollection.IntToString(offset)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = GetAllPaginationrole(r.db, perPage, offset)
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

//GetByID get data
func (r *RoleRepo) GetByID(id int) (*entity.Role, error) {
	var data *entity.Role
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getByIDRole(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getByIDRole_" + stnccollection.IntToString(id)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getByIDRole(r.db, id)
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

func getByIDRole(db *gorm.DB, id int) (*entity.Role, error) {
	repo := repository.RoleRepositoryInit(db)
	data, _ := repo.GetByID(id)
	return data, nil
}

func (r *RoleRepo) Update(data *entity.Role) (*entity.Role, map[string]string) {
	repo := repository.RoleRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

func (r *RoleRepo) UpdateTitle(id int, title string) {
	repo := repository.RoleRepositoryInit(r.db)
	repo.UpdateTitle(id, title)
}
