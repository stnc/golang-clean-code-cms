package cacheRepository

import (
	"stncCms/app/domain/entity"
	repository "stncCms/app/domain/repository/dbRepository"
)

//Save data
func (r *PostRepo) Save(data *entity.Post) (*entity.Post, map[string]string) {
	repo := repository.PostRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

//Update upate data
func (r *PostRepo) Update(data *entity.Post) (*entity.Post, map[string]string) {
	repo := repository.PostRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err
}

//Count fat
func (r *PostRepo) Count(totalCount *int64) {
	var count int64
	repo := repository.PostRepositoryInit(r.db)
	repo.Count(&count)
	*totalCount = count
}

//Delete data
func (r *PostRepo) Delete(id uint64) error {
	repo := repository.PostRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}
