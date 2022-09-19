package cacheRepository

import (
	"stncCms/app/domain/entity"
	repository "stncCms/app/domain/repository/dbRepository"
)

//PostRepo implements the repository.PostRepository interface
// var _ interfaces.CatAppInterface = &CatRepo{}

//Save data
func (r *CatRepo) Save(data *entity.Categories) (*entity.Categories, map[string]string) {
	repo := repository.CatRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

//Update upate data
func (r *CatRepo) Update(data *entity.Categories) (*entity.Categories, map[string]string) {
	repo := repository.CatRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err
}

//Delete delete data
func (r *CatRepo) Delete(id uint64) error {
	repo := repository.CatRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}
