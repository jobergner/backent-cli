package enginefactory

import (
	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/Java-Jonas/bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeAny() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeAnyFields(func(field ast.Field) {

		k := anyKindWriter{
			f: field,
		}

		decls.File.Func().Params(k.receiverParams()).Id("Kind").Params().Id("ElementKind").Block(
			k.reassignAnyContainer(),
			Return(k.containedElementKind()),
		)

		field.RangeValueTypes(func(valueType *ast.ConfigType) {
			s := anySetterWriter{
				f: field,
				v: *valueType,
			}
			decls.File.Func().Params(s.wrapperReceiverParams()).Id("Set"+Title(valueType.Name)).Params().Id(valueType.Name).Block(
				s.createChild(),
				s.callSetter(),
				Return(Id(valueType.Name)),
			)
			decls.File.Func().Params(s.receiverParams()).Id("set"+Title(valueType.Name)).Params(s.params()).Block(
				s.reassignAnyContainer(),
				ForEachValueOfField(field, func(_valueType *ast.ConfigType) *Statement {
					if _valueType.Name == valueType.Name {
						return Empty()
					}
					s._v = _valueType
					return If(s.otherValueIsSet()).Block(
						s.deleteOtherValue(),
						s.unsetIDInContainer(),
					)
				}),
				s.setElementKind(),
				s.setChildID(),
				s.updateContainerInPatch(),
			)
		})

		d := anyDeleteChildWriter{
			f: field,
		}
		decls.File.Func().Params(d.receiverParams()).Id("deleteChild").Params().Block(
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

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeAnyRefs() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeAnyFields(func(field ast.Field) {
	})

	decls.Render(s.buf)
	return s
}
