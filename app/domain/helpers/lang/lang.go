package lang

import "github.com/leonelquinteros/gotext"

func LoadLanguages(menu string) (defaultMenu *gotext.Locale, menus *gotext.Locale) {
	//var defaultMenu, menus *gotext.Locale
	defaultMenu = gotext.NewLocale("public/locales", "en_US")
	defaultMenu.AddDomain(menu)

	menus = gotext.NewLocale("public/locales", "en_US")
	menus.AddDomain("menus")
	return defaultMenu, menus
}
