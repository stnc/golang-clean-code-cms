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

//CatRepo struct
type CatRepo struct {
	db *gorm.DB
}

//CatRepositoryInit initial
func CatRepositoryInit(db *gorm.DB) *CatRepo {
	return &CatRepo{db}
}

//PostRepo implements the repository.PostRepository interface
// var _ interfaces.CatAppInterface = &CatRepo{}
func getByIDCategories(db *gorm.DB, id uint64) (*entity.Categories, error) {
	repo := repository.CatRepositoryInit(db)
	datas, _ := repo.GetByID(id)
	return datas, nil
}

//GetByID get data
func (r *CatRepo) GetByID(id uint64) (*entity.Categories, error) {
	var data *entity.Categories
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByIDCategories(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()

		key := "getByIDCategories" + stnccollection.Uint64toString(id)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByIDCategories(r.db, id)
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

//GetAll all data
func (r *CatRepo) GetAll() ([]entity.Categories, error) {
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	var data []entity.Categories
	if cacheControl == "false" {
		data, _ = getAllCategories(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllCategories"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllCategories(r.db)
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

func getAllCategories(db *gorm.DB) ([]entity.Categories, error) {
	repo := repository.CatRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}
