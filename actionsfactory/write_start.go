package actionsfactory

import (
	"text/template"
)

const startTemplateString string = `
func main() {
	err := statemachine.Start(<(- range .Actions )>
	<( .Name )>,
<(- end )>
	)

	if err != nil {
		panic(err)
	}
}
`

var startTemplate *template.Template = newTemplateFrom("actionTemplate", startTemplateString)

func (a *actionsFactory) writeStart() *actionsFactory {
	startTemplate.Execute(a.buf, a.ast)
	return a
}
