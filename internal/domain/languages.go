package domain

type SupportedLangages struct {
	langages map[string]Converter
}

var support = &SupportedLangages{langages: make(map[string]Converter)}

func GetSupportedLangages() *SupportedLangages {
	return support
}

func (sl *SupportedLangages) AddSupport(newLang string, converter Converter) bool {
	sl.langages[newLang] = converter
	return true
}
