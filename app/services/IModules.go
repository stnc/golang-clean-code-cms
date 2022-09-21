package services

import (
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
)

//ModuleAppInterface interface
type ModulesAppInterface interface {
	GetAll() ([]entity.Modules, error)
	GetAllModulesMerge() ([]dto.ModulesAndPermission, error)
	GetAllModulesMergePermission() ([]dto.ModulesAndPermissionRole, error)
}

type moduleApp struct {
	request ModulesAppInterface
}

//UserApp implements the UserAppInterface
var _ ModulesAppInterface = &moduleApp{}

func (f *moduleApp) GetAll() ([]entity.Modules, error) {
	return f.request.GetAll()
}

func (f *moduleApp) GetAllModulesMerge() ([]dto.ModulesAndPermission, error) {
	return f.request.GetAllModulesMerge()
}
func (f *moduleApp) GetAllModulesMergePermission() ([]dto.ModulesAndPermissionRole, error) {
	return f.request.GetAllModulesMergePermission()
}
