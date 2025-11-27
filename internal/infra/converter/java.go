package converter

import (
	"code-gen/internal/domain/graphql"
	logger "code-gen/pkg/log"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

var log = logger.GetLogger("JavaConverter")

type JavaConverter struct{}

type JavaFile struct {
	Name    string
	Fields  []graphql.FieldInfo
	Package string
}

func NewJavaConverter() *JavaConverter {
	return &JavaConverter{}
}

func (jc *JavaConverter) Convert(pathFile string, outputPath string) error {
	log.Info("Start convertion")
	var files []string
	info, err := os.Stat(pathFile)
	if err != nil {
		fmt.Println("erro:", err)
		os.Exit(1)
	}
	if info.IsDir() {
		_ = filepath.Walk(pathFile, func(p string, fi os.FileInfo, err error) error {
			if err == nil && !fi.IsDir() && strings.HasSuffix(p, ".graphqls") {
				files = append(files, p)
			}
			return nil
		})
	} else {
		files = append(files, pathFile)
	}
	var combined string
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			fmt.Println("erro lendo", f, err)
			os.Exit(1)
		}
		combined += "\n" + string(b)
	}

	doc, err := gqlparser.LoadSchema(&ast.Source{Input: combined})
	if err != nil {
		fmt.Println("erro parseando schema:", err)
		os.Exit(1)
	}
	tmpl := template.Must(template.New("java").Parse(javaTpl))

	_ = os.MkdirAll(outputPath, 0755)
	for name, def := range doc.Types {
		if strings.HasPrefix(name, "__") {
			continue
		}
		if def.Kind != ast.Object {
			continue
		}
		if name == "Query" || name == "Mutation" || name == "Subscription" {
			continue
		}

		ti := JavaFile{
			Name:    name,
			Fields:  []graphql.FieldInfo{},
			Package: definePackageNameBasedInOutputDir(outputPath),
		}
		for _, f := range def.Fields {
			fi := graphql.FieldInfo{
				FieldName: toCamel(f.Name),
				Type:      gqlTypeToJava(f.Type),
			}
			ti.Fields = append(ti.Fields, fi)
		}

		outFile := filepath.Join(outputPath, ti.Name+".java")
		out, err := os.Create(outFile)
		if err != nil {
			fmt.Println("erro criar arquivo", outFile, err)
			continue
		}
		if err := tmpl.Execute(out, ti); err != nil {
			fmt.Println("erro gerar template para", name, err)
		}
		out.Close()
		log.Info(fmt.Sprintf("gerado: %s", outFile))
	}
	return nil
}

var javaTpl = `package {{.Package}};

public class {{.Name}} {
{{- range .Fields}}
    private {{.Type}} {{.FieldName}};
{{- end}}

{{- range .Fields}}

    public {{.Type}} get{{.FieldName}}() {
        return this.{{.FieldName}};
    }
    public void set{{.FieldName}}({{.Type}} {{.FieldName}}) {
        this.{{.FieldName}} = {{.FieldName}};
    }
{{- end}}
}
`

func toCamel(s string) string {
	if s == "" {
		return s
	}
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.Title(s)
	s = strings.ReplaceAll(s, " ", "")
	return strings.ToLower(s[:1]) + s[1:]
}

func gqlTypeToJava(t *ast.Type) string {
	base := t.NamedType
	if base == "" && t.Elem != nil {
		return "java.util.List<" + gqlTypeToJava(t.Elem) + ">"
	}
	switch base {
	case "ID", "String":
		return "String"
	case "Int":
		return "Integer"
	case "Float":
		return "Double"
	case "Boolean":
		return "Boolean"
	default:
		return base
	}
}

func definePackageNameBasedInOutputDir(s string) string {
	re := regexp.MustCompile(`[^A-Za-z]+`)
	parts := re.Split(s, -1)

	out := parts[:0]
	for _, p := range parts {
		if p != "" {
			out = append(out, p)
		}
	}

	return strings.Join(out, ".")
}
