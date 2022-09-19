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

//TODO: ilerde sub Region gelebilir

//Region strcut
type Region struct {
	ID               uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID           uint64 `gorm:"not null;" json:"userId"`
	ParentCategoryID uint64 `gorm:"not null;DEFAULT:'0'" json:"parent_categoryId"`
	Name             string `gorm:"type:varchar(255);not null;" json:"name" validate:"required"`
	Description      string `gorm:"type:text;" json:"description"`
	Slug             string `gorm:"size:255;null;" json:"slug"`
	SelectedID       uint64 `gorm:"-"` // ignore this field when write and read
	//Picture      string     `gorm:"size:255;null;" json:"picture" `
	CreatedAt time.Time  ` json:"createdAt"`
	UpdatedAt time.Time  ` json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

//RegionSave strcut
type RegionSaveDTO struct {
	ID               uint64 `gorm:"primary_key;auto_increment" json:"id"`
	ParentCategoryID uint64 `gorm:"not null;DEFAULT:'0'" json:"parentCategoryID"`
	Name             string `gorm:"size:100 ;not null;" json:"name" validate:"required"`
	Description      string `gorm:"type:text ;" json:"description"`
	Slug             string `gorm:"size:255 ;null;" json:"slug"`
	SelectedID       uint64 `json:"selectedID"`
	PostType         int    `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" json:"postType" validate:"required"`
}

//BeforeSave init
func (f *Region) BeforeSave() {
	f.Name = html.EscapeString(strings.TrimSpace(f.Name))
}

//Prepare init
func (f *Region) Prepare() {
	f.Name = html.EscapeString(strings.TrimSpace(f.Name))
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
}

//TableName override
func (gk *Region) TableName() string {
	return "br_region"
}

//ValidateV1 old version
func (f *Region) ValidateV1() map[string]string {
	var errorMessages = make(map[string]string)

	if f.Name == "" || f.Name == "null" {
		errorMessages["PostTitle_required"] = "PostPostTitle is required"
	}
	if f.Description == "" || f.Description == "null" {
		errorMessages["desc_required"] = "content is required"
	}

	return errorMessages
}

//Validate fluent validation
func (f *Region) Validate() map[string]string {
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
