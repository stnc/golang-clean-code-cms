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

//MediaRepo struct
type MediaRepo struct {
	db *gorm.DB
}

//MediaRepositoryInit initial
func MediaRepositoryInit(db *gorm.DB) *MediaRepo {
	return &MediaRepo{db}
}

//MediaRepo implements the repository.MediaRepository interface
// var _ interfaces.MediaAppInterface = &MediaRepo{}

//GetByID get data
func (r *MediaRepo) GetByID(id uint64) (*entity.Media, error) {
	var data *entity.Media
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getByIDMedia(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getByIDMedia_" + stnccollection.Uint64toString(id)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getByIDMedia(r.db, id)
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

func getByIDMedia(db *gorm.DB, id uint64) (*entity.Media, error) {
	repo := repository.MediaRepositoryInit(db)
	data, _ := repo.GetByID(id)
	return data, nil
}

//GetAll all data
func (r *MediaRepo) GetAll(modulID int, contentID int) ([]entity.Media, error) {
	var data []entity.Media
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllMedia(r.db, modulID, contentID)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllMedia" + stnccollection.IntToString(modulID) + stnccollection.IntToString(contentID)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllMedia(r.db, modulID, contentID)
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
func getAllMedia(db *gorm.DB, modulID int, contentID int) ([]entity.Media, error) {
	repo := repository.MediaRepositoryInit(db)
	data, _ := repo.GetAll(modulID, contentID)
	return data, nil
}

//GetAllforModul all data
func (r *MediaRepo) GetAllforModul(modulID int, contentID int) ([]entity.Media, error) {
	var data []entity.Media
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllforModulMedia(r.db, modulID, contentID)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllforModulMedia" + stnccollection.IntToString(modulID) + stnccollection.IntToString(contentID)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllforModulMedia(r.db, modulID, contentID)
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

//GetAllforModul all data
func getAllforModulMedia(db *gorm.DB, modulID int, contentID int) ([]entity.Media, error) {
	repo := repository.MediaRepositoryInit(db)
	data, _ := repo.GetAllforModul(modulID, contentID)
	return data, nil
}

//GetAllPagination pagination all data
func GetAllPaginationMedia(db *gorm.DB, perPage int, offset int) ([]entity.Media, error) {
	repo := repository.MediaRepositoryInit(db)
	data, _ := repo.GetAllPagination(perPage, offset)
	return data, nil
}

//GetAllPagination pagination all data
func (r *MediaRepo) GetAllPagination(perPage int, offset int) ([]entity.Media, error) {
	var data []entity.Media
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = GetAllPaginationMedia(r.db, perPage, offset)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetAllPaginationMedia" + stnccollection.IntToString(perPage) + stnccollection.IntToString(offset)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = GetAllPaginationMedia(r.db, perPage, offset)
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
