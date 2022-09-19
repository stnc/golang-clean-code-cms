package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	tr_translations "gopkg.in/go-playground/validator.v9/translations/tr"
)

//Role Based Access Control (RBAC)

//RolePermissonTableName global table name
var RolePermissonTableName string = "rbca_role_permisson"

//RolePermisson
type RolePermisson struct {
	ID           int        `gorm:"primary_key;auto_increment" json:"id"`
	RoleID       int        `gorm:"not null;" json:"roleID"`
	PermissionID int        `gorm:"not null;" json:"permissionID"`
	Active       int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'0'" validate:"required" json:"active"`
	Status       int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required" json:"status"`
	CreatedAt    time.Time  ` json:"createdAt"`
	UpdatedAt    time.Time  ` json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
}

func (gk *RolePermisson) TableName() string {
	return RolePermissonTableName
}

//Validate fluent validation
func (f *RolePermisson) Validate() map[string]string {
	var (
		validate *validator.Validate
		uni      *ut.UniversalTranslator
	)
	tr := en.New()
	uni = ut.New(tr, tr)
	trans, _ := uni.GetTranslator("tr")
	validate = validator.New()
	tr_translations.RegisterDefaultTranslations(validate, trans)

	errorLog := make(map[string]string)

	err := validate.Struct(f)
	fmt.Println(err)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs)
		for _, e := range errs {
			// can translate each error one at a time.
			lng := strings.Replace(e.Translate(trans), e.Field(), "Burası", 1)
			errorLog[e.Field()+"_error"] = e.Translate(trans)
			// errorLog[e.Field()] = e.Translate(trans)
			errorLog[e.Field()] = lng
			errorLog[e.Field()+"_valid"] = "is-invalid"
		}
	}
	return errorLog
}
