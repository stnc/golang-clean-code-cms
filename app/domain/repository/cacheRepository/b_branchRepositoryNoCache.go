package cacheRepository

import (
	"stncCms/app/domain/entity"
	repository "stncCms/app/domain/repository/dbRepository"
)

//Save data
func (r *BranchRepo) Save(data *entity.Branches) (*entity.Branches, map[string]string) {
	repo := repository.BranchRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

//Update upate data
func (r *BranchRepo) Update(data *entity.Branches) (*entity.Branches, map[string]string) {
	repo := repository.BranchRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err

}

//Delete delete data
func (r *BranchRepo) Delete(id uint64) error {
	repo := repository.BranchRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}
