package services

import (
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
)

//UserAppInterface interface
type RoleAppInterface interface {
	GetAll() ([]entity.Role, error)
	GetByID(int) (*entity.Role, error)
	GetAllPagination(int, int) ([]entity.Role, error)
	Save(*entity.Role) (*entity.Role, map[string]string)
	Update(*entity.Role) (*entity.Role, map[string]string)
	EditList(modulID int, roleID int) ([]dto.RoleEditList, error)
	UpdateTitle(id int, title string)
	Count(*int64)
	Delete(uint64) error
}

type roleApp struct {
	request RoleAppInterface
}

//UserApp implements the UserAppInterface
var _ RoleAppInterface = &roleApp{}

func (f *roleApp) GetAll() ([]entity.Role, error) {
	return f.request.GetAll()
}

func (f *roleApp) Save(data *entity.Role) (*entity.Role, map[string]string) {
	return f.request.Save(data)
}

func (f *roleApp) EditList(modulId int, roleID int) ([]dto.RoleEditList, error) {
	return f.request.EditList(modulId, roleID)
}
func (f *roleApp) Count(totalCount *int64) {
	f.request.Count(totalCount)
}

func (f *roleApp) Delete(ID uint64) error {
	return f.request.Delete(ID)
}

func (f *roleApp) GetByID(ID int) (*entity.Role, error) {
	return f.request.GetByID(ID)
}

func (f *roleApp) GetAllPagination(postsPerPage int, offset int) ([]entity.Role, error) {
	return f.request.GetAllPagination(postsPerPage, offset)
}

//Update service init
func (f *roleApp) Update(roles *entity.Role) (*entity.Role, map[string]string) {
	return f.request.Update(roles)
}
func (f *roleApp) UpdateTitle(id int, title string) {
	f.request.UpdateTitle(id, title)
}
