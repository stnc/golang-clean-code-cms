package services

import (
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
)

//PermissionAppInterface interface
type PermissionAppInterface interface {
	GetAll() ([]entity.Permission, error)
	GetAllPaginationermissionForModulID(modulId int) ([]entity.Permission, error)
	GetUserPermission(roleID int) ([]dto.RbcaCheck, error)
}

type permissionApp struct {
	request PermissionAppInterface
}

//UserApp implements the UserAppInterface
var _ PermissionAppInterface = &permissionApp{}

func (f *permissionApp) GetAll() ([]entity.Permission, error) {
	return f.request.GetAll()
}

func (f *permissionApp) GetAllPaginationermissionForModulID(modulId int) ([]entity.Permission, error) {
	return f.request.GetAllPaginationermissionForModulID(modulId)
}

func (f *permissionApp) GetUserPermission(roleID int) ([]dto.RbcaCheck, error) {
	return f.request.GetUserPermission(roleID)
}
