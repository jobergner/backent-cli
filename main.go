package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func scanFiles(directoryPath string) []inputFile {
	var files []inputFile

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		check(err)
		files = append(files, newInputFile(info.Name(), path, content))
		return nil
	})

	check(err)

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
