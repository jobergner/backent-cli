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
					return writeEvaluateAnyChildPath(w, field)
				}
				return For(w.anyChildLoopConditions()).Block(
					writeEvaluateAnyChildPath(w, field),
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

func writeEvaluateAnyChildPath(w walkElementWriter, field ast.Field) *Statement {
	firstValueTypeIteration := true
	statement := w.declareAnyContainer().Line()
	field.RangeValueTypes(func(valueType *ast.ConfigType) {
		w.v = valueType
		if !firstValueTypeIteration {
			statement = statement.Else()
		}
		firstValueTypeIteration = false

		statement.If(Id(field.Name + "Container").Dot("ElementKind").Op("==").Id("ElementKind" + title(valueType.Name))).Block(
			writeEvaluateChildPath(w),
		)

	})

	return statement
}
