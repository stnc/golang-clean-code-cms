package dbRepository

import (
	"errors"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
	"stncCms/app/services"
	"strings"

	"github.com/jinzhu/gorm"
)

type RoleRepo struct {
	db *gorm.DB
}

func RoleRepositoryInit(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db}
}

//RoleRepo implements the repository.RoleRepository interface
var _ services.RoleAppInterface = &RoleRepo{}

//GetAll all data
func (r *RoleRepo) GetAll() ([]entity.Role, error) {
	var datas []entity.Role
	var err error
	err = r.db.Debug().Order("created_at desc").Find(&datas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

//Save data
func (r *RoleRepo) Save(data *entity.Role) (*entity.Role, map[string]string) {
	dbErr := map[string]string{}
	var err error
	err = r.db.Debug().Create(&data).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "post title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return data, nil
}

func (r *RoleRepo) EditList(modulID int, roleID int) ([]dto.RoleEditList, error) {
	var err error
	var data []dto.RoleEditList
	/*
	  // update listeleme icin

	   SELECT role_permission.id AS role_Permission_ID,modul.id AS modul_id,role_permission.role_id,role_permission.permission_id,
	   permission.title,permission.slug,modul.modul_name,role_permission.active AS role_permission_active
	    FROM   rbca_role_permisson AS role_permission
	             INNER JOIN  rbca_permission AS  permission ON permission.id=role_permission.permission_id
	   INNER JOIN modules AS modul ON modul.id= permission.modul_id
	   INNER JOIN rbca_role AS role ON role.id=role_permission.role_id  ORDER BY modul.modul_name

	*/

	query := r.db.Table(entity.RolePermissonTableName + " AS role_permission")
	query = query.Select(`role_permission.id AS role_Permission_ID,permission.modul_id AS modul_id,role_permission.role_id,role_permission.permission_id,
	role_permission.active AS   role_permission_active,permission.title AS permission_Title,permission.slug AS permission_slug`)
	query = query.Joins("INNER JOIN  rbca_permission AS  permission ON permission.id=role_permission.permission_id ")
	query = query.Joins("INNER JOIN rbca_role AS role ON role.id=role_permission.role_id")
	query = query.Where("permission.modul_id=? ", modulID)
	query = query.Where("role_permission.role_id=? ", roleID)
	err = query.Find(&data).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return data, nil
}

//Count
func (r *RoleRepo) Count(totalCount *int64) {
	var data entity.Role
	var count int64
	r.db.Debug().Model(data).Count(&count)
	*totalCount = count
}

//Delete data
func (r *RoleRepo) Delete(id uint64) error {
	var data entity.Role
	var err error
	err = r.db.Debug().Where("id = ?", id).Delete(&data, id).Error
	// r.db.Debug().Where("role_id = ?", id).Delete(&entity.Role{})
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

//GetByID get data
func (r *RoleRepo) GetByID(id int) (*entity.Role, error) {
	var data entity.Role
	var err error
	err = r.db.Debug().Where("id = ?", id).Take(&data).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return &data, nil
}

//GetAllPagination pagination all data
func (r *RoleRepo) GetAllPagination(perPage int, offset int) ([]entity.Role, error) {
	var data []entity.Role
	var err error
	err = r.db.Debug().Limit(perPage).Offset(offset).Order("created_at desc").Find(&data).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return data, nil
}

//Update upate data
func (r *RoleRepo) Update(roles *entity.Role) (*entity.Role, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Update(&roles).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "Duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return roles, nil
}

func (r *RoleRepo) UpdateTitle(id int, title string) {
	r.db.Debug().Debug().Model(&entity.Role{}).Where("id = ?", id).Update("title", title)
}

/*
func (r *KurbanRepo) ListDataTableUserID(c *gin.Context, id string) (dto.KurbanBilgisiDataTableResult, error) {
	var total, filtered int
	var err error
	var data []dto.KurbanTable
	query := r.db.Table(entity.KurbanTableName)
	query = query.Select("sacrifice_kurbanlar.*, kisiler.ad_soyad as kisi_ad_soyad,   kisiler.telefon as kisi_telefon")
	query = query.Where("sacrifice_kurbanlar.user_id = ?", id)
	query = query.Joins(" join sacrifice_kisiler as kisiler on sacrifice_kurbanlar.kisi_id <> 1  and sacrifice_kurbanlar.kisi_id = kisiler.id ")
	// query = query.Joins(" join kisiler on kurbanlar.kisi_id <> 1  ")
	query = query.Offset(stnchelper.QueryOffset(c))
	query = query.Limit(stnchelper.QueryLimit(c))
	query = query.Order(r.queryOrder(c))
	query = query.Scopes(r.searchScope(c), stnchelper.DateTimeScope(c))
	err = query.Find(&data).Error

	query = query.Offset(0)

	query.Table(entity.KurbanTableName).Count(&filtered)
	// Total data count
	// r.db.Table(entity.KurbanTableName).Count(&total)
	query.Table(entity.KurbanTableName).Count(&total)

	result := dto.KurbanBilgisiDataTableResult{
		Total:    total,
		Filtered: filtered,
		Data:     data,
	}
	return result, err
}*/
