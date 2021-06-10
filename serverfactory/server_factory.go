package serverfactory

import (
	. "bar-cli/factoryutils"
	"bytes"
	"go/format"
	"go/parser"
	"go/token"

	"bar-cli/ast"
)

type ServerFactory struct {
	config *ast.AST
	buf    *bytes.Buffer
}

// isIDTypeOfType evaluates whether a given type name is the respective ID-Type
// of a user-defined type.
// Background:
// Every user-defined type has a generated ID type.
// E.g. a defined type "person" has its ID-Type "PersonID" generated automatically
func (s ServerFactory) isIDTypeOfType(typeName string) bool {
	for _, configType := range s.config.Types {
		if configType.Name+"ID" == typeName {
			return true
		}
	}
	return false
}

func newServerFactory(config *ast.AST) *ServerFactory {
	return &ServerFactory{
		config: config,
		buf:    &bytes.Buffer{},
	}
}

// WriteServerFrom writes source code for a given ActionsConfig
func WriteServer(buf *bytes.Buffer, stateConfigData, actionsConfigData map[interface{}]interface{}) {
	config := ast.Parse(stateConfigData, actionsConfigData)
	s := newServerFactory(config).
		writePackageName(). // to be able to format the code without errors
		writeMessageKinds().
		writeActions().
		writeParameters().
		writeProcessClientMessage().
		writeStart()

	err := s.format()
	if err != nil {
		// unexpected error
		panic(err)
	}

	buf.WriteString(TrimPackageName(s.buf.String()))
}

func (s *ServerFactory) writePackageName() *ServerFactory {
	s.buf.WriteString("package main\n")
	return s
}

func (s *ServerFactory) format() error {
	config, err := parser.ParseFile(token.NewFileSet(), "", s.buf.String(), parser.AllErrors)
	if err != nil {
		return err
	}

	s.buf.Reset()
	err = format.Node(s.buf, token.NewFileSet(), config)
	return err
}
