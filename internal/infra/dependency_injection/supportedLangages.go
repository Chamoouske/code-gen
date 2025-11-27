package dependencyinjection

import "code-gen/internal/domain"

var sl domain.SupportedLangages

func init() {
	sl = domain.GetNewSupportedLangages()
	addConverters(sl)
}

func GetSupportedLangages() *domain.SupportedLangages {
	return &sl
}

func AddSupport(lang string, converter domain.Converter) {
	sl.AddSupport(lang, converter)
}
