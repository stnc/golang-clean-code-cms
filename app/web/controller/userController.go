package controller

import (
	"net/http"
	"stncCms/app/domain/helpers/lang"
	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/infrastructure/security"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stncdatetime"
	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/services"
	"strconv"
)

//https://github.com/mailgun/mailgun-go
//https://testmail.app/pricing
//https://mailtrap.io/pricing/
// https://sendgrid.com/pricing/?sg_product=ei
//https://www.sendinblue.com/pricing/
//Postmark’, ‘SparkPost’ ve ‘Ses Driver’
//https://birhankarahasan.com/laravel-mailgun-kullanimi-nasil-mail-gonderilir

// userControl constructor
type UserControl struct {
	UserControlApp services.UserAppInterface
	Branch         services.BranchAppInterface
}

const viewPathuserControl = "admin/user/"

// InitUserControl userControl controller constructor
func InitUserControl(KiApp services.UserAppInterface, BranchApp services.BranchAppInterface) *UserControl {
	return &UserControl{
		UserControlApp: KiApp,
		Branch:         BranchApp,
	}
}

// Index list
func (access *UserControl) Index(c *gin.Context) {
	// alluserControl, err := access.userControlApp.GetAlluserControl()

	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("user")
	flashMsg := stncsession.GetFlashMessage(c)

	var tarih stncdatetime.Inow
	var total int64
	access.UserControlApp.Count(&total)
	userControlsPerPage := 10
	paginator := pagination.NewPaginator(c.Request, userControlsPerPage, total)
	offset := paginator.Offset()

	userControls, _ := access.UserControlApp.GetAllPagination(userControlsPerPage, offset)

	// var tarih stncdatetime.Inow
	// fmt.Println(tarih.TarihFullSQL("2020-05-21 05:08"))
	// fmt.Println(tarih.AylarListe("May"))
	// fmt.Println(tarih.Tarih())
	// //	tarih.FormatTarihForMysql("2020-05-17 05:08:40")
	//	tarih.FormatTarihForMysql("2020-05-17 05:08:40")

	viewData := pongo2.Context{
		"paginator":   paginator,
		"title":       "İçerik Ekleme",
		"allData":     userControls,
		"tarih":       tarih,
		"flashMsg":    flashMsg,
		"csrf":        csrf.GetToken(c),
		"locale":      locale,
		"localeMenus": menuLanguage,
	}

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		viewPathuserControl+"index.html",
		// Pass the data that the page uses
		viewData,
	)
}

// Create all list f
func (access *UserControl) Create(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("user")
	flashMsg := stncsession.GetFlashMessage(c)
	cats, _ := access.Branch.GetAll()

	viewData := pongo2.Context{
		"title":       "İçerik Ekleme",
		"catsData":    cats,
		"flashMsg":    flashMsg,
		"csrf":        csrf.GetToken(c),
		"locale":      locale,
		"localeMenus": menuLanguage,
	}
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		viewPathuserControl+"create.html",
		// Pass the data that the page uses
		viewData,
	)
}

// Store save method
func (access *UserControl) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("user")
	flashMsg := stncsession.GetFlashMessage(c)

	var userSave = userModel(c)
	var userSavePostError = make(map[string]string)
	userSavePostError = userSave.Validate()

	if len(userSavePostError) == 0 {
		saveData, saveErr := access.UserControlApp.Save(&userSave)
		if saveErr != nil {
			userSavePostError = saveErr
		}
		lastID := strconv.FormatUint(uint64(saveData.ID), 10)
		stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/user/edit/"+lastID)
		return
	} else {
		stncsession.SetFlashMessage("Zorunlu alanları lütfen doldurunuz", "danger", c)
	}
	viewData := pongo2.Context{
		"title":       "content add",
		"csrf":        csrf.GetToken(c),
		"err":         userSavePostError,
		"data":        userSave,
		"flashMsg":    flashMsg,
		"locale":      locale,
		"localeMenus": menuLanguage,
	}
	c.HTML(
		http.StatusOK,
		viewPathuserControl+"create.html",
		viewData,
	)

}

