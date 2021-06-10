package main

import (
	"bar-cli/enginefactory"
	"bar-cli/examples"
	"bar-cli/serverfactory"
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBufferString("package state\n")
	writeCombinedImport(buf)
	writeImportedFiles(buf)
	enginefactory.WriteEngine(buf, examples.StateConfig)
	serverfactory.WriteServer(buf, examples.StateConfig, examples.ActionsConfig)
	fmt.Println(buf.String())
}
