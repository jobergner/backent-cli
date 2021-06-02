package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeWalkElement() *EngineFactory {
	decls := NewDeclSet()

	s.config.RangeTypes(func(configType ast.ConfigType) {

		w := walkElementWriter{
			t: configType,
		}

		if configType.IsLeafType {
			decls.File.Func().Params(w.receiverParams()).Id(w.name()).Params(w.params()).Block(
				w.updatePath(),
			)
			return
		}

		decls.File.Func().Params(w.receiverParams()).Id(w.name()).Params(w.params()).Block(
			w.getElementFromPatch(),
			If(Id("!hasUpdated")).Block(
				w.getElementFromState(),
			),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				if field.ValueType().IsBasicType || field.HasPointerValue {
					return Empty()
				}

				w.f = &field
				w.v = field.ValueType()

				if !field.HasAnyValue {
					if !field.HasSliceValue {
						return writeEvaluateChildPath(w)
					}
					return For(w.childrenLoopConditions()).Block(
						writeEvaluateChildPath(w),
					)
				}

				if !field.HasSliceValue {
					return &Statement{
						w.declareAnyContainer().Line(),
						forEachFieldValueComparison(field, *Id(field.Name + "Container").Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
							w.v = valueType
							return writeEvaluateChildPath(w)

						}),
					}
				}
				return For(w.anyChildLoopConditions()).Block(
					w.declareAnyContainer(),
					forEachFieldValueComparison(field, *Id(field.Name + "Container").Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
						w.v = valueType
						return writeEvaluateChildPath(w)

					}),
				)
			}),
			w.updatePath(),
		)
	})

	decls.Render(s.buf)
	return s
}

func writeEvaluateChildPath(w walkElementWriter) *Statement {
	return &Statement{
		w.declarePathVar().Line(),
		If(w.getChildPath(), w.pathNeedsUpdate()).Block(
			w.setChildPathNew(),
		).Else().Block(
			w.setChildPathExisting(),
		).Line(),
		w.walkChild(),
	}
}
