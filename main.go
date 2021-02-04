package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func scanFiles(directoryPath string) []inputFile {
	var files []inputFile

	fileInfos, err := ioutil.ReadDir(directoryPath)
	check(err)
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		if filepath.Ext(fileInfo.Name()) != ".go" {
			continue
		}

		filePath := filepath.Join(directoryPath, fileInfo.Name())
		content, err := ioutil.ReadFile(filePath)
		check(err)
		files = append(files, newInputFile(fileInfo.Name(), filePath, content))
	}

	return files
}

func main() {
	inputDirectoryFlag := flag.String("i", "./", "input directory")
	outputFileName := flag.String("o", "stringified_decls.go", "output file")
	packageName := flag.String("p", "main", "package name")
	outputDeclsPrefix := flag.String("prefix", "", "prefix of output declaraton names")
	flag.Parse()

	inputFiles := scanFiles(*inputDirectoryFlag)
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
