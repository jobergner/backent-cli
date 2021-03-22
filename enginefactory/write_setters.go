package enginefactory

import (
	"text/template"
)

const setterTemplateString string = `
<( range .Types )><( $Type := . )><( range .Fields )><( if not .HasSliceValue )><( if .ValueType.IsBasicType)>
func (_e <( toTitleCase $Type.Name )>) Set<( toTitleCase .Name )>(se *Engine, new<( toTitleCase .Name )> <( .ValueType.Name )>) <( toTitleCase $Type.Name )> {
	e := se.<( toTitleCase $Type.Name )>(_e.<( $Type.Name )>.ID)
	if e.<( $Type.Name )>.OperationKind_ == OperationKindDelete {
		return e
	}
	e.<( $Type.Name )>.<( toTitleCase .Name )> = new<( toTitleCase .Name )>
	e.<( $Type.Name )>.OperationKind_ = OperationKindUpdate
	se.Patch.<( toTitleCase $Type.Name )>[e.<( $Type.Name )>.ID] = e.<( $Type.Name )>
	return e
}
<( end )><( end )>
<( end )><( end )>
`

var setterTemplate *template.Template = newTemplateFrom("setterTemplate", setterTemplateString)

func (s *stateFactory) writeSetters() *stateFactory {
	setterTemplate.Execute(s.buf, s.ast)
	return s
}
