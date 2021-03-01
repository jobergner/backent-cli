package statefactory

import (
	"text/template"
)

const entityKindsTemplateString string = `
type EntityKind string

const ({{range .Decls}}
	EntityKind{{ toTitleCase .Name }} EntityKind = "{{.Name}}"{{end}}
)`

var entityKindsTemplate *template.Template = newTemplateFrom("entityKindsTemplate", entityKindsTemplateString)

func (s *stateFactory) writeEntityKinds() *stateFactory {
	entityKindsTemplate.Execute(&s.buf, s.ast)
	return s
}
