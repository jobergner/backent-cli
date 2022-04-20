package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	fi, err := os.Stat(*outDirPath)

	if err != nil {
		panic(err)
	}

	mode := fi.Mode()
	if !mode.IsDir() {
		panic(fmt.Sprintf("defined out target \"%s\" is not a directory", *outDirPath))
	}

	cmd := exec.Command("go", "env", "GOMOD")
	cmd.Dir = *outDirPath

	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	if string(stdout) == "/dev/null\n" {
		panic(fmt.Sprintf("defined out target \"%s\" is not within GOPATH which is required for generating marshallers\ntip: initialize a go module in directory or it's parent!", *outDirPath))
	}
}

func tidyModules() {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = *outDirPath

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// getModuleName returns the module name and the abolute path of the directory file
// where the go.mod file is located
func getModuleName() (string, string) {

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	p, err := filepath.Abs(filepath.Join(*outDirPath))
	if err != nil {
		panic(err)
	}

	for {
		cmd := exec.Command("go", "mod", "why")
		cmd.Dir = p

		out, err := cmd.Output()
		if err != nil {
			panic(err)
		}

		writtenLines := strings.Split(string(out), "\n")

		if len(writtenLines) == 3 {
			abs, err := filepath.Abs(p)
			if err != nil {
				panic(err)
			}
			return writtenLines[1], abs
		}

		p, err = filepath.Abs(filepath.Join(p, "../"))
		if err != nil {
			panic(err)
		}

		if userHomeDir == p {
			panic("could not find module name")
		}
	}
}
