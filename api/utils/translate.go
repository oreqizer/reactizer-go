package utils

import (
	"github.com/nicksnyder/go-i18n/i18n"
	"github.com/kataras/iris"
	"github.com/golang/glog"

	"reactizer-go/config"
)

// 'getT' returns i18n.TranslateFunc based on the request headers.
// Prority: 'X-Lang' > 'Accept-Language' > default language
//
// In case of an error, it returns i18n.IdentityTfunc().
func GetT(c *iris.Context) i18n.TranslateFunc {
	selectLang := c.RequestHeader("X-Lang")
	acceptLang := c.RequestHeader("Accept-Language")
	T, err := i18n.Tfunc(selectLang, acceptLang, config.DefaultLanguage)
	if err != nil {
		glog.Error(err)
		return i18n.IdentityTfunc()
	}
	return T
}
