// module provides utility methods for analyzing the user's module.
package module

import (
	"bytes"
	"errors"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	ErrModuleNotFound = errors.New("module not found")
)

// Find returns the module name and the absolute path of the module's root directory.
// `outPath` is the absolute path to the directory the packages will be generated into.
// It starts `outPath` and walks towards the user's home directory until it finds a go.mod file.
// If no go.mod file is found it returns ErrModuleNotFound.
func Find(outPath string) (string, string, error) {
	absOut, err := filepath.Abs(outPath)
	if err != nil {
		return "", "", err
	}

	modFilePath, err := modFilePath(absOut)
	if err != nil {
		return "", "", err
	}

	modDirPath := filepath.Dir(modFilePath)

	modName, err := modName(modDirPath)
	if err != nil {
		return "", "", err
	}

	return modName, modDirPath, nil
}

// modName takes the absolute path of a module's root directory
// and evaluates the module's name
func modName(path string) (string, error) {
	cmd := exec.Command("go", "mod", "why")
	cmd.Dir = path

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	writtenLines := strings.Split(string(output), "\n")

	if len(writtenLines) == 3 {
		return writtenLines[1], nil
	}

	return "", ErrModuleNotFound
}

// modFilePath returns the absolute path of the
// go.mod `path` relates to.
// `path` must be absolute.
// returns error if path can't be evaluated or
// `path` is not part of a module.
func modFilePath(path string) (string, error) {
	cmd := exec.Command("go", "env", "GOMOD")
	cmd.Dir = path

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	if bytes.Equal(output, []byte("/dev/null\n")) {
		return "", ErrModuleNotFound
	}

	filePath := string(bytes.TrimSpace(output))

	return filePath, nil
}
