package main

import (
	dependencyinjection "code-gen/internal/infra/dependency_injection"
	"flag"
	"fmt"
	"os"
)

func main() {
	schemaPath := flag.String("f", "", "arquivo .graphqls ou pasta com .graphqls")
	outDir := flag.String("o", "generated", "diretório de saída")
	lang := flag.String("l", "java", "linguagem alvo")
	flag.Parse()

	fmt.Printf("%s", *lang)
	var supported = dependencyinjection.GetSupportedLangages()

	var converter, err = supported.GetConverterForLangage(*lang)

	if err == nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	converter.Convert(*schemaPath, *outDir)
}
