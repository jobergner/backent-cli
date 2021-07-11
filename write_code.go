package main

import (
	"github.com/jobergner/backent-cli/enginefactory"
	"github.com/jobergner/backent-cli/serverfactory"

	"bytes"
)

func writeCode(c *config, configJson []byte) []byte {
	buf := bytes.NewBufferString("package state\n")

	if *engineOnlyFlag {
		buf.WriteString("\n" + engine_only_import_decl)
	} else {
		buf.WriteString("\n" + import_decl)
		buf.WriteString("\n" + imported_server_example_files)
	}

	enginefactory.WriteEngine(buf, c.State)
	if !*engineOnlyFlag {
		serverfactory.WriteServer(buf, c.State, c.Actions, c.Responses, configJson)
	}

	return buf.Bytes()
}
