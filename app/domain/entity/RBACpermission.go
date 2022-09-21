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

//Role Based Access Control (RBAC)

//Permission
type Permission struct {
	ID          int        `gorm:"primary_key;auto_increment" json:"id"`
	ModulID     int        `gorm:"not null;" json:"modulID"`
	Title       string     `gorm:"type:varchar(75) ;not null;" json:"title"`
	Description string     `gorm:"type:varchar(255);" json:"description"`
	Controller  string     `gorm:"type:varchar(255);" json:"controller"`
	FuncName    string     `gorm:"type:varchar(255);" json:"func_name"`
	Context     string     `gorm:"type:text ;" json:"context"`
	Active      int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" json:"active" `
	CreatedAt   time.Time  ` json:"createdAt"`
	UpdatedAt   time.Time  ` json:"updatedAt"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

//TableName override
func (gk *Permission) TableName() string {
	return "rbca_permission"
}

//BeforeSave is a gorm hook
func (u *Permission) BeforeSave() error {
	u.Title = html.EscapeString(strings.TrimSpace(u.Title))
	u.Description = html.EscapeString(strings.TrimSpace(u.Description))
	u.Context = html.EscapeString(strings.TrimSpace(u.Context))
	return nil
}

func (u *Permission) Prepare() {
	u.Title = html.EscapeString(strings.TrimSpace(u.Title))
	u.Description = html.EscapeString(strings.TrimSpace(u.Description))
	u.Context = html.EscapeString(strings.TrimSpace(u.Context))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

//Validate fluent validation
func (f *Permission) Validate() map[string]string {
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
