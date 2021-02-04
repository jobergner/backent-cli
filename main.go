package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func scanFiles(directoryPath, excludedIdentifier string) []inputFile {
	var files []inputFile
	excludeExp := regexp.MustCompile(excludedIdentifier)

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
		if excludeExp.MatchString(fileName) {
			continue
		}

		filePath := filepath.Join(directoryPath, fileName)
		content, err := ioutil.ReadFile(filePath)
		check(err)
		files = append(files, newInputFile(fileName, filePath, content))
	}

	return files
}

func main() {
	inputDirectoryFlag := flag.String("input", "./", "input directory")
	outputFileName := flag.String("output", "stringified_decls.go", "output file")
	packageName := flag.String("package", "main", "package name")
	outputDeclsPrefix := flag.String("prefix", "", "prefix of output declaraton names")
	excludedIdentifier := flag.String("exclude", "$^", "files to exclude")
	flag.Parse()

	inputFiles := scanFiles(*inputDirectoryFlag, *excludedIdentifier)
	outputFile := newOutputFile(*outputFileName, *outputDeclsPrefix, *packageName)
	for _, inputFile := range inputFiles {
		for _, decl := range inputFile.decls {
			outputFile.writeDecl(decl, inputFile.name)
		}
	}

	outputFile.formatContent()
	err := ioutil.WriteFile(outputFile.name, outputFile.content.Bytes(), 0644)
	check(err)
}
