package main

import (
	"bar-cli/enginefactory"
	"bar-cli/examples"
	"bar-cli/serverfactory"
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

const outDir = "./tmp"
const outFile = "state.go"

func main() {
	buf := bytes.NewBufferString("package state\n")
	writeCombinedImport(buf)
	writeImportedFiles(buf)
	enginefactory.WriteEngine(buf, examples.StateConfig)
	serverfactory.WriteServer(buf, examples.StateConfig, examples.ActionsConfig)
	if err := ioutil.WriteFile(filepath.Join(outDir, outFile), buf.Bytes(), 0644); err != nil {
		panic(err)
	}
	cmd := exec.Command("easyjson", "-all", "-omit_empty", filepath.Join(outDir, outFile))
	if out, err := cmd.Output(); err != nil {
		panic(err)
	} else {
		fmt.Println(string(out))
	}
}
