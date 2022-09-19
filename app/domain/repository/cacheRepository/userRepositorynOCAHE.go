package cacheRepository

import (
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
	repository "stncCms/app/domain/repository/dbRepository"
)

func (r *UserRepo) SaveUser(data *entity.Users) (*entity.Users, map[string]string) {
	repo := repository.UserRepositoryInit(r.db)
	datas, err := repo.SaveUser(data)
	return datas, err
}

// Save data
func (r *UserRepo) Save(data *entity.Users) (*entity.Users, map[string]string) {
	repo := repository.UserRepositoryInit(r.db)
	datas, err := repo.Save(data)
	return datas, err
}

func (r *UserRepo) SaveDto(data *dto.User) (*dto.User, map[string]string) {
	repo := repository.UserRepositoryInit(r.db)
	datas, err := repo.SaveDto(data)
	return datas, err
}

// Update upate data
func (r *UserRepo) Update(data *entity.Users) (*entity.Users, map[string]string) {
	repo := repository.UserRepositoryInit(r.db)
	datas, err := repo.Update(data)
	return datas, err
}

func (r *UserRepo) UpdateDto(data *dto.User) (*dto.User, map[string]string) {
	repo := repository.UserRepositoryInit(r.db)
	datas, err := repo.UpdateDto(data)
	return datas, err
}

// Count fat
func (r *UserRepo) Count(totalCount *int64) {
	var count int64
	repo := repository.UserRepositoryInit(r.db)
	repo.Count(&count)
	*totalCount = count
}

// Delete data
func (r *UserRepo) Delete(id uint64) error {
	repo := repository.UserRepositoryInit(r.db)
	err := repo.Delete(id)
	return err
}

// SetKioskSliderUpdate update data
func (r *UserRepo) SetUserStatusUpdate(id uint64, status int) {
	repo := repository.UserRepositoryInit(r.db)
	repo.SetUserStatusUpdate(id, status)
}

// api kullanacak
func (r *UserRepo) GetUserByEmailAndPassword(u *entity.Users) (*entity.Users, map[string]string) {
	repo := repository.UserRepositoryInit(r.db)
	data, _ := repo.GetUserByEmailAndPassword(u)
	return data, nil
}
func (r *UserRepo) GetUserByEmailAndPassword2(email string, InputPassword string) (*entity.Users, bool) {
	repo := repository.UserRepositoryInit(r.db)
	data, result := repo.GetUserByEmailAndPassword2(email, InputPassword)
	return data, result
}
