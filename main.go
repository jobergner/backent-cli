//go:generat bash ./generate.sh
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jobergner/backent-cli/pkg/factory/utils"
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

	ensureDir(*outDirPath)

	validateOutDir()

	absOutDirPath, err := filepath.Abs(*outDirPath)

	modName, modDirPath := getModuleName()
	pathToLibrary, err := filepath.Rel(modDirPath, absOutDirPath)
	if err != nil {
		panic(err)
	}

	libPath := filepath.Join(modName, pathToLibrary)

	for _, pkg := range packages.Packages(config) {
		writePackage(pkg, libPath)
	}
}

func writePackage(pkg packages.PackageInfo, libPath string) {
	dirPath := filepath.Join(*outDirPath, pkg.Name)

	ensureDir(dirPath)

	buf := bytes.NewBuffer(nil)

	code := strings.ReplaceAll(staticCode[pkg.StaticCodeIdentifier], "{{path}}", libPath)
	buf.WriteString(code)

	if pkg.DynamicCodeFactory != nil {
		buf.WriteString(pkg.DynamicCodeFactory.Write())
	}

	filePath := filepath.Join(dirPath, fmt.Sprintf("%s.go", pkg.Name))

	err := utils.Format(buf)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filePath, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	generateMarshallers(filePath, true)
}

func generateMarshallers(fileName string, firstAttempt bool) {
	if ok := commandExists("easyjson"); !ok {

		cmd := exec.Command("go", "install", "github.com/mailru/easyjson/...")

		out, err := cmd.Output()
		if err != nil {
			panic(fmt.Sprintf("error installing mailru/easyjson: %s \n %s", string(out), err))
		}
	}

	cmd := exec.Command("easyjson", "-all", "-byte", "-omit_empty", fileName)
	// error is being printed as a warning as easyjson throws errors while actually functioning properly
	// all underlying requirements have already been checked with `validateOutDir` at this point
	output, err := cmd.CombinedOutput()
	if err != nil {
		if firstAttempt {

			tidyModules()

			generateMarshallers(fileName, false)

		} else {
			panic(fmt.Sprintf("generating marshallers caused issues:\n %s %s", err, string(output)))
		}
	}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
