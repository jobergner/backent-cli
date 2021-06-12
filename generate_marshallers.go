package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func generateMarshallers() error {
	if ok := commandExists("easyjson"); !ok {
		panic("easyjson is required!\n\ninstall with `go get -u github.com/mailru/easyjson/...`")
	}
	cmd := exec.Command("easyjson", "-all", "-omit_empty", filepath.Join(*outDirname, outFile))
	if out, err := cmd.Output(); err != nil {
		fmt.Printf("error generating marshallers - is the output directory `%s` in GOPATH?\n", *outDirname)
		return err
	} else {
		fmt.Println(string(out))
	}
	return nil
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
