package domain

import (
	logger "code-gen/pkg/log"
	"fmt"
	"strings"
)

var log = logger.GetLogger("Languages")

type SupportedLangages struct {
	langages map[string]Converter
}

func GetNewSupportedLangages() SupportedLangages {
	return SupportedLangages{langages: make(map[string]Converter)}
}

func (sl *SupportedLangages) AddSupport(newLang string, converter Converter) bool {
	log.Info(fmt.Sprintf("Add Converter: %s", newLang))
	sl.langages[strings.ToLower(newLang)] = converter
	return true
}

func (sl *SupportedLangages) GetConverterForLangage(lang string) (Converter, error) {
	var converter = sl.langages[lang]

	if converter == nil {
		return nil, fmt.Errorf("Converter must not be exists")
	}

	return converter, nil
}
