package dependencyinjection

import "code-gen/internal/domain"

var sl domain.SupportedLangages

func init() {
	sl = domain.GetNewSupportedLangages()
}

func GetSupportedLangages() *domain.SupportedLangages {
	return &sl
}
