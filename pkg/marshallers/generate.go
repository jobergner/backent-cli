package marshallers

import (
	"errors"
	"fmt"
	"os/exec"
)

var (
	ErrEasyjsonNotFound = errors.New("install https://github.com/mailru/easyjson")
)

func Generate(filePath string) error {
	if ok := commandExists("easyjson"); !ok {
		return ErrEasyjsonNotFound
	}

	cmd := exec.Command("easyjson", "-all", "-byte", "-omit_empty", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s : %s", err.Error(), string(output))
	}

	return nil
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
