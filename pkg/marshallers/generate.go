package marshallers

import (
	"fmt"
	"os/exec"
)

func Generate(filePath string) error {
	cmd := exec.Command("easyjson", "-all", "-byte", "-omit_empty", filePath)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s : %s", err.Error(), string(output))
	}

	return nil
}
