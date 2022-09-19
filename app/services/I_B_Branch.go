package services

import (
	"stncCms/app/domain/entity"
)

//BranchAppInterface service
type BranchAppInterface interface {
	Save(*entity.Branches) (*entity.Branches, map[string]string)
	GetByID(uint64) (*entity.Branches, error)
	GetAll() ([]entity.Branches, error)
	Update(*entity.Branches) (*entity.Branches, map[string]string)
	Delete(uint64) error
}

//BranchApp struct  init
type BranchApp struct {
	request BranchAppInterface
}

var _ BranchAppInterface = &BranchApp{}

//Save service init
func (f *BranchApp) Save(Cat *entity.Branches) (*entity.Branches, map[string]string) {
	return f.request.Save(Cat)
}

//GetAll service init
func (f *BranchApp) GetAll() ([]entity.Branches, error) {
	return f.request.GetAll()
}

//GetByID service init
func (f *BranchApp) GetByID(CatID uint64) (*entity.Branches, error) {
	return f.request.GetByID(CatID)
}

//Update service init
func (f *BranchApp) Update(Cat *entity.Branches) (*entity.Branches, map[string]string) {
	return f.request.Update(Cat)
}

//Delete service init
func (f *BranchApp) Delete(CatID uint64) error {
	return f.request.Delete(CatID)
}
