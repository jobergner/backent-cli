package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func generateMarshallers() error {
	if ok := commandExists("easyjson"); !ok {
		return fmt.Errorf("easyjson is required!\n\ninstall with `go get -u github.com/mailru/easyjson/...`")
	}

	cmd := exec.Command("easyjson", "-all", "-byte", "-omit_empty", filepath.Join(*outDirname, outFile))
	// error is being swallowed as easyjson throws errors while actually functioning properly
	// all underlying requirements have already been checked with `validateOutDir` at this point
	// whether generating the marshallers was successfull will be validated with running `go build` later
	cmd.Run()

	return nil
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
