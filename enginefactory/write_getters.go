package enginefactory

import (
	"github.com/jobergner/backent-cli/ast"
	. "github.com/jobergner/backent-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeGetters() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {

		e := everyTypeGetterWriter{
			t: configType,
		}
		writeEveryTypeGetter(&decls, e)

		t := typeGetterWriter{
			name: func() string {
				return Title(configType.Name)
			},
			typeName: func() string {
				return configType.Name
			},
		}
		writeTypeGetter(&decls, t)

		i := idGetterWriter{
			typeName: func() string {
				return configType.Name
			},
			returns: func() string {
				return Title(configType.Name) + "ID"
			},
			idFieldToReturn: func() string {
				return "ID"
			},
		}
		writeIDGetter(&decls, i)

		p := pathGetterWriter{
			t: configType,
		}
		decls.File.Func().Params(p.receiverParams()).Id("Path").Params().String().Block(
			Return(p.returnPath()),
		)

		configType.RangeFields(func(field ast.Field) {
			f := fieldGetterWriter{
				t: configType,
				f: field,
			}
			decls.File.Func().Params(f.receiverParams()).Id(f.name()).Params().Id(f.returns()).Block(
				f.reassignElement(),
				// if slice
				OnlyIf(field.HasSliceValue, f.declareSliceOfElements()),
				OnlyIf(field.HasSliceValue, For(f.loopConditions().Block(
					f.appendElement(),
				))),
				OnlyIf(field.HasSliceValue, Return(f.returnSliceOfType())),
				// if not slice
				OnlyIf(!field.HasSliceValue, Return(f.returnSingleType())),
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
				return Title(anyNameByField(field)) + "ID"
			}
		} else {
			i.returns = func() string {
				return Title(field.ValueType().Name) + "ID"
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
				return Title(t.name()) + "ID"
			},
		}

		t.typeName = t.name
		i.typeName = t.name

		writeTypeGetter(&decls, t)
		writeIDGetter(&decls, i)

		field.RangeValueTypes(func(valueType *ast.ConfigType) {
			decls.File.Func().Params(Id("_"+t.name()).Id(t.name())).Id(Title(valueType.Name)).Params().Id(valueType.Name).Block(
				Id(t.name()).Op(":=").Id("_"+t.name()).Dot(t.name()).Dot("engine").Dot(t.name()).Call(Id("_"+t.name()).Dot(t.name()).Dot("ID")),
				Return(Id(t.name()).Dot(t.name()).Dot("engine").Dot(Title(valueType.Name)).Call(Id(t.name()).Dot(t.name()).Dot(Title(valueType.Name)))),
			)
		})
	})

	decls.Render(s.buf)
	return s
}

func writeEveryTypeGetter(decls *DeclSet, e everyTypeGetterWriter) {
	decls.File.Func().Params(e.receiverParams()).Id("Every"+Title(e.t.Name)).Params().Add(e.returns()).Block(
		e.allIDs(),
		e.declareSlice(),
		For(e.loopConditions()).Block(
			e.declareElement(),
			OnlyIf(!e.t.IsRootType, If(e.elementHasParent()).Block(
				Continue(),
			)),
			e.appendElement(),
		),
		e.returnToIdsSliceToPool(),
		Return(Id(e.sliceName())),
	)
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
