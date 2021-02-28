package statefactory

import (
	"bytes"
	"fmt"
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

const entityKindsTemplateString string = `
type EntityKind string

const ({{range $index, $element := .Decls}}
	EntityKind{{ .Name }} EntityKind = "{{.Name}}"{{end}}
)
`

var entityKindsTemplate *template.Template = template.Must(
	template.New("entityKindsTemplate").Parse(entityKindsTemplateString),
)

func (s *stateFactory) writeEntityKinds() *stateFactory {
	entityKindsTemplate.Execute(&s.buf, s.ast)
	fmt.Println(s.buf.String())
	return s
}
