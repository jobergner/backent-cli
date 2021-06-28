//go:generate bash ./generate.sh
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Java-Jonas/bar-cli/getstartedfactory"
)

const outFile = "state.go"

var configNameFlag = flag.String("config", "./example.config.json", "path of config")
var engineOnlyFlag = flag.Bool("engine_only", false, "only state")
var outDirname = flag.String("out", "./tmp", "where to write the files to")

func main() {
	flag.Parse()

	c, configJson, err := readConfig()
	if err != nil {
		panic(err)
	}

	validationErrs := validateConfig(c)
	if len(validationErrs) != 0 {
		for _, validationErr := range validationErrs {
			fmt.Println(validationErr)
		}
		panic("\nthe above errors have occured while validating " + *configNameFlag)
	}

	if err := validateOutDir(); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(filepath.Join(*outDirname, outFile), writeCode(c, configJson), 0644); err != nil {
		panic(fmt.Errorf("error while writing generated code to file system: %s", err))
	}

	if err := generateMarshallers(); err != nil {
		panic(fmt.Errorf("error while generating marshallers: %s", err))
	}

	if err := validateBuild(); err != nil {
		panic("something went wrong. Please create an issue containing your environment details and config! You may also want to run `go build` in the out directory an include any errors.")
	}

	fmt.Println(getstartedfactory.WriteGetStarted(c.State, c.Actions, c.Responses))
}

func validateOutDir() error {
	fi, err := os.Stat(*outDirname)

	if err != nil {
		return fmt.Errorf("defined out target \"%s\" does not exist", *outDirname)
	}

	mode := fi.Mode()
	if !mode.IsDir() {
		return fmt.Errorf("defined out target \"%s\" is not a directory", *outDirname)
	}

	cmd := exec.Command("go", "env", "GOMOD")
	cmd.Dir = *outDirname

	stdout, err := cmd.Output()
	if len(stdout) == 1 && string(stdout[0]) == "\n" {
		return fmt.Errorf("defined out target \"%s\" is not within GOPATH which is required for generating marshallers\ntip: initialize a go module in directory or its parent!", *outDirname)
	}

	return nil
}

func validateBuild() error {
	cmd := exec.Command("go", "test", ".")
	cmd.Dir = *outDirname

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
