package statefactory

import (
	"text/template"
)

const adderTemplateString string = `
<( range .Types )><( $Type := . )><( range .Fields )><( if .HasSliceValue )>
func (_e <( toTitleCase $Type.Name )>) Add
<(- if .ValueType.IsBasicType -)>
	<( toTitleCase .Name )>
<(- else -)>
	<( .Name | toSingular | toTitleCase )>
<(- end -)>
(sm *StateMachine<(if .ValueType.IsBasicType )>, <( .Name)> ...<( .ValueType.Name )><( end )>) <( if not .ValueType.IsBasicType )><( toTitleCase .ValueType.Name )><( end )> {
	e := sm.<( toTitleCase $Type.Name )>(_e.<( $Type.Name )>.ID)
	if e.<( $Type.Name )>.OperationKind == OperationKindDelete {
		return<( if not .ValueType.IsBasicType )> <( toTitleCase .ValueType.Name )>{<( .ValueType.Name )>Core{OperationKind: OperationKindDelete}}<( end )>
	}
	<(- if not .ValueType.IsBasicType )>
		<( .ValueType.Name )> := sm.create<( toTitleCase .ValueType.Name )>(true)
	<(- end )>
	e.<( $Type.Name )>.<( toTitleCase .Name )> = append(e.<( $Type.Name )>.<( toTitleCase .Name )>,<(print " ")>
	<(- if .ValueType.IsBasicType -)>
		<( .Name )>...
	<(- else -)>
		<( .ValueType.Name )>.<( .ValueType.Name )>.ID
	<(- end -)>
	)
	e.<( $Type.Name )>.OperationKind = OperationKindUpdate
	sm.Patch.<( toTitleCase $Type.Name )>[e.<( $Type.Name )>.ID] = e.<( $Type.Name )><( if not .ValueType.IsBasicType )>
	return <( .ValueType.Name )><( end )>
}
<( end )>
<( end )><( end )>
`

var adderTemplate *template.Template = newTemplateFrom("adderTemplate", adderTemplateString)

func (s *stateFactory) writeAdders() *stateFactory {
	adderTemplate.Execute(s.buf, s.ast)
	return s
}
