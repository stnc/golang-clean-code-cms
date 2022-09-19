package controller

import (
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
	"stncCms/app/domain/cache"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/lang"
	"stncCms/app/domain/helpers/stncsession"
	repository "stncCms/app/domain/repository/cacheRepository"
)

const viewPathIndex = "admin/index/"

// index
func Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	locale, menuLanguage := lang.LoadLanguages("dashboard")
	viewData := pongo2.Context{
		"title":       locale.Get("Dashboard"),
		"flashMsg":    flashMsg,
		"csrf":        csrf.GetToken(c),
		"locale":      locale,
		"localeMenus": menuLanguage,
	}
	c.HTML(
		http.StatusOK,
		viewPathIndex+"dashboard.html",
		viewData,
	)
}

// OptionsDefault all list f
// TODO: bu kisim baska isimle acilibilir
func OptionsDefault(c *gin.Context) {
	// stncsession.IsLoggedInRedirect(c)

	//buraya bir oprion otılacak bunlar giriş yaptıktan sonra veri varmı yok mu bakacak

	db := repository.DB

	option1 := entity.Options{OptionName: "siteUrl", OptionValue: "http://localhost:8888/"}
	db.Debug().Create(&option1)

	option2 := entity.Options{OptionName: "kurban_yili", OptionValue: "2022"}
	db.Debug().Create(&option2)

	option3 := entity.Options{OptionName: "hisse_adeti", OptionValue: "7"}
	db.Debug().Create(&option3)

	option4 := entity.Options{OptionName: "satis_birim_fiyati_1", OptionValue: "20"}
	db.Debug().Create(&option4)

	option5 := entity.Options{OptionName: "satis_birim_fiyati_2", OptionValue: "25"}
	db.Debug().Create(&option5)

	option6 := entity.Options{OptionName: "satis_birim_fiyati_3", OptionValue: "30"}
	db.Debug().Create(&option6)

	option7 := entity.Options{OptionName: "hayvan_dusuk_agirligi", OptionValue: "0-200"}
	db.Debug().Create(&option7)

	option78 := entity.Options{OptionName: "hayvan_orta_agirligi", OptionValue: "200-600"}
	db.Debug().Create(&option78)

	option786 := entity.Options{OptionName: "hayvan_yuksek_agirligi", OptionValue: "600-1500"}
	db.Debug().Create(&option786)

	option8 := entity.Options{OptionName: "alis_birim_fiyati_1", OptionValue: "10"}
	db.Debug().Create(&option8)

	option9 := entity.Options{OptionName: "alis_birim_fiyati_2", OptionValue: "15"}
	db.Debug().Create(&option9)

	option10 := entity.Options{OptionName: "alis_birim_fiyati_3", OptionValue: "20"}
	db.Debug().Create(&option10)

	option11 := entity.Options{OptionName: "otomatik_sira_buyukbas", OptionValue: "1"}
	db.Debug().Create(&option11)

	option12 := entity.Options{OptionName: "otomatik_sira_kuyukbas", OptionValue: "1"}
	db.Debug().Create(&option12)

	option13 := entity.Options{OptionName: "whatsAppMsg", OptionValue: "Merhaba ."}
	db.Debug().Create(&option13)

	option14 := entity.Options{OptionName: "whatsAppMsgMap", OptionValue: "Merhaba Efendim bize bu adresden ulaşın "}
	db.Debug().Create(&option14)

	option15 := entity.Options{OptionName: "whatsAppMsgEk1", OptionValue: "ek mesaj "}
	db.Debug().Create(&option15)

	//mutluerF9E
	user := entity.Users{FirstName: "Sel", LastName: "t", Email: "i.com", Password: ""} //mutluerF9E
	db.Debug().Create(&user)

	ModulKurban := entity.Modules{ModulName: "kurban", Status: 1, UserID: 1}
	db.Debug().Create(&ModulKurban)

	Modulodemeler := entity.Modules{ModulName: "odemeler", Status: 1, UserID: 1}
	db.Debug().Create(&Modulodemeler)

	ModulHayvanBilgisi := entity.Modules{ModulName: "hayvanBilgisi", Status: 1, UserID: 1}
	db.Debug().Create(&ModulHayvanBilgisi)

	ModulDashborad := entity.Modules{ModulName: "dashboard", Status: 1, UserID: 1}
	db.Debug().Create(&ModulDashborad)

	ModulAyarlar := entity.Modules{ModulName: "ayarlar", Status: 1, UserID: 1}
	db.Debug().Create(&ModulAyarlar)

	ModulGruplar := entity.Modules{ModulName: "gruplar", Status: 1, UserID: 1}
	db.Debug().Create(&ModulGruplar)

	ModulhayvanSatisYerleri := entity.Modules{ModulName: "hayvanSatisYerleri", Status: 1, UserID: 1}
	db.Debug().Create(&ModulhayvanSatisYerleri)

	ModulKisiler := entity.Modules{ModulName: "kisiler", Status: 1, UserID: 1}
	db.Debug().Create(&ModulKisiler)

	Modulkullanici := entity.Modules{ModulName: "kullanici", Status: 1, UserID: 1}
	db.Debug().Create(&Modulkullanici)

	// db.Debug().Delete(&entity.Kisiler{}, 1)

	c.JSON(http.StatusOK, "yapıldı")
}

// OptionsDefault all list f
func CacheReset(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	stncsession.SetFlashMessage("Cache Temizlendi", "success", c)
	redisClient := cache.RedisDBInit()
	redisClient.FlushAll()
	c.Redirect(http.StatusMovedPermanently, "/admin/options/")
}
