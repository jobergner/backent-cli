package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"
)

var packageName = flag.String("pkg", "webclient", "name of package of out file")
var inputFilePath = flag.String("in", "./examples/webclient/index.ts", "path of input file")
var outputFilePath = flag.String("out", "./pkg/factory/webclient/stringified_decls.go", "path of output file")

var literalMatcher = regexp.MustCompile(`([a-zA-Z])\w*`)

func declarationName(decl string) string {
	matches := literalMatcher.FindAllString(decl, 3)

	if matches[0] == "export" {
		if len(matches) < 3 {
			panic(fmt.Sprintf("found < 3 matches for this decl:\n%s", decl))
		}
		return fmt.Sprintf("%s_%s", matches[1], matches[2])
	}

	if len(matches) < 2 {
		panic(fmt.Sprintf("found < 2 matches for this decl:\n%s", decl))
	}
	return fmt.Sprintf("%s_%s", matches[0], matches[1])
}

func main() {
	flag.Parse()

	content, err := os.ReadFile(*inputFilePath)
	if err != nil {
		panic(err)
	}

	decls := strings.Split(string(content), "\n\n")

	file := jen.NewFile(*packageName)
	collectionContent := make(jen.Dict)
	for _, d := range decls {
		declName := declarationName(d)
		collectionContent[jen.Lit(declName)] = jen.Id(declName)
		file.Const().Id(declName).Op("=").Id("`" + escapeBackticks(strings.TrimSpace(d)) + "`")
	}

	file.Var().Id("decl_to_string_decl_collection").Op("=").Map(jen.String()).String().Values(collectionContent)

	buf := bytes.NewBuffer(nil)
	file.Render(buf)
	if err := os.WriteFile(*outputFilePath, buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func escapeBackticks(s string) string {
	return strings.Replace(s, "`", "` + \"`\" +  `", -1)
}
