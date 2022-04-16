//go:generat bash ./generate.sh
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/jobergner/backent-cli/pkg/packages"
)

var configNameFlag = flag.String("config", "./example.config.json", "path of config")
var outDirName = flag.String("out", "./tmp", "where to write the files to")
var exampleFlag = flag.Bool("example", false, "when enabled starts example")
var devModeFlag = flag.Bool("dev", false, "start in dev mode")
var portFlag = flag.String("port", "3100", "start in dev mode")

func main() {
	flag.Parse()

	config, errs := newAST()
	if len(errs) != 0 {

		for _, validationErr := range errs {
			fmt.Println(validationErr)
		}

		panic("\nthe above errors have occured while validating " + *configNameFlag)
	}

	ensureDir(*outDirName)

	validateOutDir()

	modName := getModuleName()

	for _, pkg := range packages.Packages(config) {
		writePackage(pkg, modName)
	}
}

func writePackage(pkg packages.PackageInfo, modName string) {
	dirPath := path.Join(*outDirName, pkg.Name)

	ensureDir(dirPath)

	buf := bytes.NewBuffer(nil)

	if pkg.DynamicCodeFactory != nil {
		pkg.DynamicCodeFactory.Write(buf)
	} else {
		buf.WriteString(fmt.Sprintf("package %s\n", pkg.Name))
	}

	moduleName := getModuleName()
	code := strings.ReplaceAll(staticCode[pkg.StaticCodeIdentifier], "{{path}}", moduleName)
	buf.WriteString(code)

	filePath := path.Join(dirPath, fmt.Sprintf("%s.go", pkg.Name))

	err := os.WriteFile(filePath, buf.Bytes(), 0644)
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

			if err := tidyModules(); err != nil {
				panic(fmt.Errorf("something went wrong while tidying modules: %s", err))
			}

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
