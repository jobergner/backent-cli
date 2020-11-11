package main

const main_import string = `import (
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
)`

const file string = `type file struct {
	name	string
	path	string
}`

const check string = `func check(err error) {
	if err != nil {
		panic(err)
	}
}`

const scanFiles string = `func scanFiles(directoryPath string) []file {
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
}`

const evalDeclName string = `func evalDeclName(decl ast.Decl, containingFile file) string {
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
}`

const printDecl string = `func printDecl(decl ast.Decl) string {
	var buf bytes.Buffer
	printer.Fprint(&buf, token.NewFileSet(), decl)
	return buf.String()
}`

const forEachDeclInFile string = `func forEachDeclInFile(file file, fn func(decl ast.Decl)) {
	content, err := ioutil.ReadFile(file.path)
	check(err)
	f, err := parser.ParseFile(token.NewFileSet(), "", content, 0)
	check(err)
	for _, decl := range f.Decls {
		fn(decl)
	}
}`

const outputDeclaration string = `type outputDeclaration struct {
	name	string
	value	string
}`

const writeToOutputFile string = `func writeToOutputFile(outputDecls []outputDeclaration) {
	var buf bytes.Buffer
	buf.WriteString("package main\n\n")
	for _, od := range outputDecls {
		buf.WriteString("\n\nconst " + od.name + " string = ` + "`" + ` " + od.value + )
	}
	fmt.Println(buf.String())
}`

const main string = `func main() {
	inputDirectoryFlag := flag.String("i", "./", "input directory")
	flag.Parse()
	files := scanFiles(*inputDirectoryFlag)
	var outputDecls []outputDeclaration
	for _, file := range files {
		forEachDeclInFile(file, func(decl ast.Decl) {
			newOutputDecl := outputDeclaration{name: evalDeclName(decl, file), value: printDecl(decl)}
			outputDecls = append(outputDecls, newOutputDecl)
		})
	}
	writeToOutputFile(outputDecls)
}`

const isImportDecl string = `func isImportDecl(decl ast.Decl) bool {
	if genDecl, ok := decl.(*ast.GenDecl); ok {
		if genDecl.Tok == token.IMPORT {
			return true
		}
	}
	return false
}`

const isFuncDecl string = `func isFuncDecl(decl ast.Decl) bool {
	if _, ok := decl.(*ast.FuncDecl); ok {
		return true
	}
	return false
}`

const isGenDecl string = `func isGenDecl(decl ast.Decl) bool {
	if _, ok := decl.(*ast.GenDecl); ok {
		return true
	}
	return false
}`

const getFuncName string = `func getFuncName(decl *ast.FuncDecl) string {
	return decl.Name.Name
}`

const getGenDeclName string = `func getGenDeclName(decl *ast.GenDecl) string {
	return decl.Specs[0].(*ast.TypeSpec).Name.Name
}`
