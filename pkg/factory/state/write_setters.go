package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeSetters() *Factory {

	RangeBasicTypes(func(b BasicType) {
		s.file.Func().Params(Id("engine").Id("*Engine")).Id("set"+Title(b.Value)).Params(Id("id").Id(Title(b.Value)+"ID"), Id("val").Id(b.Name)).Block(
			Id(b.Value).Op(":=").Id("engine").Dot(b.Value).Call(Id("id")),
			If(Id(b.Value).Dot("OperationKind").Op("==").Id("OperationKindDelete")).Block(
				Return(),
			),
			If(Id(b.Value).Dot("Value").Op("==").Id("val")).Block(
				Return(),
			),
			Id(b.Value).Dot("Value").Op("=").Id("val"),
			Id(b.Value).Dot("OperationKind").Op("=").Id("OperationKindUpdate"),
			Id("engine").Dot("Patch").Dot(Title(b.Value)).Index(Id("id")).Op("=").Id(b.Value),
		)
	})

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
				sw.setAttribute(),
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
				If(srfw.isRefAlreadyAssigned()).Block(
					If(srfw.referenceEquals(), Id(Title(valueType.Name)+"ID").Call(Id("childID")).Op("==").Id(srfw.idParam())).Block(
						Return(Id(field.Parent.Name)),
					),
					srfw.deleteExistingRef(),
				),
				OnlyIf(field.HasAnyValue, srfw.createAnyContainer()),
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
