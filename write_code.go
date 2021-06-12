package main

import (
	"github.com/Java-Jonas/bar-cli/enginefactory"
	"github.com/Java-Jonas/bar-cli/serverfactory"

	"bytes"
)

func writeCode(c *config) []byte {
	buf := bytes.NewBufferString("package state\n")

	writeCombinedImport(buf)
	writeImportedFiles(buf)

	enginefactory.WriteEngine(buf, c.State)
	serverfactory.WriteServer(buf, c.State, c.Actions)

	return buf.Bytes()
}
