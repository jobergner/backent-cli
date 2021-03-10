package statefactory

import (
	"bytes"
	"github.com/gertd/go-pluralize"
	"sort"
	"strings"
	"text/template"
)

var pluralizeClient *pluralize.Client = pluralize.NewClient()

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

// TODO prettier
func newTemplateFrom(name, templateString string) *template.Template {
	return template.Must(
		template.New(name).
			Funcs(template.FuncMap{
				"toTitleCase": strings.Title,
				"toSingular":  pluralizeClient.Singular,
				"toFieldValue": func(field simpleFieldDecl) string {
					var valueStringWriter bytes.Buffer
					if field.HasSliceValue {
						valueStringWriter.WriteString("[]")
					}
					if field.ValueType.IsBasicType {
						valueStringWriter.WriteString(field.ValueType.Name)
					} else {
						valueStringWriter.WriteString(strings.Title(field.ValueType.Name) + "ID")
					}
					return valueStringWriter.String()
				},
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
