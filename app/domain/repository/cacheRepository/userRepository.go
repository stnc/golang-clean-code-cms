package cacheRepository

import (
	"encoding/json"
	"fmt"
	"stncCms/app/domain/cache"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stnccollection"
	repository "stncCms/app/domain/repository/dbRepository"
	"stncCms/app/services"
	"time"

	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func UserRepositoryInit(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

// UserRepo implements the repository.UserRepository interface
var _ services.UserAppInterface = &UserRepo{}

func getUser(db *gorm.DB, id uint64) (*entity.Users, error) {
	repo := repository.UserRepositoryInit(db)
	data, _ := repo.GetUser(id)
	return data, nil
}

func (r *UserRepo) GetUser(id uint64) (*entity.Users, error) {
	var data *entity.Users
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getUser(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getUser_" + stnccollection.Uint64toString(id)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getUser(r.db, id)
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

func getUsers(db *gorm.DB) ([]entity.Users, error) {
	repo := repository.UserRepositoryInit(db)
	data, _ := repo.GetUsers()
	return data, nil
}

func (r *UserRepo) GetUsers() ([]entity.Users, error) {
	var data []entity.Users
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getUsers(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getUsers"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getUsers(r.db)
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

///*******************************  sondaradan eklendı

// GetByID get data
func (r *UserRepo) GetByID(id uint64) (*entity.Users, error) {
	var data *entity.Users
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getByIDuser(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getByIDuser_" + stnccollection.Uint64toString(id)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getByIDuser(r.db, id)
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

func getByIDuser(db *gorm.DB, id uint64) (*entity.Users, error) {
	repo := repository.UserRepositoryInit(db)
	data, _ := repo.GetByID(id)
	return data, nil
}

// GetAll all data
func (r *UserRepo) GetAll() ([]entity.Users, error) {
	var data []entity.Users
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = getAlluser(r.db)
	} else {
		redisClient := cache.RedisDBInit()
		key := "getAlluser"
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = getAlluser(r.db)
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
func getAlluser(db *gorm.DB) ([]entity.Users, error) {
	repo := repository.UserRepositoryInit(db)
	data, _ := repo.GetAll()
	return data, nil
}

// GetAllPagination pagination all data
func (r *UserRepo) GetAllPagination(perPage int, offset int) ([]entity.Users, error) {
	var data []entity.Users
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")
	if cacheControl == "false" {
		data, _ = GetAllPaginationuser(r.db, perPage, offset)
	} else {
		redisClient := cache.RedisDBInit()
		key := "GetAllPaginationuser_" + stnccollection.IntToString(perPage) + stnccollection.IntToString(offset)
		cachedProducts, err := redisClient.GetKey(key)
		if err != nil {
			data, _ = GetAllPaginationuser(r.db, perPage, offset)
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
func GetAllPaginationuser(db *gorm.DB, perPage int, offset int) ([]entity.Users, error) {
	repo := repository.UserRepositoryInit(db)
	data, _ := repo.GetAllPagination(perPage, offset)
	return data, nil
}
