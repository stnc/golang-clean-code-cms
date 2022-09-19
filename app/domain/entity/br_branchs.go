package entity

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	tr_translations "gopkg.in/go-playground/validator.v9/translations/tr"
)

//Branches strcut
type Branches struct {
	ID                   uint64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID               uint64     `gorm:"not null;" json:"userId"`
	RegionId             int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required"`
	Title                string     `gorm:"type:varchar(255);not null;" json:"title" validate:"required"`
	BranchCode           string     `gorm:"type:varchar(255);not null;" json:"branch_code" validate:"required"`
	Description          string     `gorm:"type:text;" json:"description"`
	Picture              string     `gorm:"type:varchar(255);not null;" json:"picture" validate:"required"`
	ManagerName          string     `gorm:"type:varchar(255);not null;" json:"manager_name" validate:"required"`
	ManagerPhone         string     `gorm:"type:varchar(255);not null;" json:"manager_phone" validate:"required"`
	ManagerMail          string     `gorm:"type:varchar(255);not null;" json:"manager_mail" validate:"required"`
	Html                 string     `gorm:"type:text;" json:"html"`
	Status               int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required"`
	RootBranchCategoryId int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required"`
	CreatedAt            time.Time  ` json:"createdAt"`
	UpdatedAt            time.Time  ` json:"updatedAt"`
	DeletedAt            *time.Time `json:"deletedAt"`
	SelectedID           uint64     `gorm:"-"` // ignore this field when write and read

}

//BranchesSave strcut
type BranchSaveDTO struct {
	ID               uint64 `gorm:"primary_key;auto_increment" json:"id"`
	ParentCategoryID uint64 `gorm:"not null;DEFAULT:'0'" json:"parentCategoryID"`
	Title            string `gorm:"size:100 ;not null;" json:"title" validate:"required"`
	Description      string `gorm:"type:text ;" json:"description"`
	Slug             string `gorm:"size:255 ;null;" json:"slug"`
	SelectedID       uint64
	PostType         int `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required" json:"postType"`
}

//BeforeSave init
func (f *Branches) BeforeSave() {
	f.Title = html.EscapeString(strings.TrimSpace(f.Title))
}

//TableName override
func (gk *Branches) TableName() string {
	return "br_branches"
}

//Prepare init
func (f *Branches) Prepare() {
	f.Title = html.EscapeString(strings.TrimSpace(f.Title))
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
}

//ValidateV1 old version
func (f *Branches) ValidateV1() map[string]string {
	var errorMessages = make(map[string]string)

	if f.Title == "" || f.Title == "null" {
		errorMessages["PostTitle_required"] = "PostPostTitle is required"
	}
	if f.Description == "" || f.Description == "null" {
		errorMessages["desc_required"] = "content is required"
	}

	return errorMessages
}

//Validate fluent validation
func (f *Branches) Validate() map[string]string {
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
