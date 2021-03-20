package enginefactory

import (
	"text/template"
)

const getterTemplateString string = `
<( define "returnValue" -)>
	<( if .HasSliceValue -)>
		[]
	<(- end -)>
	<( if .ValueType.IsBasicType -)>
		<( .ValueType.Name )>
	<(- else -)>
		<( toTitleCase .ValueType.Name )>
	<(- end )>
<(- end )>
<( range .Types )>
func (se *Engine) <( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) <( toTitleCase .Name )> {
	patching<( toTitleCase .Name )>, ok := se.Patch.<( toTitleCase .Name )>[<( .Name )>ID]
	if ok {
		return <( toTitleCase .Name )>{patching<( toTitleCase .Name )>}
	}
	current<( toTitleCase .Name )> := se.State.<( toTitleCase .Name )>[<( .Name )>ID]
	return <( toTitleCase .Name )>{current<( toTitleCase .Name )>}
}
func (_e <( toTitleCase .Name )>) ID(se *Engine) <( toTitleCase .Name )>ID {
	return _e.<( .Name )>.ID
}
<( $Type := . )><( range .Fields )>
func (_e <( toTitleCase $Type.Name )>) <( toTitleCase .Name )>(se *Engine) <( template "returnValue" . )> {
	e := se.<( toTitleCase $Type.Name )>(_e.<( $Type.Name )>.ID)
	<( if .HasSliceValue -)>
		var <( .Name )> <( template "returnValue" . )>
		for _,<( print " " )>
		<(- if .ValueType.IsBasicType -)>
			element
		<(- else -)>
			<( .ValueType.Name )>ID
		<(- end -)>
		<( print " " )>:= range e.<( $Type.Name )>.<( toTitleCase .Name )> {
			<( .Name )> = append(<( .Name )>,<( print " " )>
			<(- if .ValueType.IsBasicType -)>
				element
			<(- else -)>
				se.<( toTitleCase .ValueType.Name )>(<( .ValueType.Name )>ID)
			<(- end -)>
			)
		}
		return <( .Name )>
	<(- else -)>
		<( if .ValueType.IsBasicType -)>
			return e.<( $Type.Name )>.<( toTitleCase .Name )>
		<(- else -)>
			return se.<( toTitleCase .Name )>(e.<( $Type.Name )>.<( toTitleCase .Name )>)
		<(- end -)>
	<(- end )>
}
<( end )>
<( end )>
`

var getterTemplate *template.Template = newTemplateFrom("getterTemplate", getterTemplateString)

func (s *stateFactory) writeGetters() *stateFactory {
	getterTemplate.Execute(s.buf, s.ast)
	return s
}
