package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func generateMarshallers(firstAttempt bool) error {
	if ok := commandExists("easyjson"); !ok {
		cmd := exec.Command("go", "get", "-u", "github.com/mailru/easyjson/...")
		out, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("error installing mailru/easyjson: %s \n %s", string(out), err)
		}
	}

	cmd := exec.Command("easyjson", "-all", "-byte", "-omit_empty", filepath.Join(*outDirName, outFile))
	// error is being printed as a warning as easyjson throws errors while actually functioning properly
	// all underlying requirements have already been checked with `validateOutDir` at this point
	output, err := cmd.CombinedOutput()
	if err != nil {
		if firstAttempt {

			if err := tidyModules(); err != nil {
				panic(fmt.Errorf("something went wrong while tidying modules: %s", err))
			}

			if err := generateMarshallers(false); err != nil {
				return err
			}

		} else {
			return fmt.Errorf("generating marshallers caused issues:\n %s %s", err, string(output))
		}
	}

	return nil
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
