package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func generateMarshallers() error {
	cmd := exec.Command("easyjson", "-all", "-omit_empty", filepath.Join(outDir, outFile))
	if out, err := cmd.Output(); err != nil {
		return err
	} else {
		fmt.Println(string(out))
	}
	return nil
}
