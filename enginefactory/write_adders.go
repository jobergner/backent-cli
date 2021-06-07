package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeAdders() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {

			if !field.HasSliceValue {
				return
			}

			a := adderWriter{
				t: configType,
				f: field,
				v: nil,
			}

			field.RangeValueTypes(func(valueType *ast.ConfigType) {
				a.v = valueType
				decls.File.Func().Params(a.receiverParams()).Id(a.name()).Params(a.params()).Id(a.returns()).Block(
					a.reassignElement(),
					If(a.isOperationKindDelete()).Block(
						Return(a.returnDeletedElement()),
					),
					OnlyIf(field.HasPointerValue, If(a.referencedElementDoesntExist()).Block(
						Return(),
					)),
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

	decls.Render(s.buf)
	return s
}
