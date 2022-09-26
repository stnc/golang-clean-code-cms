package services

import (
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
)

// UserAppInterface interface
type UserAppInterface interface {
	SaveUser(*entity.Users) (*entity.Users, map[string]string)
	GetUser(uint64) (*entity.Users, error)
	GetUsers() ([]entity.Users, error)
	GetUserByEmailAndPassword(*entity.Users) (*entity.Users, map[string]string)
	GetUserByEmailAndPassword2(email string, password string) (*entity.Users, bool)

	Save(*entity.Users) (*entity.Users, map[string]string)
	SaveDto(*dto.User) (*dto.User, map[string]string)
	GetByID(uint64) (*entity.Users, error)

	GetAll() ([]entity.Users, error)
	GetAllPagination(int, int) ([]entity.Users, error)
	Update(*entity.Users) (*entity.Users, map[string]string)
	SetUserPassword(id uint64, password string)
	UpdateDto(*dto.User) (*dto.User, map[string]string)
	Count(*int64)
	Delete(uint64) error

	// SetUserUpdate(uint64, int)

}

type userApp struct {
	request UserAppInterface
}

// UserApp implements the UserAppInterface
var _ UserAppInterface = &userApp{}

func (u *userApp) SaveUser(user *entity.Users) (*entity.Users, map[string]string) {
	return u.request.SaveUser(user)
}

func (u *userApp) GetUser(userID uint64) (*entity.Users, error) {
	return u.request.GetUser(userID)
}

func (u *userApp) GetUsers() ([]entity.Users, error) {
	return u.request.GetUsers()
}

func (u *userApp) GetUserByEmailAndPassword(user *entity.Users) (*entity.Users, map[string]string) {
	return u.request.GetUserByEmailAndPassword(user)
}
func (u *userApp) GetUserByEmailAndPassword2(email string, password string) (*entity.Users, bool) {
	return u.request.GetUserByEmailAndPassword2(email, password)
}

///

func (f *userApp) Count(UserTotalCount *int64) {
	f.request.Count(UserTotalCount)
}

func (f *userApp) Save(User *entity.Users) (*entity.Users, map[string]string) {
	return f.request.Save(User)
}

func (f *userApp) SaveDto(User *dto.User) (*dto.User, map[string]string) {
	return f.request.SaveDto(User)
}

func (f *userApp) GetAll() ([]entity.Users, error) {
	return f.request.GetAll()
}

func (f *userApp) GetAllPagination(UsersPerPage int, offset int) ([]entity.Users, error) {
	return f.request.GetAllPagination(UsersPerPage, offset)
}

func (f *userApp) GetByID(UserID uint64) (*entity.Users, error) {
	return f.request.GetByID(UserID)
}

func (f *userApp) Update(User *entity.Users) (*entity.Users, map[string]string) {
	return f.request.Update(User)
}

func (f *userApp) UpdateDto(User *dto.User) (*dto.User, map[string]string) {
	return f.request.UpdateDto(User)
}

func (f *userApp) Delete(UserID uint64) error {
	return f.request.Delete(UserID)
}
func (f *userApp) SetUserPassword(id uint64, password string) {
	f.request.SetUserPassword(id, password)
}

// func (f *userApp) SetUserUpdate(id uint64, status int) {
// 	f.request.SetUserUpdate(id, status)
// }
