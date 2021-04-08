package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeSetters() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {

			if field.HasSliceValue || !field.ValueType.IsBasicType {
				return
			}

			s := setter{
				t: configType,
				f: field,
			}

			decls.File.Func().Params(s.receiverParams()).Id(s.name()).Params(s.params()).Id(s.returns()).Block(
				s.reassignElement(),
				If(s.isOperationKindDelete()).Block(
					Return(Id(configType.Name)),
				),
				s.setAttribute(),
				s.setOperationKind(),
				s.updateElementInPatch(),
				Return(Id(configType.Name)),
			)
		})
	})

	decls.Render(s.buf)
	return s
}

type setter struct {
	t ast.ConfigType
	f ast.Field
}

func (s setter) receiverParams() *Statement {
	return Id(s.receiverName()).Id(title(s.t.Name))
}

func (s setter) name() string {
	if s.f.ValueType.IsBasicType {
		return "Set" + title(s.f.Name)
	}
	return "Add" + title(pluralizeClient.Singular(s.f.Name))
}

func (s setter) newValueParam() string {
	return "new" + title(s.f.Name)
}

func (s setter) params() *Statement {
	return List(Id("se").Id("*Engine"), Id(s.newValueParam()).Id(s.f.ValueType.Name))
}

func (s setter) returns() string {
	return title(s.t.Name)
}

func (s setter) reassignElement() *Statement {
	return Id(s.t.Name).Op(":=").Id("se").Dot(title(s.t.Name)).Call(Id(s.receiverName()).Dot(s.t.Name).Dot("ID"))
}

func (s setter) isOperationKindDelete() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("OperationKind_").Op("==").Id("OperationKindDelete")
}

func (s setter) setAttribute() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot(title(s.f.Name)).Op("=").Id(s.newValueParam())
}

func (s setter) setOperationKind() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("OperationKind_").Op("=").Id("OperationKindUpdate")
}

func (s setter) setOperationKindUpdate() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("OperationKind_").Op("=").Id("OperationKindUpdate")
}

func (s setter) updateElementInPatch() *Statement {
	return Id("se").Dot("Patch").Dot(title(s.t.Name)).Index(s.elementID()).Op("=").Id(s.t.Name).Dot(s.t.Name)
}

func (s setter) elementID() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("ID")
}

func (s setter) receiverName() string {
	return "_" + s.t.Name
}
