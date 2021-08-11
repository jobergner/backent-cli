package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeSetters() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {

			if field.HasSliceValue || !field.ValueType().IsBasicType {
				return
			}

			s := setterWriter{
				t: configType,
				f: field,
			}

			decls.File.Func().Params(s.receiverParams()).Id(s.name()).Params(s.params()).Id(s.returns()).Block(
				s.reassignElement(),
				If(s.isOperationKindDelete()).Block(
					Return(Id(configType.Name)),
				),
				If(s.valueHasNotChanged()).Block(
					Return(Id(configType.Name)),
				),
				s.setAttribute(),
				s.setOperationKind(),
				s.updateElementInPatch(),
				Return(Id(configType.Name)),
			)
		})
	})

	s.config.RangeRefFields(func(field ast.Field) {
		if field.HasSliceValue {
			return
		}

		s := setRefFieldWeiter{
			f: field,
		}

		field.RangeValueTypes(func(valueType *ast.ConfigType) {
			s.v = valueType
			decls.File.Func().Params(s.receiverParams()).Id(s.name()).Params(s.params()).Id(s.returns()).Block(
				s.reassignElement(),
				If(s.isOperationKindDelete()).Block(
					Return(Id(field.Parent.Name)),
				),
				If(s.isReferencedElementDeleted()).Block(
					Return(Id(field.Parent.Name)),
				),
				If(s.isSameID()).Block(
					Return(Id(field.Parent.Name)),
				),
				If(s.isRefAlreadyAssigned()).Block(
					s.deleteExistingRef(),
				),
				OnlyIf(field.HasAnyValue, s.createAnyContainer()),
				OnlyIf(field.HasAnyValue, s.setAnyContainer()),
				s.createNewRef(),
				s.setNewRef(),
				s.setOperationKind(),
				s.setItemInPatch(),
				Return(Id(field.Parent.Name)),
			)
		})

	})

	decls.Render(s.buf)
	return s
}
