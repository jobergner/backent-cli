package serverfactory

import (
	"bytes"
	"strings"

	"bar-cli/ast"
)

type ServerFactory struct {
	config *ast.AST
	buf    *bytes.Buffer
}

func newServerFactory(config *ast.AST) *ServerFactory {
	return &ServerFactory{
		config: config,
		buf:    &bytes.Buffer{},
	}
}

func title(name string) string {
	return strings.Title(name)
}
