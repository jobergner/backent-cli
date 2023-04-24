package module

import (
	"os/exec"
)

func Tidy() error {
	cmd := exec.Command("go", "mod", "tidy")

	if _, err := cmd.Output(); err != nil {
		return err
	}

	return nil
}
