package modules

import (
	"net/http"

	"github.com/nicksnyder/go-i18n/i18n"

	"reactizer-go/config"
)

// 'getT' returns i18n.TranslateFunc based on the request headers.
// Prority: 'X-Lang' > 'Accept-Language' > default language
//
// In case of an error, it returns i18n.IdentityTfunc().
func getT(r *http.Request) i18n.TranslateFunc {
	selectLang := r.Header.Get("X-Lang")
	acceptLang := r.Header.Get("Accept-Language")
	T, err := i18n.Tfunc(selectLang, acceptLang, config.DefaultLanguage)
	if err != nil {
		return i18n.IdentityTfunc()
	}
	return T
}
