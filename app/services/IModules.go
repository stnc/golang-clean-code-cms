package services

import (
	"stncCms/app/domain/entity"
)

// ModuleAppInterface interface
type ModulesAppInterface interface {
	GetAll() ([]entity.Modules, error)
}

type moduleApp struct {
	request ModulesAppInterface
}

// UserApp implements the UserAppInterface
var _ ModulesAppInterface = &moduleApp{}

func (f *moduleApp) GetAll() ([]entity.Modules, error) {
	return f.request.GetAll()
}
