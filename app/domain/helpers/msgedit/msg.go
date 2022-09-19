package msgedit

import (
	"net/url"

	repository "stncCms/app/domain/repository/cacheRepository"
	"strings"
)

//oku
//https://faq.whatsapp.com/general/chats/how-to-use-click-to-chat/?lang=tr

type (
	// Inow : Custom Renderer for templates
	Msg struct{ Debug bool }
)

func (msgModul Msg) Wp4() bool {
	return SendMsgWp4()
}

func (msgModul Msg) Message1(phone string, slug string) string {
	return mesaage1Func(phone, slug)
}

func (msgModul Msg) Message2(phone string) string {
	return mesaage2Func(phone)
}

func (msgModul Msg) Message3(phone string) string {
	return mesaage3Func(phone)
}

func mesaage1Func(phone string, slug string) string {
	db := repository.DB
	tel := telReplace(phone)
	appOptions := repository.OptionRepositoryInit(db)
	msgg := msgReplace(appOptions.GetOption("whatsAppMsg"), slug, appOptions.GetOption("siteUrl"))
	// fmt.Println(appOptions.GetOption("siteUrl"))
	return textWhatsApp(tel, msgg)
}

func mesaage2Func(phone string) string {
	db := repository.DB
	tel := telReplace(phone)
	appOptions := repository.OptionRepositoryInit(db)
	msgg := appOptions.GetOption("whatsAppMsgMap")
	return textWhatsApp(tel, msgg)
}

func mesaage3Func(phone string) string {
	db := repository.DB
	tel := telReplace(phone)
	appOptions := repository.OptionRepositoryInit(db)
	msgg := appOptions.GetOption("whatsAppMsgEk1")
	return textWhatsApp(tel, msgg)
}

//for test
func SendMsgWp4() bool {
	db := repository.DB
	access := repository.OptionRepositoryInit(db)

	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "true" {
		return true
	} else {
		return false
	}
}

//for test
func textWhatsApp(tel string, msgg string) string {
	return "whatsapp://send/?phone=+90" + tel + "&text=" + url.QueryEscape(msgg) //v3
}

func telReplace(tel string) string {
	return strings.Replace(tel, " ", "", -1)
}

func msgReplace(msg string, slug string, url string) string {
	return strings.Replace(msg, "[link]", url+"kurbanBilgi/"+slug+"/", -1)
}

// func HasCacheIsOpen(db2 *gorm.DB) bool {

// 	// db := repository.DB
// 	access := repository.OptionRepositoryInit(db2)

// 	cacheControl := access.GetOption("cache_open_close")

// 	if cacheControl == "true" {
// 		return true
// 	} else {
// 		return false
// 	}
// }

//Example Not Used
// func HasCacheIsOpenSession() bool {
// 	var c *gin.Context
// 	cacheStore := stncsession.GetStore(c, "cache_open_close")
// 	cacheControl := reflect.ValueOf(cacheStore).String()
// 	if cacheControl == "true" {
// 		return true
// 	} else {
// 		return false
// 	}
// }
