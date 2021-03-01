package statefactory

import (
	"bytes"
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

func newTemplateFrom(name, templateString string) *template.Template {
	return template.Must(
		template.New(name).
			Funcs(template.FuncMap{
				"toTitleCase": strings.Title,
			}).
			Parse(templateString),
	)
}
