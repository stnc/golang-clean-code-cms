package cacheRepository

import (
	"stncCms/app/domain/entity"
	repository "stncCms/app/domain/repository/dbRepository"
)

// Save data
func (r *RegionRepo) Save(data *entity.Region) (*entity.Region, map[string]string) {
	repo := repository.RegionRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

// Update upate data
func (r *RegionRepo) Update(data *entity.Region) (*entity.Region, map[string]string) {
	repo := repository.RegionRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err

}

// Delete delete data
func (r *RegionRepo) Delete(id uint64) error {
	repo := repository.RegionRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}
