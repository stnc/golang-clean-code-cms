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

//CatPostRepo struct
type CatPostRepo struct {
	db *gorm.DB
}

//CatPostRepositoryInit initial
func CatPostRepositoryInit(db *gorm.DB) *CatPostRepo {
	return &CatPostRepo{db}
}

//PostRepo implements the repository.PostRepository interface
// var _ interfaces.CatPostAppInterface = &CatPostRepo{}

//GetByID get data
func (r *CatPostRepo) GetByID(id uint64) (*entity.CategoryPosts, error) {
	var data *entity.CategoryPosts
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = getByIDCatPost(r.db, id)
	} else {
		redisClient := cache.RedisDBInit()

		key := "getByIDCatPost" + stnccollection.Uint64toString(id)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = getByIDCatPost(r.db, id)
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

func getByIDCatPost(db *gorm.DB, id uint64) (*entity.CategoryPosts, error) {
	repo := repository.CatPostRepositoryInit(db)
	datas, _ := repo.GetByID(id)
	return datas, nil
}

//GetAllforCatID get data
func (r *CatPostRepo) GetAllforCatID(catid uint64) ([]entity.CategoryPosts, error) {
	var cat []entity.CategoryPosts
	err := r.db.Debug().Limit(100).Where("category_id = ?", catid).Order("created_at desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}

	// if err.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")
	// }
	return cat, nil
}

//GetAllforPostID all data
func (r *CatPostRepo) GetAllforPostID(postid uint64) ([]entity.CategoryPosts, error) {
	var cat []entity.CategoryPosts
	err := r.db.Debug().Limit(100).Where("post_id = ?", postid).Order("created_at desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")
	// }
	return cat, nil
}

//GetAll all data
func (r *CatPostRepo) GetAll() ([]entity.CategoryPosts, error) {
	var cat []entity.CategoryPosts
	err := r.db.Debug().Limit(100).Order("created_at desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")
	// }
	return cat, nil
}
