package serverfactory

import (
	"bytes"

	. "github.com/Java-Jonas/bar-cli/factoryutils"

	"github.com/Java-Jonas/bar-cli/ast"
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
func WriteServer(
	buf *bytes.Buffer,
	stateConfigData, actionsConfigData, responsesConfigData map[interface{}]interface{},
) {
	config := ast.Parse(stateConfigData, actionsConfigData, responsesConfigData)
	s := newServerFactory(config).
		writePackageName(). // to be able to format the code without errors
		writeMessageKinds().
		writeActions().
		writeParameters().
		writeResponses().
		writeProcessClientMessage().
		writeStart()

	err := Format(s.buf)
	if err != nil {
		// unexpected error
		panic(err)
	}

	buf.WriteString(TrimPackageName(s.buf.String()))
}

func (s *ServerFactory) writePackageName() *ServerFactory {
	s.buf.WriteString("package state\n")
	return s
}
