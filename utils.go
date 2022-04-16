package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ensureDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func validateOutDir() {
	fi, err := os.Stat(*outDirName)

	if err != nil {
		panic(err)
	}

	mode := fi.Mode()
	if !mode.IsDir() {
		panic(fmt.Sprintf("defined out target \"%s\" is not a directory", *outDirName))
	}

	cmd := exec.Command("go", "env", "GOMOD")
	cmd.Dir = *outDirName

	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	if string(stdout) == "/dev/null\n" {
		panic(fmt.Sprintf("defined out target \"%s\" is not within GOPATH which is required for generating marshallers\ntip: initialize a go module in directory or it's parent!", *outDirName))
	}
}

func tidyModules() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = *outDirName

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func getModuleName() string {
	cmd := exec.Command("go", "mod", "why")
	cmd.Dir = *outDirName

	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	writtenLines := strings.Split(string(out), "\n")
	if len(writtenLines) != 3 || len(writtenLines[1]) == 0 {
		panic("could not read module name of out target")
	}

	return writtenLines[1]
}
