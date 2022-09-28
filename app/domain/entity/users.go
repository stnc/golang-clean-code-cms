package entity

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	tr_translations "gopkg.in/go-playground/validator.v9/translations/tr"
)

type UsersData []Users

// User
type Users struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID      uint64 `gorm:"not null;" json:"userId"`
	RoleID      int    ` json:"roleId" validate:"required"`
	Username    string `gorm:"type:varchar(255) ;not null;" json:"username"`
	FirstName   string `gorm:"type:varchar(255) ;not null;" json:"firstName"`
	LastName    string `gorm:"type:varchar(255) ;not null;" json:"lastName"`
	Email       string `gorm:"type:varchar(255) ;" validate:"email"  json:"emailAdres"` //`gorm:"type:varchar(255) ;" validate:"required,email"  json:"emailAdres"`
	Password    string `gorm:"type:varchar(255) ;column:password" validate:"required"  json:"password"`
	TimeZone    string `gorm:"type:varchar(255) ;column:time_zone" json:"timezone"  `
	Description string `gorm:"type:text ;" json:"description"`

	PasswordReset string `gorm:"type:varchar(255) ;column:password_reset"  `
	Phone         string `gorm:"type:varchar(255) ;column:phone" validate:"required" `
	LastLogin     string `gorm:"type:varchar(255) ;column:last_login"  `
	Status        int    `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'"  json:"status" `
	Activation    int    `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" json:"activation" `
	BranchID      int    `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" json:"branchID" `

	CreatedAt time.Time  ` json:"createdAt"`
	UpdatedAt time.Time  ` json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type UsersGetByUserForBranchIDDTO struct {
	//ID         uint64 `json:"id"`
	//UserID     uint64 `json:"userId"`
	//Username   string `json:"username"`
	//FirstName  string `json:"first_name"`
	//LastName   string `json:"last_name"`
	BranchID   uint64 `json:"branchID"`
	RegionName string `json:"regionName"`
	BranchName string `json:"branchName"`
	RegionID   uint64 `json:"regionID"`
}

// PublicUser
type PublicUser struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `gorm:"size:100;not null;" json:"firstName"`
	LastName  string `gorm:"size:100;not null;" json:"lastName"`
}

// BeforeSave is a gorm hook
func (u *Users) BeforeSave() error {
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	// hashPassword := security.Hash(u.Password)
	// u.Password = string(hashPassword)
	return nil
}

// So that we dont expose the user's email address and password to the world
func (sers UsersData) PublicUsers() []interface{} {
	result := make([]interface{}, len(sers))
	for index, user := range sers {
		result[index] = user.PublicUser()
	}
	return result
}

// So that we dont expose the user's email address and password to the world
func (u *Users) PublicUser() interface{} {
	return &PublicUser{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

func (u *Users) Prepare() {
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// bu kısım  API den kullanılan onun için silinmeyecek kopya oluşturulablir
func (u *Users) ValidateAPI(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			errorMessages["email_required"] = "email required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "email email"
			}
		}

	case "login":
		if u.Password == "" {
			errorMessages["password_required"] = "password is required"
		}
		if u.Email == "" {
			errorMessages["email_required"] = "email is required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	case "forgotpassword":
		if u.Email == "" {
			errorMessages["email_required"] = "email required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	default:
		if u.FirstName == "" {
			errorMessages["firstname_required"] = "first name is required"
		}
		if u.LastName == "" {
			errorMessages["lastname_required"] = "last name is required"
		}
		if u.Password == "" {
			errorMessages["password_required"] = "password is required"
		}
		if u.Password != "" && len(u.Password) < 6 {
			errorMessages["invalid_password"] = "password should be at least 6 characters"
		}
		if u.Email == "" {
			errorMessages["email_required"] = "email is required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	}
	return errorMessages
}

// Validate form için
func (u *Users) ValidateLoginForm(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			errorMessages["Email"] = "email required"

		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["Email"] = "email formati"
			}
		}

	case "login":
		if u.Password == "" {
			errorMessages["Password"] = "password is required"
			errorMessages["Password_error"] = "email required"
			errorMessages["Password_valid"] = "is-invalid"
		}
		if u.Email == "" {
			errorMessages["Email"] = "email is required"
			errorMessages["Email_error"] = "email is required"
			errorMessages["Email_valid"] = "is-invalid"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["Email"] = "please provide a valid email"
				errorMessages["Email_error"] = "please provide a valid email"
				errorMessages["Email_valid"] = "is-invalid"
			}
		}
	case "forgotpassword":
		if u.Email == "" {
			errorMessages["Email"] = "email required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["Email"] = "please provide a valid email"
			}
		}
	default:
		if u.FirstName == "" {
			errorMessages["FirstName"] = "first name is required"
		}
		if u.LastName == "" {
			errorMessages["LastName"] = "last name is required"
		}
		if u.Password == "" {
			errorMessages["Password"] = "password is required"
		}
		if u.Password != "" && len(u.Password) < 6 {
			errorMessages["Password"] = "password should be at least 6 characters"
		}
		if u.Email == "" {
			errorMessages["Email"] = "email is required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["Email"] = "please provide a valid email"
			}
		}
	}
	return errorMessages
}

// Validate fluent validation
func (f *Users) Validate() map[string]string {
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