// Edit edit data
func (access *UserControl) Edit(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("user")
	flashMsg := stncsession.GetFlashMessage(c)
	if userID, err := strconv.ParseUint(c.Param("UserID"), 10, 64); err == nil {
		if userControls, err := access.UserControlApp.GetByID(userID); err == nil {
			viewData := pongo2.Context{
				"title":       "kullanıcı düzenleme",
				"data":        userControls,
				"csrf":        csrf.GetToken(c),
				"flashMsg":    flashMsg,
				"locale":      locale,
				"localeMenus": menuLanguage,
			}
			c.HTML(
				http.StatusOK,
				viewPathuserControl+"edit.html",
				viewData,
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// Update data
// TODO: resim silmeyi unutma
// Delete data
func (access *UserControl) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	// var userControl, _, id = userControlModel(c)
	// var saveuserControlError = make(map[string]string)

	// if len(saveuserControlError) == 0 {

	// 	// user bilgisi guncelleyecekılerde acabilirim
	// 	_, saveErr := access.UserControlApp.UpdateDto(&userControl)
	// 	if saveErr != nil {
	// 		saveuserControlError = saveErr
	// 	}

	// 	stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)

	// 	c.Redirect(http.StatusModPermanently, "/"+viewPathuserControl+"edit/"+id)
	// 	return:= pongo2.Context{
	// 	"title": "Kullanıcı düzenlme
	// 	"err":   saveuserControlEr,
	//	"csrf":  csrf.GetToken(c),

	// 	ata": userControl,
	// }
	// .HTML(
	// 	ttp.SttusOK,
	// 	viewPahurCotrol+"edit.html",
	// 	ewData,
	// )
}

// OdemeEkleCreateModalBox takistler
func (access *UserControl) NewPasswordModalBox(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	viewData := pongo2.Context{
		"title": "New Password",
		"csrf":  csrf.GetToken(c),
	}
	c.HTML(
		http.StatusOK,
		viewPathuserControl+"NewPasswordModalBox.html",
		viewData,
	)
}

// referansEkleAjax save method
func (access *UserControl) NewPasswordCreateModalBox(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, _ := lang.LoadLanguages("user")
	//var KurbanID uint64
	//var kalanKurbanFiyati float64
	//KurbanID, _ = strconv.ParseUint(c.PostForm("kurbanID"), 10, 64)
	//kalanKurbanFiyati = access.kurbanApp.KalanUcret(KurbanID)
	//db := repository.DB
	//appKurban := repository.KurbanRepositoryInit(db)
	//var kisiID int
	//appKurban.GetKurbanKisiVarmi(KurbanID, &kisiID)
	//if kisiID == 1 {
	// sahte veri girişi yani kişi atanmamış kurbana ödeme yapmaya çalışıyor  TODO: bunun loglanması lazım
	viewData := pongo2.Context{
		"title":  "Kurban Ekleme",
		"csrf":   csrf.GetToken(c),
		"status": "err",
		"err":    "fk", // sahte veri girişi TODO: bunun loglanması lazım
		"errMsg": "beklenmeyen bir hata oluştu",
		"locale": locale,
	}
	c.JSON(http.StatusOK, viewData)
	//} else {
	//	if odeme, _, errorR := odemelerModel(kalanKurbanFiyati, c); errorR == nil {
	//		var savePostError = make(map[string]string)
	//		savePostError = odeme.Validate()
	//		fmt.Printf("%+v\n", odeme)
	//		if len(savePostError) == 0 {
	//			_, saveErr := access.OdemelerApp.Save(&odeme)
	//			stncsession.SetFlashMessage("Password changed successfully", "success", c)
	//			if saveErr != nil {
	//				savePostError = saveErr
	//			}
	//			viewData := pongo2.Context{
	//				"title":  "Password Change",
	//				"csrf":   csrf.GetToken(c),
	//				"err":    savePostError,
	//				"status": "ok",
	//				"path":   "/admin/kurban/edit/" + c.PostForm("kurbanID"),
	//				"id":     c.PostForm("kurbanID"),
	//				"post":   odeme,
	//			}
	//			c.JSON(http.StatusOK, viewData)
	//
	//		}
	//	}
	//}
}

// form post model
func userModel(c *gin.Context) (user entity.Users) {
	//	var post = entit.Post{}
	user.Username = c.PostForm("UserNameNew")
	user.Email = c.PostForm("EmailNew")
	user.FirstName = c.PostForm("FirstName")
	user.LastName = c.PostForm("LastName")
	user.Phone = c.PostForm("Phone")
	pass := c.PostForm("PasswordNew")
	sifre := security.PassGenerate(pass)
	user.Password = sifre
	user.RoleID = stnccollection.StringToint(c.PostForm("RoleID"))
	return user
}
