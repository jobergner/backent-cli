package build

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	ErrModuleNotFound = errors.New("module not found")
)

// ImportPath evaluates the base path that can be used to import
// the individual packages: `{module_name}/path/to/out`
func ImportPath(outPath string) (string, error) {
	absOut, err := filepath.Abs(outPath)

	modName, modPath, err := findMod(absOut)
	if err != nil {
		return "", err
	}

	modToOut, err := filepath.Rel(modPath, outPath)
	if err != nil {
		return "", err
	}

	importPath := filepath.Join(modName, modToOut)

	return importPath, nil
}

// findMod returns the module name and the absolute root directory path of the module.
// `outPath` is the absolute path to the directory the packages will be generated into.
// It starts `outPath` and walks towards the user's home directory until it finds a go.mod file.
// If no go.mod file is found it returns ErrModuleNotFound.
func findMod(outPath string) (string, string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}

	current := outPath

	for {
		modName, ok, err := hasModFile(current)
		if err != nil {
			return "", "", err
		}

		if ok {
			return modName, current, nil
		}

		current, err = filepath.Abs(filepath.Join(current, "../"))
		if err != nil {
			return "", "", err
		}

		if userHomeDir == current {
			return "", "", ErrModuleNotFound
		}
	}
}

// hasModFile evals whether the `path` directory contains a go.mod file.
// `path` must be absolute.
func hasModFile(path string) (string, bool, error) {
	cmd := exec.Command("go", "mod", "why")
	cmd.Dir = path

	out, err := cmd.Output()
	if err != nil {
		return "", false, err
	}

	writtenLines := strings.Split(string(out), "\n")

	if len(writtenLines) == 3 {
		return writtenLines[1], true, nil
	}

	return "", false, nil
}
