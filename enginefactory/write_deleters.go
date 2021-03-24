package enginefactory

import (
	. "github.com/dave/jennifer/jen"
	"text/template"
)

const deleterTemplateString string = `
<( range .Types )>
func (se *Engine) Delete<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) {
	<( if not .IsRootType -)>
		<( encrypt .Name )> := se.<( toTitleCase .Name )>(<( .Name )>ID).<( .Name )>
		if <( encrypt .Name )>.HasParent_ {
			return
		}
	<( end -)>
	se.delete<( toTitleCase .Name )>(<( .Name )>ID)
}
func (se *Engine) delete<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) {
	<( encrypt .Name )> := se.<( toTitleCase .Name )>(<( .Name )>ID).<( .Name )>
	<( encrypt .Name )>.OperationKind_ = OperationKindDelete
	se.Patch.<( toTitleCase .Name )>[<( encrypt .Name )>.ID] = <( encrypt .Name -)>
	<( $Type := . )>
	<(- range .Fields )><( if not .ValueType.IsBasicType )>
		<(- if .HasSliceValue )>
			for _, <( .ValueType.Name )>ID := range <( encrypt $Type.Name )>.<( toTitleCase .Name )> {
				se.delete<( toTitleCase .ValueType.Name )>(<( .ValueType.Name )>ID)
			}
		<(- else )>
			se.delete<( toTitleCase .ValueType.Name )>(<( encrypt $Type.Name )>.<( toTitleCase .ValueType.Name )>)
		<(- end -)>
	<(- end )><(- end )>
}
<( end )>
`

var deleterTemplate *template.Template = newTemplateFrom("deleterTemplate", deleterTemplateString)

func (s *stateFactory) writeDeleters() *stateFactory {
	decls := newDeclSet()
	s.ast.rangeTypes(func(configType stateConfigType) {
		configType.rangeFields(func(field stateConfigField) {

			if field.HasSliceValue || !field.ValueType.IsBasicType {
				return
			}

			t := typeDeleterWrapper{
				t: configType,
			}

			decls.file.Func().Params(t.receiverParams()).Id(t.name()).Params(t.params()).Block(
				onlyIf(!configType.IsRootType, t.reassignElement()),
				If(t.hasParent()).Block(
					Return(),
				),
				t.deleteElement(),
			)

			td := typeDeleter{
				t: configType,
			}

			decls.file.Func().Params(td.receiverParams()).Id(td.name()).Params(td.params()).Block(
				onlyIf(!configType.IsRootType, td.reassignElement()),
				If(td.hasParent()).Block(
					Return(),
				),
				td.deleteElement(),
			)
		})
	})

	decls.render(s.buf)
	return s
}

type typeDeleterWrapper struct {
	t stateConfigType
}

func (t typeDeleterWrapper) receiverParams() *Statement {
	return Id("se").Id("*Engine")
}

func (t typeDeleterWrapper) name() string {
	return "Delete" + title(t.t.Name)
}

func (t typeDeleterWrapper) idParam() string {
	return t.t.Name + "ID"
}

func (t typeDeleterWrapper) params() *Statement {
	return Id(t.idParam()).Id(title(t.t.Name) + "ID")
}

func (t typeDeleterWrapper) reassignElement() *Statement {
	return Id(t.t.Name).Op(":=").Id("se").Dot(title(t.t.Name)).Params(Id(t.idParam())).Dot(t.t.Name)
}

func (t typeDeleterWrapper) hasParent() *Statement {
	return Id(t.t.Name).Dot("HasParent_")
}

func (t typeDeleterWrapper) deleteElement() *Statement {
	return Id("se").Dot("delete" + title(t.t.Name)).Params(Id(t.idParam()))
}

type typeDeleter struct {
	t stateConfigType
	f stateConfigField
}

func (td typeDeleter) receiverParams() *Statement {
	return Id("se").Id("*Engine")
}

func (td typeDeleter) name() string {
	return "delete" + title(td.t.Name)
}

func (td typeDeleter) idParam() string {
	return td.t.Name + "ID"
}

func (td typeDeleter) params() *Statement {
	return Id(td.idParam()).Id(title(td.t.Name) + "ID")
}

func (td typeDeleter) reassignElement() *Statement {
	return Id(td.t.Name).Op(":=").Id("se").Dot(title(td.t.Name)).Params(Id(td.idParam())).Dot(td.t.Name)
}

func (td typeDeleter) setOperationKind() *Statement {
	return Id(td.t.Name).Dot("OperationKind_").Op("=").Id("OperationKindUpdate")
}

func (td typeDeleter) updateElementInPatch() *Statement {
	return Id("se").Dot("Patch").Dot(title(td.t.Name)).Index(Id(td.t.Name).Dot("ID")).Op("=").Id(td.t.Name)
}

func (td *typeDeleter) withField(field stateConfigField) *typeDeleter {
	td.f = field
	return td
}

func (td typeDeleter) loopConditions() *Statement {
	return List(Id("_"), Id(td.f.ValueType.Name+"ID")).Op(":=").Range().Id(td.t.Name).Dot(title(td.f.Name))
}

func (td typeDeleter) deleteElementInLoop() *Statement {
	return Id("se").Dot("delete" + title(td.f.ValueType.Name)).Params(Id(td.f.ValueType.Name + "ID"))
}

func (td typeDeleter) deleteElement() *Statement {
	return Id("se").Dot("delete" + title(td.f.ValueType.Name)).Params(Id(td.t.Name).Dot(title(td.f.ValueType.Name)))
}
