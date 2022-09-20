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

type ModulesRepo struct {
	db *gorm.DB
}

func ModulesRepositoryInit(db *gorm.DB) *ModulesRepo {
	return &ModulesRepo{db}
}

var _ services.ModulesAppInterface = &ModulesRepo{}

// GetAll all data
func (r *ModulesRepo) GetAll() ([]entity.Modules, error) {
	var data []entity.Modules
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAllModules(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAllModules"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAllModules(r.db)
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
func getAllModules(db *gorm.DB) ([]entity.Modules, error) {
	repo := repository.ModulesRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}
