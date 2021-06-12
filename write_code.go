package main

import (
	"bar-cli/enginefactory"
	"bar-cli/serverfactory"

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
