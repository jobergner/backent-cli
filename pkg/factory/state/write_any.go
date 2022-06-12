package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeAny() *Factory {
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

			s.file.Func().Params(asw.wrapperReceiverParams()).Id("Be"+Title(valueType.Name)).Params().Id(Title(valueType.Name)).Block(
				asw.reassignAnyContainerWrapper(),
				If(asw.isAlreadyRequestedElement().Op("||").Add(asw.isOperationKindDelete())).Block(
					Return(asw.currentElement()),
				),
				asw.createChild(),
				asw.callSetter(),
				Return(Id(valueType.Name)),
			)

			s.file.Func().Params(asw.receiverParams()).Id("be"+Title(valueType.Name)).Params(asw.params()).Block(
				asw.reassignAnyContainer(),
				Id("any").Dot("engine").Dot("delete"+Title(anyNameByField(field))).Call(Id("any").Dot("ID"), Id("deleteCurrentChild")),
				Id("any").Op("=").Id("any").Dot("engine").Dot("create"+Title(anyNameByField(field))).Call(Id("any").Dot("ParentID"), Int().Call(Id(asw.v.Name+"ID")), Id("ElementKind"+Title(asw.v.Name)), Id("any").Dot("ParentElementPath"), Id("any").Dot("FieldIdentifier")).Dot(anyNameByField(field)),
				Switch(Id("any").Dot("FieldIdentifier")).Block(
					ForEachFieldInAST(s.config, func(_field ast.Field) *Statement {
						if _field.HasSliceValue || _field.HasPointerValue {
							return Empty()
						}
						if anyNameByField(field) != anyNameByField(_field) {
							return Empty()
						}
						return Case(Id(FieldPathIdentifier(_field))).Block(
							Id(field.Parent.Name).Op(":=").Id("any").Dot("engine").Dot(Title(field.Parent.Name)).Call(Id(Title(field.Parent.Name)+"ID").Call(Id("any").Dot("ParentID"))).Dot(field.Parent.Name),
							Id(field.Parent.Name).Dot(Title(field.Name)).Op("=").Id("any").Dot("ID"),
							Id(field.Parent.Name).Dot("engine").Dot("Patch").Dot("Item").Index(Id(field.Parent.Name).Dot("ID")).Op("=").Id(field.Parent.Name),
						)
					}),
				),
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

func (s *Factory) writeAnyRefs() *Factory {
	s.config.RangeAnyFields(func(field ast.Field) {

		a := anyRefWriter{
			f: field,
		}

		s.file.Type().Id(a.typeRefName()).Interface(
			Id("Kind").Params().Id("ElementKind").Line(),
			ForEachValueOfField(field, func(configType *ast.ConfigType) *Statement {
				return Id(Title(configType.Name)).Params().Id(Title(configType.Name)).Line()
			}),
		)

		s.file.Type().Id(Title(a.typeName())+"SliceElement").Interface(
			Id("Kind").Params().Id("ElementKind").Line(),
			ForEachValueOfField(field, func(configType *ast.ConfigType) *Statement {
				return Id(Title(configType.Name)).Params().Id(Title(configType.Name)).Line()
			}),
		)
	})

	return s
}
