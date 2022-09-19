package services

import (
	"stncCms/app/domain/entity"
)

//PermissionAppInterface interface
type RolePermissionAppInterface interface {
	GetAll() ([]entity.RolePermisson, error)
	Save(*entity.RolePermisson) (*entity.RolePermisson, map[string]string)
	Update(*entity.RolePermisson) (*entity.RolePermisson, map[string]string)
	UpdateActiveStatus(roleId int, permissionId int, active int)
}

type rolePermissionApp struct {
	request RolePermissionAppInterface
}

//UserApp implements the UserAppInterface
var _ RolePermissionAppInterface = &rolePermissionApp{}

func (f *rolePermissionApp) GetAll() ([]entity.RolePermisson, error) {
	return f.request.GetAll()
}

func (f *rolePermissionApp) Save(data *entity.RolePermisson) (*entity.RolePermisson, map[string]string) {
	return f.request.Save(data)
}
func (f *rolePermissionApp) Update(data *entity.RolePermisson) (*entity.RolePermisson, map[string]string) {
	return f.request.Update(data)
}

func (f *rolePermissionApp) UpdateActiveStatus(roleId int, permissionId int, active int) {
	f.request.UpdateActiveStatus(roleId, permissionId, active)
}
