package main

import (
	"bytes"
	"go/ast"
	"strings"

	"github.com/dave/jennifer/jen"
)

type outputFile struct {
	writtenDeclNames []string
	declPrefix       string
	file             *jen.File
}

func newOutputFile(packageName, declPrefix string) *outputFile {
	file := jen.NewFile(packageName)
	return &outputFile{nil, declPrefix, file}
}

func escapeBackticks(s string) string {
	return strings.Replace(s, "`", "` + \"`\" +  `", -1)
}

func (o *outputFile) newDeclName(decl ast.Decl, containingFileName string) string {
	outputDeclName := evalDeclName(decl, containingFileName)

	if o.declPrefix != "" {
		outputDeclName = o.declPrefix + outputDeclName
	}

	return outputDeclName
}

func (o *outputFile) addDecl(decl ast.Decl, containingFileName string) {

	outputDeclName := o.newDeclName(decl, containingFileName)
	outputDeclValue := escapeBackticks(stringifyDecl(decl))

	o.file.Const().Id(outputDeclName).String().Op("=").Id("`" + outputDeclValue + "`")

	if isImportDecl(decl) { // we currently don't care about import decls
		return
	}

	o.writtenDeclNames = append(o.writtenDeclNames, outputDeclName)
}

func (o *outputFile) write(buf *bytes.Buffer, fileName string) {
	mapContent := make(jen.Dict)
	for _, n := range o.writtenDeclNames {
		mapContent[jen.Lit(n)] = jen.Id(n)
	}

	o.file.Var().Id(o.declPrefix + "decl_to_string_decl_collection").Op("=").Map(jen.String()).String().Values(mapContent)

	o.file.Render(buf)
}
