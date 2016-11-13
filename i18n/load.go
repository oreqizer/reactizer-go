package i18n

import (
	"fmt"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/nicksnyder/go-i18n/i18n"
)

func LoadTranslations(locales []string) {
	for _, lang := range locales {
		path, _ := filepath.Abs(fmt.Sprintf("locales/%s.all.json", lang))
		err := i18n.LoadTranslationFile(path)
		if err != nil {
			glog.Error(err)
		}
	}
}
