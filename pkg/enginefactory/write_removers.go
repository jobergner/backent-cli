package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeRemovers() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {

			if !field.HasSliceValue {
				return
			}

			r := remover{
				t: configType,
				f: field,
			}

			field.RangeValueTypes(func(valueType *ast.ConfigType) {

				r.v = valueType

				decls.File.Func().Params(r.receiverParams()).Id(r.name()).Params(r.params()).Id(r.returns()).Block(
					r.reassignElement(),
					If(r.isOperationKindDelete()).Block(
						Return(Id(configType.Name)),
					),
					OnlyIf(r.f.HasAnyValue && r.f.HasPointerValue, &Statement{
						Id("refs").Op(":=").Make(Map(Id(Title(r.f.ValueTypeName + "ID"))).Id(Title(anyNameByField(r.f) + "ID"))).Line(),
						For(List(Id("_"), Id("refID")).Op(":=").Range().Add(r.elementCore()).Dot(Title(r.f.Name))).Block(
							Id("ref").Op(":=").Add(r.engine()).Dot(r.f.ValueTypeName).Call(Id("refID")),
							Id("refs").Index(Id("refID")).Op("=").Id("ref").Dot(r.f.ValueTypeName).Dot("ReferencedElementID"),
						).Line(),
					}),
					OnlyIf(r.f.HasPointerValue || r.f.HasAnyValue, &Statement{
						Id("wrappers").Op(":=").Make(Map(Id(Title(r.f.ValueTypeName + "ID"))).Id(Title(r.v.Name + "ID"))).Line(),
						For(r.eachWrapper()).Block(
							Id("wrapper").Op(":=").Add(r.getWrapper()),
							OnlyIf(r.f.HasAnyValue, If(Id("wrapper").Dot("Kind").Call()).Op("!=").Id("ElementKind"+Title(r.v.Name)).Block(
								Continue(),
							)),
							Id("wrappers").Index(r.usedWrapperID()).Op("=").Add(r.getElementID()),
						).Line(),
					}),
					For(r.eachElement()).Block(
						OnlyIf(r.f.HasAnyValue || r.f.HasPointerValue, Id(r.v.Name+"ID").Op(":=").Id("wrappers").Index(Id("wrapperID"))),
						If(r.toRemoveComparator().Op("!=").Id(r.toRemoveParamName())).Block(
							Continue(),
						),
						r.field().Index(Id("i")).Op("=").Add(r.field()).Index(Len(r.field()).Lit(-1)),
						r.field().Index(Len(r.field()).Lit(-1)).Op("=").Lit(r.defaultValueForBasicType(r.v.Name)),
						r.field().Op("=").Add(r.field()).Index(Id(""), Len(r.field()).Lit(-1)),
						OnlyIf(!r.v.IsBasicType, r.deleteElement()),
						r.elementCore().Dot("OperationKind").Op("=").Id("OperationKindUpdate"),
						r.engine().Dot("Patch").Dot(Title(r.t.Name)).Index(r.elementCore().Dot("ID")).Op("=").Add(r.elementCore()),
						Break(),
					),
					Return(Id(configType.Name)),
				)
			})

		})
	})

	decls.Render(s.buf)
	return s
}
