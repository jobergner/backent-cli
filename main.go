package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type file struct {
	name string
	path string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func scanFiles(directoryPath string) []file {
	var files []file

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		files = append(files, file{name: info.Name(), path: path})
		return nil
	})

	check(err)

	return files
}

func evalDeclName(decl ast.Decl, containingFile file) string {
	if isImportDecl(decl) {
		name := strings.TrimSuffix(containingFile.name, filepath.Ext(containingFile.name))
		splitByDot := strings.Split(name, ".")
		rejoined := strings.Join(splitByDot, "_")
		return rejoined + "_import"
	}
	if isFuncDecl(decl) {
		return getFuncName(decl.(*ast.FuncDecl))
	}
	if isGenDecl(decl) {
		return getGenDeclName(decl.(*ast.GenDecl))
	}
	panic("unknown decl kind")
}

func printDecl(decl ast.Decl) string {
	var buf bytes.Buffer
	// ast.Print(token.NewFileSet(), decl)
	printer.Fprint(&buf, token.NewFileSet(), decl)
	return buf.String()
}

func formatFileContent(content string) string {
	f, err := parser.ParseFile(token.NewFileSet(), "", content, 0)
	check(err)
	var buf bytes.Buffer
	// ast.Print(token.NewFileSet(), decl)
	printer.Fprint(&buf, token.NewFileSet(), f)
	return buf.String()
}

func forEachDeclInFile(file file, fn func(decl ast.Decl)) {
	content, err := ioutil.ReadFile(file.path)
	check(err)
	f, err := parser.ParseFile(token.NewFileSet(), "", content, 0)
	check(err)
	for _, decl := range f.Decls {
		fn(decl)
	}
}

type outputDeclaration struct {
	name  string
	value string
}

func writeToOutputFile(outputDecls []outputDeclaration) {
	var buf bytes.Buffer

	buf.WriteString("package main\n\n")
	for _, outputDecl := range outputDecls {
		escapedValue := strings.Replace(outputDecl.value, "`", "` + \"`\" +  `", -1)
		buf.WriteString("\n\nconst " + outputDecl.name + " string = `" + escapedValue + "`")
	}

	fmt.Println(formatFileContent(buf.String()))
}

func main() {
	inputDirectoryFlag := flag.String("i", "./", "input directory")
	flag.Parse()
	files := scanFiles(*inputDirectoryFlag)

	var outputDecls []outputDeclaration
	for _, file := range files {
		forEachDeclInFile(file, func(decl ast.Decl) {
			newOutputDecl := outputDeclaration{
				name:  evalDeclName(decl, file),
				value: printDecl(decl),
			}
			outputDecls = append(outputDecls, newOutputDecl)
		})
	}

	writeToOutputFile(outputDecls)
}

func isImportDecl(decl ast.Decl) bool {
	if genDecl, ok := decl.(*ast.GenDecl); ok {
		if genDecl.Tok == token.IMPORT {
			return true
		}
	}
	return false
}

func isFuncDecl(decl ast.Decl) bool {
	if _, ok := decl.(*ast.FuncDecl); ok {
		return true
	}
	return false
}

func isGenDecl(decl ast.Decl) bool {
	if _, ok := decl.(*ast.GenDecl); ok {
		return true
	}
	return false
}

func getFuncName(decl *ast.FuncDecl) string {
	return decl.Name.Name
}

func getGenDeclName(decl *ast.GenDecl) string {
	return decl.Specs[0].(*ast.TypeSpec).Name.Name
}
