package statefactory

import (
	"bytes"
	"sort"
	"strings"
	"text/template"
)

type stateFactory struct {
	ast simpleAST
	buf bytes.Buffer
}

func newStateFactory(ast simpleAST) stateFactory {
	return stateFactory{
		ast: ast,
		buf: bytes.Buffer{},
	}
}

func indexOfDecl(decls map[string]simpleTypeDecl, currentDecl simpleTypeDecl) int {
	var keys []string
	for k := range decls {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var indexOf int
	for i, key := range keys {
		if key == currentDecl.Name {
			indexOf = i
		}
	}
	return indexOf
}

func newTemplateFrom(name, templateString string) *template.Template {
	return template.Must(
		template.New(name).
			Funcs(template.FuncMap{
				"toTitleCase": strings.Title,
				"doNotWriteOnIndex": func(decls map[string]simpleTypeDecl, currentDecl simpleTypeDecl, requiredIndex int, toWrite string) string {
					currentIndex := indexOfDecl(decls, currentDecl)
					if requiredIndex < 0 {
						if currentIndex == len(decls)+requiredIndex {
							return ""
						}
					} else {

						if currentIndex == requiredIndex {
							return ""
						}
					}
					return toWrite
				},
				"writeOnIndex": func(decls map[string]simpleTypeDecl, currentDecl simpleTypeDecl, requiredIndex int, toWrite string) string {
					currentIndex := indexOfDecl(decls, currentDecl)
					if requiredIndex < 0 {
						if currentIndex == len(decls)+requiredIndex {
							return toWrite
						}
					} else {

						if currentIndex == requiredIndex {
							return toWrite
						}
					}
					return ""
				},
			}).
			Delims("<(", ")>").
			Parse(templateString),
	)
}
