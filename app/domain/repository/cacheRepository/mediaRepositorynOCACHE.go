package cacheRepository

import (
	"stncCms/app/domain/entity"
	repository "stncCms/app/domain/repository/dbRepository"
)

//Save data
func (r *MediaRepo) Save(data *entity.Media) (*entity.Media, map[string]string) {

	repo := repository.MediaRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

//Update upate data
func (r *MediaRepo) Update(data *entity.Media) (*entity.Media, map[string]string) {
	repo := repository.MediaRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

//Count fat
func (r *MediaRepo) Count(modulID int, contentID int, MediaTotalCount *int) {
	var count int
	repo := repository.MediaRepositoryInit(r.db)
	repo.Count(modulID, contentID, &count)
	*MediaTotalCount = count
}

//Delete data
func (r *MediaRepo) Delete(id uint64) error {

	repo := repository.MediaRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}
