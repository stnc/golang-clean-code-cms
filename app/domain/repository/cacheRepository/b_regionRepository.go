package cacheRepository

import (
	"encoding/json"
	"fmt"
	"stncCms/app/domain/cache"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stnccollection"
	repository "stncCms/app/domain/repository/dbRepository"
	"time"

	"github.com/jinzhu/gorm"
)

// RegionRepo struct
type RegionRepo struct {
	db *gorm.DB
}

// RegionRepositoryInit initial
func RegionRepositoryInit(db *gorm.DB) *RegionRepo {
	return &RegionRepo{db}
}

func getByIDRegion(db *gorm.DB, id uint64) (*entity.Region, error) {
	repo := repository.RegionRepositoryInit(db)
	datas, _ := repo.GetByID(id)
	return datas, nil
}

// GetByID get data
func (r *RegionRepo) GetByID(id uint64) (*entity.Region, error) {

	var data *entity.Region
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByIDRegion(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()

		key := "RegionGetByID" + stnccollection.Uint64toString(id)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByIDRegion(r.db, id)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
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

// GetAll all data
func (r *RegionRepo) GetAll() ([]entity.Region, error) {
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	var data []entity.Region
	if cacheControl == "false" {
		data, _ = getAllRegion(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "RegionGetAll"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllRegion(r.db)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
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

func getAllRegion(db *gorm.DB) ([]entity.Region, error) {
	repo := repository.RegionRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}
