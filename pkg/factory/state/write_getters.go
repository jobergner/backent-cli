package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeGetters() *Factory {

	RangeBasicTypes(func(b BasicType) {
		s.file.Func().Params(Id("engine").Id("*Engine")).Id(b.Value).Params(Id(b.Value+"ID").Id(Title(b.Value)+"ID")).Id(b.Value).Block(
			List(Id("patching"+Title(b.Value)), Id("ok")).Op(":=").Id("engine").Dot("Patch").Dot(Title(b.Value)).Index(Id(b.Value+"ID")),
			If(Id("ok")).Block(
				Return(Id("patching"+Title(b.Value))),
			),
			Return(Id("engine").Dot("State").Dot(Title(b.Value)).Index(Id(b.Value+"ID"))),
		)
	})

	s.config.RangeTypes(func(configType ast.ConfigType) {

		ex := existsGetterWriter{
			t: configType,
		}

		writeExistsGetter(s.file, ex)

		e := everyTypeGetterWriter{
			t: configType,
		}

		writeQueryTypeGetter(s.file, e)
		writeEveryTypeGetter(s.file, e)

		t := typeGetterWriter{
			name: func() string {
				return Title(configType.Name)
			},
			typeName: func() string {
				return configType.Name
			},
		}

		if !configType.IsRootType {
			p := parentKindGetter{
				t: configType,
			}

			s.file.Func().Params(p.receiverParams()).Id("ParentKind").Params().Params(Id("ElementKind"), Bool()).Block(
				If(p.hasNoParent()).Block(
					Return(Lit(""), False()),
				),
				Return(p.parentKind(), True()),
			)
		}

		writeTypeGetter(s.file, t)

		configType.RangeImplementedBy(func(parentType *ast.ConfigType) {

			p := parentGetter{
				t:      configType,
				parent: *parentType,
			}

			s.file.Func().Params(p.receiverParams()).Id("Parent"+Title(p.parent.Name)).Params().Id(Title(p.parent.Name)).Block(
				p.reassignElement(),
				If(p.hasNoParent()).Block(
					Return(p.emptyReturn()),
				),
				p.declareParentSeg(),
				Return(p.getParent()),
			)
		})

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

		writeIDGetter(s.file, i)

		p := pathGetterWriter{
			t: configType,
		}

		s.file.Func().Params(p.receiverParams()).Id("Path").Params().String().Block(
			Return(p.returnPath()),
		)

		configType.RangeFields(func(field ast.Field) {

			f := fieldGetterWriter{
				t: configType,
				f: field,
			}

			s.file.Func().Params(f.receiverParams()).Id(f.name()).Params().Id(f.returns()).Block(
				f.reassignElement(),
				// if slice
				OnlyIf(field.HasSliceValue, f.declareSliceOfElements()),
				OnlyIf(field.HasSliceValue, Id("sort").Dot("Slice").Call(f.element().Dot(Title(f.f.Name)), Func().Params(Id("i"), Id("j").Int()).Bool().Block(
					Return(f.element().Dot(Title(f.f.Name)).Index(Id("i")).Op("<").Add(f.element()).Dot(Title(f.f.Name)).Index(Id("j"))),
				))),
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

		writeTypeGetter(s.file, t)

		writeIDGetter(s.file, i)

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

		writeTypeGetter(s.file, t)

		writeIDGetter(s.file, i)

		field.RangeValueTypes(func(valueType *ast.ConfigType) {
			s.file.Func().Params(Id("_"+t.name()).Id(Title(t.name()))).Id(Title(valueType.Name)).Params().Id(Title(valueType.Name)).Block(
				Id(t.name()).Op(":=").Id("_"+t.name()).Dot(t.name()).Dot("engine").Dot(t.name()).Call(Id("_"+t.name()).Dot(t.name()).Dot("ID")),
				Return(Id(t.name()).Dot(t.name()).Dot("engine").Dot(Title(valueType.Name)).Call(Id(Title(valueType.Name+"ID")).Call(Id(t.name()).Dot(t.name()).Dot("ChildID")))),
			)
		})
	})

	return s
}

func writeEveryTypeGetter(file *File, e everyTypeGetterWriter) {
	file.Func().Params(e.receiverParams()).Id("Every"+Title(e.t.Name)).Params().Add(e.returns()).Block(
		e.allIDs(),
		e.sortIDs(),
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

func writeQueryTypeGetter(file *File, e everyTypeGetterWriter) {
	file.Func().Params(e.receiverParams()).Id("Query"+Plural(Title(e.t.Name))).Params(Id("matcher").Func().Params(Id(Title(e.t.Name))).Bool()).Add(e.returns()).Block(
		e.allIDs(),
		e.sortIDs(),
		e.declareSlice(),
		For(e.loopConditions()).Block(
			e.declareElement(),
			If(e.elementSatisfiesMatcher()).Block(
				e.appendElement(),
			),
		),
		e.returnToIdsSliceToPool(),
		Return(Id(e.sliceName())),
	)
}

func writeTypeGetter(file *File, t typeGetterWriter) {
	file.Func().Params(t.receiverParams()).Id(t.name()).Params(t.params()).Id(t.returns()).Block(
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

func writeIDGetter(file *File, i idGetterWriter) {
	file.Func().Params(i.receiverParams()).Id(i.name()).Params().Id(i.returns()).Block(
		Return(i.returnID()),
	)
}

func writeExistsGetter(file *File, e existsGetterWriter) {
	file.Func().Params(e.receiverParams()).Id("Exists").Params().Params(e.returnTypes()).Block(
		e.reassignElement(),
		Return(Id(e.t.Name), e.isNotOperationKindDelete()),
	)
}
