package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeRemovers() *EngineFactory {
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

				s.file.Func().Params(r.receiverParams()).Id(r.name()).Params(r.params()).Id(r.returns()).Block(
					r.reassignElement(),
					If(r.isOperationKindDelete()).Block(
						Return(Id(configType.Name)),
					),
					For(r.eachElement()).Block(
						If(r.toRemoveComparator().Op("!=").Id(r.toRemoveParamName())).Block(
							Continue(),
						),
						r.field().Index(Id("i")).Op("=").Add(r.field()).Index(Len(r.field()).Lit(-1)),
						r.field().Index(Len(r.field()).Lit(-1)).Op("=").Add(r.zeroValueID()),
						r.field().Op("=").Add(r.field()).Index(Id(""), Len(r.field()).Lit(-1)),
						OnlyIf(!r.v.IsBasicType, r.deleteElement()),
						r.elementCore().Dot("OperationKind").Op("=").Id("OperationKindUpdate"),
						r.elementCore().Dot("Meta").Dot("sign").Call(r.engine().Dot("broadcastingClientID")),
						r.engine().Dot("Patch").Dot(Title(r.t.Name)).Index(r.elementCore().Dot("ID")).Op("=").Add(r.elementCore()),
						Break(),
					),
					Return(Id(configType.Name)),
				)
			})

		})
	})

	return s
}
