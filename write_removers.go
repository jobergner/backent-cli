package statefactory

import (
	"text/template"
)

const removerTemplateString string = `
<( range .Decls )><( $Decl := . )><( range .Fields )><( if .HasSliceValue )>
func (_e <( toTitleCase $Decl.Name )>) Remove<(if .ValueType.IsBasicType )><( toTitleCase .Name )><( else )><( toTitleCase .ValueType.Name )>s<( end )>(sm *StateMachine, <(if .ValueType.IsBasicType )><( .Name )>ToRemove ...<( .ValueType.Name )><( else )> <( .ValueType.Name )>sToRemove ...<( toTitleCase .ValueType.Name )>ID<( end )>) <( toTitleCase $Decl.Name )> {
	e := sm.Get<( toTitleCase $Decl.Name )>(_e.<( $Decl.Name )>.ID)
	if e.<( $Decl.Name )>.OperationKind == OperationKindDelete {
		return e
	}
	var elementsAltered bool
	var newElements []<(if .ValueType.IsBasicType )><( .ValueType.Name )><( else )><( toTitleCase .ValueType.Name )>ID<( end )>
	for _, element := range e.<( $Decl.Name )>.<( toTitleCase .Name )> {
		var toBeRemoved bool
		for _, elementToRemove := range <(if .ValueType.IsBasicType )><( .Name )>ToRemove<( else )> <( .ValueType.Name )>sToRemove<( end )> {
			if element == elementToRemove {
				toBeRemoved = true
				elementsAltered = true<(if not .ValueType.IsBasicType )>
				sm.Delete<( toTitleCase .ValueType.Name )>(element)<( end )>
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !elementsAltered {
		return e
	}
	e.<( $Decl.Name )>.<( toTitleCase .Name )> = newElements
	e.<( $Decl.Name )>.OperationKind = OperationKindUpdate
	sm.Patch.<( toTitleCase $Decl.Name )>[e.<( $Decl.Name )>.ID] = e.<( $Decl.Name )>
	return e
}
<( end )>
<( end )><( end )>
`

var removerTemplate *template.Template = newTemplateFrom("removerTemplate", removerTemplateString)

func (s *stateFactory) writeRemovers() *stateFactory {
	removerTemplate.Execute(&s.buf, s.ast)
	return s
}
