package i18n

import "github.com/nicksnyder/go-i18n/i18n"

func LoadTranslations() {
	i18n.MustLoadTranslationFile("/Users/oreqizer/go/src/reactizer-go/i18n/locales/en-US.all.json")
}
