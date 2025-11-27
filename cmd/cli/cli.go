package main

import (
	dependencyinjection "code-gen/internal/infra/dependency_injection"
	logger "code-gen/pkg/log"
	"flag"
	"fmt"
	"os"
)

func main() {
	var log = logger.GetLogger("CLI MAin")
	schemaPath := flag.String("f", "", "arquivo .graphqls ou pasta com .graphqls")
	outDir := flag.String("o", "generated", "diretório de saída")
	lang := flag.String("l", "java", "linguagem alvo")
	flag.Parse()

	fmt.Printf("%s", *lang)
	var supported = dependencyinjection.GetSupportedLangages()

	var converter, err = supported.GetConverterForLangage(*lang)

	if err != nil {
		log.Info(fmt.Sprintf("Error: %s", err.Error()))
		os.Exit(1)
	}

	converter.Convert(*schemaPath, *outDir)
}
