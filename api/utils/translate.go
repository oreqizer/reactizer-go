package utils

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	// Check https://github.com/iris-contrib/middleware/tree/master/go-i18n too.
	"github.com/nicksnyder/go-i18n/i18n"
	"reactizer-go/config"
)

// 'getT' returns i18n.TranslateFunc based on the request headers.
// Prority: 'X-Lang' > 'Accept-Language' > default language
//
// In case of an error, it returns i18n.IdentityTfunc().
func GetT(c iris.Context) i18n.TranslateFunc {
	selectLang := c.GetHeader("X-Lang")
	acceptLang := c.GetHeader("Accept-Language")
	T, err := i18n.Tfunc(selectLang, acceptLang, config.DefaultLanguage)
	if err != nil {
		golog.Error(err)
		return i18n.IdentityTfunc()
	}
	return T
}
