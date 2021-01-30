package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type inputFile struct {
	name  string
	path  string
	decls []ast.Decl
}

func newInputFile(name, path string, content []byte) inputFile {
	f, err := parser.ParseFile(token.NewFileSet(), "", content, 0)
	check(err)

	file := inputFile{
		name:  name,
		path:  path,
		decls: f.Decls,
	}

	return file
}
