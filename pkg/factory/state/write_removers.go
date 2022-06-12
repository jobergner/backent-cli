package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeRemovers() *Factory {
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
					If(List(Id("_"), Id("ok")).Op(":=").Add(r.engine().Dot("Patch").Dot(Title(r.t.Name)).Index(r.elementCore().Dot("ID"))), Op("!").Id("ok")).Block(
						Id("cp").Op(":=").Make(Index().Id(r.valueTypeID()), Len(r.field())),
						Copy(Id("cp"), r.field()),
						r.field().Op("=").Id("cp"),
					),
					For(r.eachElement()).Block(
						If(r.idsDontMatch()...).Block(
							Continue(),
						),
						r.field().Index(Id("i")).Op("=").Add(r.field()).Index(Len(r.field()).Lit(-1)),
						r.field().Index(Len(r.field()).Lit(-1)).Op("=").Lit(0),
						r.field().Op("=").Add(r.field()).Index(Id(""), Len(r.field()).Lit(-1)),
						r.deleteElement(),
						r.elementCore().Dot("OperationKind").Op("=").Id("OperationKindUpdate"),
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
