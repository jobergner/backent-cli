//go:generat bash ./generate.sh
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jobergner/backent-cli/pkg/env"
	"github.com/jobergner/backent-cli/pkg/factory/utils"
	"github.com/jobergner/backent-cli/pkg/marshallers"
	"github.com/jobergner/backent-cli/pkg/module"
	"github.com/jobergner/backent-cli/pkg/packages"
)

var configPath = flag.String("config", "./example.config.json", "path of config")
var outDirPath = flag.String("out", "./tmp", "where to write the files to")

func main() {
	flag.Parse()

	config, errs := newAST()
	if len(errs) != 0 {

		for _, validationErr := range errs {
			fmt.Println(validationErr)
		}

		panic("\nthe above errors have occured while validating " + *configPath)
	}

	err := env.EnsureDir(*outDirPath)
	if err != nil {
		panic(err)
	}

	importPath, err := evalImportPath(*outDirPath)
	if err != nil {
		panic(err)
	}

	for _, pkg := range packages.Packages(config) {

		dirPath := filepath.Join(*outDirPath, pkg.Name)

		err := env.EnsureDir(dirPath)
		if err != nil {
			panic(err)
		}

		staticCodeTemplate := packages.StaticCode[pkg.StaticCodeIdentifier]
		staticCode := strings.ReplaceAll(staticCodeTemplate, "{{path}}", importPath)

		buf := bytes.NewBufferString(staticCode)

		if pkg.DynamicCodeFactory != nil {
			customCode := pkg.DynamicCodeFactory.Write()

			buf.WriteString(customCode)
		}

		filePath := filepath.Join(dirPath, fmt.Sprintf("%s.go", pkg.Name))

		err = utils.Format(buf)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(filePath, buf.Bytes(), 0644)
		if err != nil {
			panic(err)
		}

		err = marshallers.Generate(filePath)
		if err != nil {
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
