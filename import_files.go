package main

import (
	"bar-cli/factoryutils"
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"path/filepath"
)

var importedDirs = [2]string{
	"./serverfactory/server_example/server",
	"./enginefactory/state_engine_example",
}

var excludedFiles = []string{
	"serverfactory/server_example/server/gets_generated.go",
	"serverfactory/server_example/server/message_easyjson.go",
	"serverfactory/server_example/server/state.go",
	"enginefactory/state_engine_example/state_engine_test.go",
	"enginefactory/state_engine_example/state_engine_bench_test.go",
	"enginefactory/state_engine_example/tree_easyjson.go",
}

func scanDeclsInDir(directoryPath string) ([]ast.Decl, error) {

	fileInfos, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}

	var decls []ast.Decl

IterateFiles:
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		fileName := fileInfo.Name()
		filePath := filepath.Join(directoryPath, fileName)

		if filepath.Ext(fileName) != ".go" {
			continue
		}

		for _, excludedFile := range excludedFiles {
			if filepath.Join(directoryPath, fileName) == excludedFile {
				continue IterateFiles
			}
		}

		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		f, err := parser.ParseFile(token.NewFileSet(), "", content, 0)
		if err != nil {
			return nil, err
		}

		decls = append(decls, f.Decls...)
	}

	return decls, nil
}

func writeImportedFiles(buf *bytes.Buffer) {

	dirDecls, err := scanDeclsInDir(importedDirs[0])
	if err != nil {
		panic(err)
	}

	var decls []ast.Decl
	for _, decl := range dirDecls {
		if _, ok := isImportDecl(decl); !ok {
			decls = append(decls, decl)
		}
	}

	printer.Fprint(buf, token.NewFileSet(), decls)
}

func writeCombinedImport(buf *bytes.Buffer) {
	importDecl := &ast.GenDecl{
		Tok: token.IMPORT,
	}

	var decls []ast.Decl
	for _, importedDir := range importedDirs {
		dirDecls, err := scanDeclsInDir(importedDir)
		if err != nil {
			panic(err)
		}
		decls = append(decls, dirDecls...)
	}

	for _, decl := range decls {
		if genDecl, ok := isImportDecl(decl); ok {
			importDecl.Specs = append(importDecl.Specs, genDecl.Specs...)
		}
	}

	importBuf := bytes.NewBufferString("package state\n")
	printer.Fprint(importBuf, token.NewFileSet(), importDecl)
	factoryutils.Format(importBuf)
	buf.WriteString(factoryutils.TrimPackageName(importBuf.String()))
}

func isImportDecl(decl ast.Decl) (*ast.GenDecl, bool) {
	if genDecl, ok := decl.(*ast.GenDecl); ok {
		if genDecl.Tok == token.IMPORT {
			return genDecl, true
		}
	}
	return nil, false
}
