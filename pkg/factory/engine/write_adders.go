package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeAdders() *EngineFactory {
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
					OnlyIf(field.HasPointerValue, a.returnIfReferencedElementIsAlreadyReferenced()),
					OnlyIf(!valueType.IsBasicType && !field.HasPointerValue, a.createNewElement()),
					OnlyIf(field.HasAnyValue, &Statement{
						a.createAnyContainer().Line(),
						a.setAnyContainer(),
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
