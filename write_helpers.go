package statefactory

import (
	"text/template"
)

const deduplicateTemplateString string = `
<( range .Types )>
func deduplicate<( toTitleCase .Name )>IDs(a []<( toTitleCase .Name )>ID, b []<( toTitleCase .Name )>ID) []<( toTitleCase .Name )>ID {

	check := make(map[<( toTitleCase .Name )>ID]bool)
	deduped := make([]<( toTitleCase .Name )>ID, 0)
	for _, val := range a {
		check[val] = true
	}
	for _, val := range b {
		check[val] = true
	}

	for letter := range check {
		deduped = append(deduped, letter)
	}

	return deduped
}<( end )>
`

var deduplicateTemplate *template.Template = newTemplateFrom("deduplicateTemplate", deduplicateTemplateString)

func (s *stateFactory) writeDeduplicate() *stateFactory {
	deduplicateTemplate.Execute(s.buf, s.ast)
	return s
}
