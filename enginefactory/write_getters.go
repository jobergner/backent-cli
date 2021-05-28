package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeGetters() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		t := typeGetterWriter{
			name: func() string {
				return title(configType.Name)
			},
			typeName: func() string {
				return configType.Name
			},
		}

		i := idGetterWriter{
			typeName: func() string {
				return configType.Name
			},
			returns: func() string {
				return title(configType.Name) + "ID"
			},
			idFieldToReturn: func() string {
				return "ID"
			},
		}

		writeTypeGetter(&decls, t)
		writeIDGetter(&decls, i)

		configType.RangeFields(func(field ast.Field) {
			f := fieldGetter{
				t: configType,
				f: field,
			}

			decls.File.Func().Params(f.receiverParams()).Id(f.name()).Params().Id(f.returns()).Block(
				f.reassignElement(),
				// if slice
				onlyIf(field.HasSliceValue, f.declareSliceOfElements()),
				onlyIf(field.HasSliceValue, For(f.loopConditions().Block(
					f.appendElement(),
				))),
				onlyIf(field.HasSliceValue, Return(f.returnSliceOfType())),
				// if not slice
				onlyIf(!field.HasSliceValue, Return(f.returnSingleType())),
			)
		})

	})

	s.config.RangeRefFields(func(field ast.Field) {

		t := typeGetterWriter{
			name: func() string {
				return field.ValueTypeName
			},
		}
		i := idGetterWriter{
			idFieldToReturn: func() string {
				return "ReferencedElementID"
			},
		}

		if field.HasAnyValue {
			i.returns = func() string {
				return title(anyNameByField(field)) + "ID"
			}
		} else {
			i.returns = func() string {
				return title(field.ValueType().Name) + "ID"
			}
		}

		t.typeName = t.name
		i.typeName = t.name

		writeTypeGetter(&decls, t)
		writeIDGetter(&decls, i)
	})

	s.config.RangeAnyFields(func(field ast.Field) {
		t := typeGetterWriter{
			name: func() string {
				return anyNameByField(field)
			},
		}
		i := idGetterWriter{
			idFieldToReturn: func() string {
				return "ID"
			},
			returns: func() string {
				return title(t.name()) + "ID"
			},
		}

		t.typeName = t.name
		i.typeName = t.name

		writeTypeGetter(&decls, t)
		writeIDGetter(&decls, i)
	})

	decls.Render(s.buf)
	return s
}

func writeTypeGetter(decls *DeclSet, t typeGetterWriter) {
	decls.File.Func().Params(t.receiverParams()).Id(t.name()).Params(t.params()).Id(t.returns()).Block(
		t.definePatchingElement(),
		If(Id("ok")).Block(
			Return(t.earlyReturnPatching()),
		),
		t.defineCurrentElement(),
		If(Id("ok")).Block(
			Return(t.earlyReturnCurrent()),
		),
		Return(t.finalReturn()),
	)
}

func writeIDGetter(decls *DeclSet, i idGetterWriter) {
	decls.File.Func().Params(i.receiverParams()).Id(i.name()).Params().Id(i.returns()).Block(
		Return(i.returnID()),
	)
}
