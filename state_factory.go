package statefactory

import (
	"bytes"
	"github.com/gertd/go-pluralize"
	"go/format"
	"go/parser"
	"go/token"
	"sort"
	"strings"
	"text/template"
)

var pluralizeClient *pluralize.Client = pluralize.NewClient()

type stateFactory struct {
	ast simpleAST
	buf *bytes.Buffer
}

func newStateFactory(ast simpleAST) *stateFactory {
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
				"toSingular":  pluralizeClient.Singular,
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
