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

//LanguageRepo struct
type LanguageRepo struct {
	db *gorm.DB
}

//LanguageRepositoryInit initial
func LanguageRepositoryInit(db *gorm.DB) *LanguageRepo {
	return &LanguageRepo{db}
}

//languageRepo implements the repository.languageRepository interface
// var _ interfaces.languageAppInterface = &languageRepo{}

//GetByID get data
func (r *LanguageRepo) GetByID(id uint64) (*entity.Languages, error) {
	var data *entity.Languages
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getByIDLanguages(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getByIDLanguages_" + stnccollection.Uint64toString(id)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getByIDLanguages(r.db, id)
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
func getByIDLanguages(db *gorm.DB, id uint64) (*entity.Languages, error) {
	repo := repository.LanguageRepositoryInit(db)
	data, _ := repo.GetByID(id)
	return data, nil
}

//GetAll all data
func getAllLanguages(db *gorm.DB) ([]entity.Languages, error) {
	repo := repository.LanguageRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}
func (r *LanguageRepo) GetAll() ([]entity.Languages, error) {
	var data []entity.Languages
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllLanguages(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllLanguages"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllLanguages(r.db)
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
