//go:generat bash ./generate.sh
package main

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/jobergner/backent-cli/pkg/ast"
	"github.com/jobergner/backent-cli/pkg/config"
	"github.com/jobergner/backent-cli/pkg/env"
	"github.com/jobergner/backent-cli/pkg/factory/utils"
	"github.com/jobergner/backent-cli/pkg/marshallers"
	"github.com/jobergner/backent-cli/pkg/module"
	"github.com/jobergner/backent-cli/pkg/packages"

	_ "github.com/mailru/easyjson/gen"
	_ "github.com/mailru/easyjson/jlexer"
	_ "github.com/mailru/easyjson/jwriter"
)

var configPath = flag.String("config", "./example.config.json", "path of config")
var outDirPath = flag.String("out", "./tmp", "where to place packages")

func main() {
	flag.Parse()

	if err := env.EnsureDir(*outDirPath); err != nil {
		panic(err)
	}

	importPath, err := evalImportPath(*outDirPath)
	if err != nil {
		panic(err)
	}

	state, actions, responses, err := config.Read(*configPath)
	if err != nil {
		panic(err)
	}

	syntaxTree := ast.Parse(state, actions, responses)

	for _, pkg := range packages.Packages(syntaxTree) {

		dirPath, filePath := pkg.Paths(*outDirPath)

		if err := env.EnsureDir(dirPath); err != nil {
			panic(err)
		}

		buf := bytes.NewBuffer(nil)

		if staticCodeTemplate, ok := packages.StaticCode[pkg.StaticCodeIdentifier]; ok {
			staticCode := strings.ReplaceAll(staticCodeTemplate, "{{path}}", importPath)
			if _, err := buf.WriteString(staticCode); err != nil {
				panic(err)
			}
		}

		if pkg.DynamicCodeFactory != nil {
			customCode := pkg.DynamicCodeFactory.Write()
			if _, err := buf.WriteString(customCode); err != nil {
				panic(err)
			}
		}

		if pkg.Lang() == packages.LangGo {
			if err = utils.Format(buf); err != nil {
				panic(err)
			}
		}

		if err := os.WriteFile(filePath, buf.Bytes(), 0644); err != nil {
			panic(err)
		}
	}

	if err := marshallers.WriteImportFile(*outDirPath); err != nil {
		panic(err)
	}

	if err := module.Tidy(); err != nil {
		panic(err)
	}

	for _, pkg := range packages.Packages(syntaxTree) {
		if pkg.Lang() != packages.LangGo {
			continue
		}

		_, filePath := pkg.Paths(*outDirPath)
		if err = marshallers.Generate(filePath); err != nil {
			panic(err)
		}
	}
}

// evalImportPath evaluates the base path that can be used to import
// the individual packages: `{module_name}/path/to/out`
func evalImportPath(outPath string) (string, error) {
	absOut, err := filepath.Abs(outPath)
	if err != nil {
		return "", err
	}

	modName, modPath, err := module.Find(absOut)
	if err != nil {
		return "", err
	}

	modToOut, err := filepath.Rel(modPath, absOut)
	if err != nil {
		return "", err
	}

	importPath := filepath.Join(modName, modToOut)

	return importPath, nil
}
