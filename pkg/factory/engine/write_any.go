package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeAny() *EngineFactory {
	s.config.RangeAnyFields(func(field ast.Field) {

		k := anyKindWriter{
			f: field,
		}

		s.file.Func().Params(k.receiverParams()).Id("Kind").Params().Id("ElementKind").Block(
			k.reassignAnyContainer(),
			Return(k.containedElementKind()),
		)

		field.RangeValueTypes(func(valueType *ast.ConfigType) {
			asw := anySetterWriter{
				f: field,
				v: *valueType,
			}

			s.file.Func().Params(asw.wrapperReceiverParams()).Id("Set"+Title(valueType.Name)).Params().Id(Title(valueType.Name)).Block(
				asw.reassignAnyContainerWrapper(),
				If(asw.isAlreadyRequestedElement().Op("||").Add(asw.isOperationKindDelete())).Block(
					Return(asw.currentElement()),
				),
				asw.createChild(),
				asw.callSetter(),
				Return(Id(valueType.Name)),
			)

			s.file.Func().Params(asw.receiverParams()).Id("set"+Title(valueType.Name)).Params(asw.params()).Block(
				asw.reassignAnyContainer(),
				If(Id("deleteCurrentChild")).Block(
					ForEachValueOfField(field, func(_valueType *ast.ConfigType) *Statement {
						if _valueType.Name == valueType.Name {
							return Empty()
						}
						asw._v = _valueType
						return If(asw.otherValueIsSet()).Block(
							asw.deleteOtherValue(),
							asw.unsetIDInContainer(),
						)
					}),
				),
				asw.setElementKind(),
				asw.setChildID(),
				asw.updateContainerInPatch(),
			)
		})

		d := anyDeleteChildWriter{
			f: field,
		}
		s.file.Func().Params(d.receiverParams()).Id("deleteChild").Params().Block(
			d.reassignAnyContainer(),
			Switch(Id("any").Dot("ElementKind")).Block(
				ForEachValueOfField(field, func(valueType *ast.ConfigType) *Statement {
					d.v = valueType
					return Case(Id("ElementKind" + Title(valueType.Name))).Block(
						d.deleteChild(),
					)
				}),
			),
		)

	})

	return s
}

func (s *EngineFactory) writeAnyRefs() *EngineFactory {
	s.config.RangeAnyFields(func(field ast.Field) {

		a := anyRefWriter{
			f: field,
		}

		s.file.Type().Id(a.typeRefName()).Struct(
			Id(a.wrapperName()).Id(Title(a.typeName())),
			Id(a.typeName()).Id(a.typeName()+"Core"),
		)

		s.file.Func().Params(a.receiverParams()).Id("Kind").Params().Id("ElementKind").Block(
			Return(a.elementKind()),
		)

		field.RangeValueTypes(func(configType *ast.ConfigType) {
			a.v = configType
			s.file.Func().Params(a.receiverParams()).Id(Title(configType.Name)).Params().Id(Title(configType.Name)).Block(
				Return(a.child()),
			)
		})
	})

	return s
}
