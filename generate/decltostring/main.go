package main

import (
	"bytes"
	"flag"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

var (
	packageComment = "// THIS IS A GENERATED FILE. DO NOT EDIT.\n"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func scanFiles(directoryPath, exclude, include, onlyIdentifier string) []inputFile {
	var files []inputFile
	excludeExp := regexp.MustCompile(exclude)
	includeExp := regexp.MustCompile(include)

	fileInfos, err := ioutil.ReadDir(directoryPath)
	check(err)
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		fileName := fileInfo.Name()
		if filepath.Ext(fileName) != ".go" {
			continue
		}
		if exclude != "" && excludeExp.MatchString(fileName) {
			continue
		}
		if onlyIdentifier != "" && fileName != onlyIdentifier {
			continue
		}
		if include != "" && !includeExp.MatchString(fileName) {
			continue
		}

		filePath := filepath.Join(directoryPath, fileName)
		content, err := ioutil.ReadFile(filePath)
		check(err)
		files = append(files, newInputFile(fileName, filePath, content))
	}

	return files
}

func formatCode(code []byte) string {
	ast, err := parser.ParseFile(token.NewFileSet(), "", code, parser.AllErrors)
	check(err)

	var buf bytes.Buffer
	err = format.Node(&buf, token.NewFileSet(), ast)
	check(err)

	return buf.String()
}

func main() {
	inputDirectoryFlag := flag.String("input", "./", "input directory")
	outputFileName := flag.String("output", "stringified_decls.go", "output file")
	packageName := flag.String("package", "main", "package name")
	outputDeclsPrefix := flag.String("prefix", "", "prefix of output declaraton names")
	exclude := flag.String("exclude", "", "files to exclude")
	include := flag.String("include", "", "files to include")
	onlyIdentifier := flag.String("only", "", "the only file to include")
	flag.Parse()

	inputFiles := scanFiles(*inputDirectoryFlag, *exclude, *include, *onlyIdentifier)
	outputFile := newOutputFile(*packageName, *outputDeclsPrefix)
	for _, inputFile := range inputFiles {
		for _, decl := range inputFile.decls {
			outputFile.addDecl(decl, inputFile.name)
		}
	}

	buf := bytes.NewBuffer(nil)

	outputFile.write(buf, *outputFileName)

	content := formatCode(buf.Bytes())

	err := ioutil.WriteFile(*outputFileName, []byte(packageComment+content), 0644)
	check(err)
}
