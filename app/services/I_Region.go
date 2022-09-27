package services

import (
	"stncCms/app/domain/entity"
)

// RegionAppInterface service
type RegionAppInterface interface {
	Save(*entity.Region) (*entity.Region, map[string]string)
	GetByID(uint64) (*entity.Region, error)

	GetAll() ([]entity.Region, error)
	Update(*entity.Region) (*entity.Region, map[string]string)
	Delete(uint64) error
}

// RegionApp struct  init
type RegionApp struct {
	request RegionAppInterface
}

var _ RegionAppInterface = &RegionApp{}

// Save service init
func (f *RegionApp) Save(Cat *entity.Region) (*entity.Region, map[string]string) {
	return f.request.Save(Cat)
}

// GetAll service init
func (f *RegionApp) GetAll() ([]entity.Region, error) {
	return f.request.GetAll()
}

// GetByID service init
func (f *RegionApp) GetByID(catID uint64) (*entity.Region, error) {
	return f.request.GetByID(catID)
}

// Update service init
func (f *RegionApp) Update(cat *entity.Region) (*entity.Region, map[string]string) {
	return f.request.Update(cat)
}

// Delete service init
func (f *RegionApp) Delete(catID uint64) error {
	return f.request.Delete(catID)
}
