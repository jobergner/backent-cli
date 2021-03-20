package enginefactory

import (
	"text/template"
)

const removerTemplateString string = `
<( range .Types )><( $Type := . )><( range .Fields )><( if .HasSliceValue )>
func (_e <( toTitleCase $Type.Name )>) Remove<( toTitleCase .Name )>(se *Engine,<( print " " )>
<(- if .ValueType.IsBasicType -)>
	<( .Name )>ToRemove ...<( .ValueType.Name )>
<(- else -)>
	<( .ValueType.Name )>sToRemove ...<( toTitleCase .ValueType.Name )>ID
<(- end -)>
) <( toTitleCase $Type.Name )> {
	e := se.<( toTitleCase $Type.Name )>(_e.<( $Type.Name )>.ID)
	if e.<( $Type.Name )>.OperationKind == OperationKindDelete {
		return e
	}
	var elementsAltered bool
	var newElements []
	<(- if .ValueType.IsBasicType -)>
		<( .ValueType.Name )>
	<(- else -)>
		<( toTitleCase .ValueType.Name )>ID
	<(- end )>
	for _, element := range e.<( $Type.Name )>.<( toTitleCase .Name )> {
		var toBeRemoved bool
		for _, elementToRemove := range<(print " ")>
		<(- if .ValueType.IsBasicType -)>
			<( .Name )>ToRemove
		<(- else -)>
			<( .ValueType.Name )>sToRemove
		<(- end )> {
			if element == elementToRemove {
				toBeRemoved = true
				elementsAltered = true<(if not .ValueType.IsBasicType )>
				se.delete<( toTitleCase .ValueType.Name )>(element)<( end )>
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !elementsAltered {
		return e
	}
	e.<( $Type.Name )>.<( toTitleCase .Name )> = newElements
	e.<( $Type.Name )>.OperationKind = OperationKindUpdate
	se.Patch.<( toTitleCase $Type.Name )>[e.<( $Type.Name )>.ID] = e.<( $Type.Name )>
	return e
}
<( end )>
<( end )><( end )>
`

var removerTemplate *template.Template = newTemplateFrom("removerTemplate", removerTemplateString)

func (s *stateFactory) writeRemovers() *stateFactory {
	removerTemplate.Execute(s.buf, s.ast)
	return s
}
