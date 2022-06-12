package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeAdders() *Factory {
	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {

			if !field.HasSliceValue {
				return
			}

			field.RangeValueTypes(func(valueType *ast.ConfigType) {

				a := adderWriter{
					t: configType,
					f: field,
					v: valueType,
				}

				s.file.Func().Params(a.receiverParams()).Id(a.name()).Params(a.params()).Id(a.returns()).Block(
					a.reassignElement(),
					If(a.isOperationKindDelete()).Block(
						Return(a.returnDeletedElement()),
					),
					OnlyIf(field.HasPointerValue, If(a.referencedElementDoesntExist()).Block(
						Return(),
					)),
					If(List(Id("_"), Id("ok")).Op(":=").Add(a.engine().Dot("Patch").Dot(Title(a.t.Name)).Index(a.elementCore().Dot("ID"))), Op("!").Id("ok")).Block(
						Id("cp").Op(":=").Make(Index().Id(a.valueTypeID()), Len(a.field())),
						Copy(Id("cp"), a.field()),
						a.field().Op("=").Id("cp"),
					),
					OnlyIf(field.HasPointerValue, a.returnIfReferencedElementIsAlreadyReferenced()),
					OnlyIf(!field.HasPointerValue, a.createNewElement()),
					OnlyIf(field.HasAnyValue, &Statement{
						a.createAnyContainer().Line(),
					}),
					OnlyIf(field.HasPointerValue, a.createRef()),
					a.appendElement(),
					a.setOperationKindUpdate(),
					a.updateElementInPatch(),
					OnlyIf(!valueType.IsBasicType && !field.HasPointerValue, Return(Id(valueType.Name))),
				)
			})
		})
	})

	return s
}
