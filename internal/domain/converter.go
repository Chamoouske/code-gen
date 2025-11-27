package domain

type Converter interface {
	Convert(pathFile string, outputPath string) error
}
