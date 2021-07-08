//go:generate bash ./generate.sh
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Java-Jonas/bar-cli/getstartedfactory"
)

const outFile = "state.go"

var configNameFlag = flag.String("config", "./example.config.json", "path of config")
var engineOnlyFlag = flag.Bool("engine_only", false, "only state")
var outDirName = flag.String("out", "./tmp", "where to write the files to")
var exampleFlag = flag.Bool("example", false, "when enabled starts example")

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

	if err := ioutil.WriteFile(filepath.Join(*outDirName, outFile), writeCode(c, configJson), 0644); err != nil {
		panic(fmt.Errorf("error while writing generated code to file system: %s", err))
	}

	if err := generateMarshallers(); err != nil {
		panic(fmt.Errorf("error while generating marshallers: %s", err))
	}

	if err := validateBuild(); err != nil {
		panic(fmt.Errorf("something went wrong when generating the code: %s", err))
	}

	if outDirModuleName, err := getModuleName(); err != nil {
		panic(err)
	} else {
		fmt.Println(getstartedfactory.WriteGetStarted(outDirModuleName, c.State, c.Actions, c.Responses))
	}
}

func validateOutDir() error {
	fi, err := os.Stat(*outDirName)

	if err != nil {
		return fmt.Errorf("defined out target \"%s\" does not exist", *outDirName)
	}

	mode := fi.Mode()
	if !mode.IsDir() {
		return fmt.Errorf("defined out target \"%s\" is not a directory", *outDirName)
	}

	cmd := exec.Command("go", "env", "GOMOD")
	cmd.Dir = *outDirName

	stdout, err := cmd.Output()
	if len(stdout) == 1 && string(stdout[0]) == "\n" {
		return fmt.Errorf("defined out target \"%s\" is not within GOPATH which is required for generating marshallers\ntip: initialize a go module in directory or it's parent!", *outDirName)
	}

	return nil
}

func validateBuild() error {
	cmd := exec.Command("go", "test", ".")
	cmd.Dir = *outDirName

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func getModuleName() (string, error) {
	cmd := exec.Command("go", "mod", "why")
	cmd.Dir = *outDirName

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	writtenLines := strings.Split(string(out), "\n")
	if len(writtenLines) != 3 || len(writtenLines[1]) == 0 {
		return "", fmt.Errorf("could not read module name of out target")
	}

	fmt.Println(writtenLines[1])

	return writtenLines[1], nil
}
