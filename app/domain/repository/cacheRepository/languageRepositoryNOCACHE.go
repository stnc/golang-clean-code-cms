package cacheRepository

import (
	"stncCms/app/domain/entity"
	repository "stncCms/app/domain/repository/dbRepository"
)

//Save data
func (r *LanguageRepo) Save(data *entity.Languages) (*entity.Languages, map[string]string) {
	repo := repository.LanguageRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

//Update upate data
func (r *LanguageRepo) Update(data *entity.Languages) (*entity.Languages, map[string]string) {
	repo := repository.LanguageRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err
}

//Delete data
func (r *LanguageRepo) Delete(id uint64) error {
	repo := repository.LanguageRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}
