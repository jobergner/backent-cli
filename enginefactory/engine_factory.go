package enginefactory

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"strings"

	"bar-cli/ast"

	"github.com/dave/jennifer/jen"
	"github.com/gertd/go-pluralize"
)

// TODO wtf
const isProductionMode = false

func title(name string) string {
	return strings.Title(name)
}

func forEachFieldInType(configType ast.ConfigType, fn func(field ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	configType.RangeFields(func(field ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func forEachTypeInAST(config *ast.AST, fn func(configType ast.ConfigType) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	config.RangeTypes(func(configType ast.ConfigType) {
		statements = append(statements, fn(configType))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func onlyIf(is bool, statement *jen.Statement) *jen.Statement {
	if is {
		return statement
	}
	return jen.Empty()
}

type declSet struct {
	file *jen.File
}

func newDeclSet() declSet {
	return declSet{
		file: jen.NewFile("main"),
	}
}

func (d declSet) render(buf *bytes.Buffer) {
	var _buf bytes.Buffer
	err := d.file.Render(&_buf)
	if err != nil {
		panic(err)
	}
	code := strings.TrimPrefix(_buf.String(), "package main")
	code = strings.TrimSpace(code)
	buf.WriteString("\n" + code + "\n")
}

// pluralizeClient is used to find the singular of field names
// this is necessary for writing coherent method names, eg. in write_adders.go (toSingular)
// with getting the singular form of a plural, this field:
// { pieces []piece }
// can have the coherent adder method of "AddPiece"
var pluralizeClient *pluralize.Client = pluralize.NewClient()

type engineFactory struct {
	config *ast.AST
	buf    *bytes.Buffer
}

// WriteEngineFrom writes source code for a given State-/ActionsConfig
func WriteEngineFrom(stateConfigData map[interface{}]interface{}) []byte {
	config := ast.Parse(stateConfigData, map[interface{}]interface{}{})
	s := newStateFactory(config).
		writePackageName().
		writeOperationKind().
		writeIDs().
		writeState().
		writeEngine().
		writeGenerateID().
		writeUpdateState().
		writeElements().
		writeAdders().
		writeCreators().
		writeDeleters().
		writeGetters().
		writeRemovers().
		writeSetters().
		writeTree().
		writeTreeElements().
		writeAssembleTree().
		writeAssembleTreeElement().
		writeDeduplicate()

	err := s.format()
	if err != nil {
		// unexpected error
		panic(err)
	}

	return s.writtenSourceCode()
}

func (s *engineFactory) writePackageName() *engineFactory {
	s.buf.WriteString("package state\n")
	return s
}

func newStateFactory(config *ast.AST) *engineFactory {
	return &engineFactory{
		config: config,
		buf:    &bytes.Buffer{},
	}
}

func (s *engineFactory) writtenSourceCode() []byte {
	return s.buf.Bytes()
}

func (s *engineFactory) format() error {
	config, err := parser.ParseFile(token.NewFileSet(), "", s.buf.String(), parser.AllErrors)
	if err != nil {
		return err
	}

	s.buf.Reset()
	err = format.Node(s.buf, token.NewFileSet(), config)
	return err
}
