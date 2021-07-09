package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Java-Jonas/bar-cli/getstartedfactory"
)

func generate() {

	config, configJson, err := readConfig()
	if err != nil {
		panic(err)
	}

	validationErrs := validateConfig(config)
	if len(validationErrs) != 0 {
		for _, validationErr := range validationErrs {
			fmt.Println(validationErr)
		}
		panic("\nthe above errors have occured while validating " + *configNameFlag)
	}

	if err := ensureOutDir(); err != nil {
		panic(err)
	}

	if err := validateOutDir(); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(filepath.Join(*outDirName, outFile), writeCode(config, configJson), 0644); err != nil {
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
	} else if !*exampleFlag {
		fmt.Println(getstartedfactory.WriteGetStarted(outDirModuleName, false, config.State, config.Actions, config.Responses))
	} else {
		cleanOutDirPath := filepath.Clean(*outDirName)
		outDirPathBase := filepath.Base(cleanOutDirPath)
		mainFilePath := cleanOutDirPath[0:len(cleanOutDirPath)-len(outDirPathBase)] + "/main.go"
		mainFileContent := getstartedfactory.WriteGetStarted(outDirModuleName, true, config.State, config.Actions, config.Responses)
		if err := ioutil.WriteFile(mainFilePath, []byte(mainFileContent), 0644); err != nil {
			panic(fmt.Errorf("error while writing generated code to file system: %s", err))
		}
	}
}

func ensureOutDir() error {
	if _, err := os.Stat(*outDirName); os.IsNotExist(err) {
		err := os.Mkdir(*outDirName, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory for generated code: %s", err)
		}
	}
	return nil
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

	return writtenLines[1], nil
}
