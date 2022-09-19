package controller

import (
	"net/http"
	"stncCms/app/domain/helpers/lang"
	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/services"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// Options constructor
type Options struct {
	OptionsApp services.OptionsAppInterface
}

var (
	paginator = &pagination.Paginator{}
)

const viewPathOptions = "admin/options/"

// InitOptions post controller constructor
func InitOptions(OptionsApp services.OptionsAppInterface) *Options {
	return &Options{
		OptionsApp: OptionsApp,
	}
}

// option list data
func (access *Options) Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	locale, menuLanguage := lang.LoadLanguages("options")
	numberOfShares := access.OptionsApp.GetOption("hisse_adeti")
	salesUnitPrice1 := access.OptionsApp.GetOption("satis_birim_fiyati_1")
	salesUnitPrice2 := access.OptionsApp.GetOption("satis_birim_fiyati_2")
	salesUnitPrice3 := access.OptionsApp.GetOption("satis_birim_fiyati_3")
	purchaseUnitPrice1 := access.OptionsApp.GetOption("alis_birim_fiyati_1")
	purchaseUnitPrice2 := access.OptionsApp.GetOption("alis_birim_fiyati_2")
	purchaseUnitPrice3 := access.OptionsApp.GetOption("alis_birim_fiyati_3")
	whatsAppMsg := access.OptionsApp.GetOption("whatsAppMsg")
	whatsAppMsgMap := access.OptionsApp.GetOption("whatsAppMsgMap")
	whatsAppMsgEk1 := access.OptionsApp.GetOption("whatsAppMsgEk1")
	otomatik_sira_buyukbas := access.OptionsApp.GetOption("otomatik_sira_buyukbas")
	receipt_otomatik_sira_no := access.OptionsApp.GetOption("receipt_otomatik_sira_no")
	cache_open_close := access.OptionsApp.GetOption("cache_open_close")
	// dusukagirlik := access.OptionsApp.GetOption("hayvan_dusuk_agirligi")
	// ortaagirlik := access.OptionsApp.GetOption("hayvan_orta_agirligi")
	// yuksekagirlik := access.OptionsApp.GetOption("hayvan_yuksek_agirligi")
	yearSacrifice := access.OptionsApp.GetOption("kurban_yili")

	viewData := pongo2.Context{
		"title":                    "Ayarlar",
		"csrf":                     csrf.GetToken(c),
		"numberOfShares":           numberOfShares,
		"satis_birim_fiyati_1":     salesUnitPrice1,
		"satis_birim_fiyati_2":     salesUnitPrice2,
		"satis_birim_fiyati_3":     salesUnitPrice3,
		"alis_birim_fiyati_1":      purchaseUnitPrice1,
		"alis_birim_fiyati_2":      purchaseUnitPrice2,
		"alis_birim_fiyati_3":      purchaseUnitPrice3,
		"whatsAppMsg":              whatsAppMsg,
		"whatsAppMsgEk1":           whatsAppMsgEk1,
		"otomatik_sira_buyukbas":   otomatik_sira_buyukbas,
		"receipt_otomatik_sira_no": receipt_otomatik_sira_no,
		"cache_open_close":         cache_open_close,
		// "hayvan_dusuk_agirligi":  dusukagirlik,
		// "hayvan_orta_agirligi":   ortaagirlik,
		// "hayvan_yuksek_agirligi": yuksekagirlik,
		"kurban_yili":    yearSacrifice,
		"flashMsg":       flashMsg,
		"whatsAppMsgMap": whatsAppMsgMap,
		"locale":         locale,
		"localeMenus":    menuLanguage,
	}

	c.HTML(
		http.StatusOK,
		viewPathOptions+"index.html",
		viewData,
	)
}

// Update list
func (access *Options) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	access.OptionsApp.SetOption("hisse_adeti", c.PostForm("hisse_adeti"))
	access.OptionsApp.SetOption("satis_birim_fiyati_1", c.PostForm("satis_birim_fiyati_1"))
	access.OptionsApp.SetOption("satis_birim_fiyati_2", c.PostForm("satis_birim_fiyati_2"))
	access.OptionsApp.SetOption("satis_birim_fiyati_3", c.PostForm("satis_birim_fiyati_3"))
	access.OptionsApp.SetOption("alis_birim_fiyati_1", c.PostForm("alis_birim_fiyati_1"))
	access.OptionsApp.SetOption("alis_birim_fiyati_2", c.PostForm("alis_birim_fiyati_2"))
	access.OptionsApp.SetOption("alis_birim_fiyati_3", c.PostForm("alis_birim_fiyati_3"))
	access.OptionsApp.SetOption("whatsAppMsg", c.PostForm("whatsAppMsg"))
	access.OptionsApp.SetOption("whatsAppMsgMap", c.PostForm("whatsAppMsgMap"))
	access.OptionsApp.SetOption("whatsAppMsgEk1", c.PostForm("whatsAppMsgEk1"))
	// access.OptionsApp.SetOption("hayvan_dusuk_agirligi", c.PostForm("hayvan_dusuk_agirligi"))
	// access.OptionsApp.SetOption("hayvan_orta_agirligi", c.PostForm("hayvan_orta_agirligi"))
	// access.OptionsApp.SetOption("hayvan_yuksek_agirligi", c.PostForm("hayvan_yuksek_agirligi"))
	access.OptionsApp.SetOption("kurban_yili", c.PostForm("kurban_yili"))
	access.OptionsApp.SetOption("otomatik_sira_buyukbas", c.PostForm("otomatik_sira_buyukbas"))
	access.OptionsApp.SetOption("receipt_otomatik_sira_no", c.PostForm("receipt_otomatik_sira_no"))
	access.OptionsApp.SetOption("cache_open_close", c.PostForm("cache_open_close"))
	stncsession.SetStore(c, "cache_open_close", c.PostForm("cache_open_close"))

	stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
	c.Redirect(http.StatusMovedPermanently, "/admin/options")
	return

}

// Receipt No generator  tr: makbuz no üretir
func (access *Options) ReceiptNo(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	receiptAutomaticQueueNo := access.OptionsApp.GetOption("receipt_automatic_queue_no ")
	receiptAutomaticQueueNoint := stnccollection.StringToint(receiptAutomaticQueueNo) + 1
	receiptAutomaticQueueNostr := stnccollection.IntToString(receiptAutomaticQueueNoint)
	access.OptionsApp.SetOption("receipt_automatic_queue_no", receiptAutomaticQueueNostr)
	viewData := pongo2.Context{
		"title":                   "Otomatik Sıra No",
		"status":                  "ok",
		"receiptAutomaticQueueNo": receiptAutomaticQueueNo,
		"errMsg":                  "",
	}
	c.JSON(http.StatusOK, viewData)
	return

}
