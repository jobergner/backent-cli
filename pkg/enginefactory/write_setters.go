package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeSetters() *EngineFactory {
	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {

			if field.HasSliceValue || !field.ValueType().IsBasicType {
				return
			}

			sw := setterWriter{
				t: configType,
				f: field,
			}

			s.file.Func().Params(sw.receiverParams()).Id(sw.name()).Params(sw.params()).Id(sw.returns()).Block(
				sw.reassignElement(),
				If(sw.isOperationKindDelete()).Block(
					Return(Id(configType.Name)),
				),
				If(sw.valueHasNotChanged()).Block(
					Return(Id(configType.Name)),
				),
				sw.setAttribute(),
				sw.setOperationKind(),
				sw.updateElementInPatch(),
				Return(Id(configType.Name)),
			)
		})
	})

	s.config.RangeRefFields(func(field ast.Field) {
		if field.HasSliceValue {
			return
		}

		srfw := setRefFieldWeiter{
			f: field,
		}

		field.RangeValueTypes(func(valueType *ast.ConfigType) {
			srfw.v = valueType
			s.file.Func().Params(srfw.receiverParams()).Id(srfw.name()).Params(srfw.params()).Id(srfw.returns()).Block(
				srfw.reassignElement(),
				If(srfw.isOperationKindDelete()).Block(
					Return(Id(field.Parent.Name)),
				),
				If(srfw.isReferencedElementDeleted()).Block(
					Return(Id(field.Parent.Name)),
				),
				If(srfw.isSameID()).Block(
					Return(Id(field.Parent.Name)),
				),
				If(srfw.isRefAlreadyAssigned()).Block(
					srfw.deleteExistingRef(),
				),
				OnlyIf(field.HasAnyValue, srfw.createAnyContainer()),
				OnlyIf(field.HasAnyValue, srfw.setAnyContainer()),
				srfw.createNewRef(),
				srfw.setNewRef(),
				srfw.setOperationKind(),
				srfw.setItemInPatch(),
				Return(Id(field.Parent.Name)),
			)
		})

	})

	return s
}
