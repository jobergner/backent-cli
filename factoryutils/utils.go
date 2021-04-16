package factoryutils

import (
	"bar-cli/ast"
	"bytes"
	"strings"

	"github.com/dave/jennifer/jen"
)

func ForEachParamInAction(action ast.Action, fn func(param ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	action.RangeParams(func(field ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachFieldInType(configType ast.ConfigType, fn func(field ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	configType.RangeFields(func(field ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachTypeInAST(config *ast.AST, fn func(configType ast.ConfigType) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	config.RangeTypes(func(configType ast.ConfigType) {
		statements = append(statements, fn(configType))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachActionInAST(config *ast.AST, fn func(action ast.Action) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	config.RangeActions(func(action ast.Action) {
		statements = append(statements, fn(action))
		statements = append(statements, jen.Line())
	})
	return &statements
}

type DeclSet struct {
	File *jen.File
}

func NewDeclSet() DeclSet {
	return DeclSet{
		File: jen.NewFile("main"),
	}
}

func (d DeclSet) Render(buf *bytes.Buffer) {
	var _buf bytes.Buffer
	err := d.File.Render(&_buf)
	if err != nil {
		panic(err)
	}
	code := strings.TrimPrefix(_buf.String(), "package main")
	code = strings.TrimSpace(code)
	buf.WriteString("\n" + code + "\n")
}
