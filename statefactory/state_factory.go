package statefactory

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"sort"
	"strings"
	"text/template"

	"github.com/gertd/go-pluralize"
)

var pluralizeClient *pluralize.Client = pluralize.NewClient()

type stateFactory struct {
	ast stateConfigAST
	buf *bytes.Buffer
}

func newStateFactory(ast stateConfigAST) *stateFactory {
	return &stateFactory{
		ast: ast,
		buf: &bytes.Buffer{},
	}
}

func (s *stateFactory) prependPackage() *stateFactory {
	s.buf = bytes.NewBufferString("package statemachine\n" + s.buf.String())
	return s
}

func (s *stateFactory) format() error {
	ast, err := parser.ParseFile(token.NewFileSet(), "", s.buf.String(), parser.AllErrors)
	if err != nil {
		return err
	}

	s.buf.Reset()
	err = format.Node(s.buf, token.NewFileSet(), ast)
	return err
}

func indexOfType(configTypes map[string]stateConfigType, currentConfigType stateConfigType) int {
	var keys []string
	for k := range configTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var indexOf int
	for i, key := range keys {
		if key == currentConfigType.Name {
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
				"toSingular":  pluralizeClient.Singular,
				"doNotWriteOnIndex": func(configTypes map[string]stateConfigType, currentConfigType stateConfigType, requiredIndex int, toWrite string) string {
					currentIndex := indexOfType(configTypes, currentConfigType)
					if requiredIndex < 0 {
						if currentIndex == len(configTypes)+requiredIndex {
							return ""
						}
					} else {

						if currentIndex == requiredIndex {
							return ""
						}
					}
					return toWrite
				},
				"writeOnIndex": func(configTypes map[string]stateConfigType, currentConfigType stateConfigType, requiredIndex int, toWrite string) string {
					currentIndex := indexOfType(configTypes, currentConfigType)
					if requiredIndex < 0 {
						if currentIndex == len(configTypes)+requiredIndex {
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
