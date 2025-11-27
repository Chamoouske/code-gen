package usecase

import (
	"code-gen/internal/domain"
	"fmt"
	"strings"
)

type ConverterFile struct {
	supportedLangs domain.SupportedLangages
}

func (cf *ConverterFile) Converter(inputPath string, langage string, outputPath string) error {
	converter, err := cf.supportedLangs.GetConverterForLangage(strings.ToLower(langage))

	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	converter.Convert(inputPath, outputPath)

	return nil
}
