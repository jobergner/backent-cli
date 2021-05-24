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

type typeGetterWriter struct {
	name     func() string
	typeName func() string
}

func (t typeGetterWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (t typeGetterWriter) idParam() string {
	return t.typeName() + "ID"
}

func (t typeGetterWriter) params() *Statement {
	return Id(t.idParam()).Id(title(t.typeName()) + "ID")
}

func (t typeGetterWriter) returns() string {
	return t.typeName()
}

func (t typeGetterWriter) definePatchingElement() *Statement {
	return List(Id("patching"+title(t.typeName())), Id("ok")).Op(":=").Id("engine").Dot("Patch").Dot(title(t.typeName())).Index(Id(t.idParam()))
}

func (t typeGetterWriter) earlyReturnPatching() *Statement {
	return Id(t.typeName()).Values(Dict{Id(t.typeName()): Id("patching" + title(t.typeName()))})
}

func (t typeGetterWriter) defineCurrentElement() *Statement {
	return List(Id("current"+title(t.typeName())), Id("ok")).Op(":=").Id("engine").Dot("State").Dot(title(t.typeName())).Index(Id(t.idParam()))
}

func (t typeGetterWriter) earlyReturnCurrent() *Statement {
	return Id(t.typeName()).Values(Dict{Id(t.typeName()): Id("current" + title(t.typeName()))})
}

func (t typeGetterWriter) finalReturn() *Statement {
	return Id(t.typeName()).Values(Dict{Id(t.typeName()): Id(t.typeName() + "Core").Values(Dict{Id("OperationKind"): Id("OperationKindDelete"), Id("engine"): Id("engine")})})
}

type idGetterWriter struct {
	typeName        func() string
	returns         func() string
	idFieldToReturn func() string
}

func (i idGetterWriter) receiverName() string {
	return "_" + i.typeName()
}

func (i idGetterWriter) receiverParams() *Statement {
	return Id(i.receiverName()).Id(i.typeName())
}

func (i idGetterWriter) name() string {
	return "ID"
}

func (i idGetterWriter) returnID() *Statement {
	return Id(i.receiverName()).Dot(i.typeName()).Dot(i.idFieldToReturn())
}

type fieldGetter struct {
	t ast.ConfigType
	f ast.Field
}

func (f fieldGetter) receiverName() string {
	return "_" + f.t.Name
}

func (f fieldGetter) receiverParams() *Statement {
	return Id(f.receiverName()).Id(f.t.Name)
}

func (f fieldGetter) name() string {
	return title(f.f.Name)
}

func (f fieldGetter) returnedType() string {

	if f.f.ValueType().IsBasicType {
		return f.f.ValueType().Name
	}
	if f.f.HasPointerValue {
		return f.f.Parent.Name + title(pluralizeClient.Singular(f.f.Name)) + "Ref"
	}
	if f.f.HasAnyValue {
		return anyNameByField(f.f)
	}
	return f.f.ValueType().Name
}

func (f fieldGetter) returns() string {
	returnedLiteral := f.returnedType()
	if f.f.HasSliceValue {
		return "[]" + returnedLiteral
	} else if f.f.HasPointerValue {
		return "(" + returnedLiteral + ", bool)"
	}
	return returnedLiteral
}

func (f fieldGetter) reassignElement() *Statement {
	return Id(f.t.Name).Op(":=").Id(f.receiverName()).Dot(f.t.Name).Dot("engine").Dot(title(f.t.Name)).Call(Id(f.receiverName()).Dot(f.t.Name).Dot("ID"))
}

func (f fieldGetter) declareSliceOfElements() *Statement {
	returnedType := f.returnedType()
	if f.f.HasSliceValue {
		returnedType = "[]" + returnedType
	}
	return Var().Id(f.f.Name).Id(returnedType)
}

func (f fieldGetter) loopedElementIdentifier() string {
	if f.f.ValueType().IsBasicType {
		return "element"
	}
	if f.f.HasPointerValue {
		return "refID"
	}
	if f.f.HasAnyValue {
		return anyNameByField(f.f) + "ID"
	}
	return f.f.ValueType().Name + "ID"
}

func (f fieldGetter) loopConditions() *Statement {
	identifier := f.loopedElementIdentifier()
	return List(Id("_"), Id(identifier)).Op(":=").Range().Id(f.t.Name).Dot(f.t.Name).Dot(title(f.f.Name))
}

func (f fieldGetter) appendedItem() *Statement {
	if f.f.ValueType().IsBasicType {
		return Id(f.loopedElementIdentifier())
	}
	returnedType := f.returnedType()
	if !f.f.HasPointerValue && !f.f.HasAnyValue {
		returnedType = title(returnedType)
	}
	return Id(f.t.Name).Dot(f.t.Name).Dot("engine").Dot(returnedType).Call(Id(f.loopedElementIdentifier()))
}

func (f fieldGetter) appendElement() *Statement {
	return Id(f.f.Name).Op("=").Append(Id(f.f.Name), f.appendedItem())
}

func (f fieldGetter) returnSliceOfType() *Statement {
	return Id(f.f.Name)
}

func (f fieldGetter) returnBasicType() *Statement {
	return Id(f.t.Name).Dot(f.t.Name).Dot(title(f.f.Name))
}

func (f fieldGetter) returnNamedType() *Statement {
	engine := Id(f.t.Name).Dot(f.t.Name).Dot("engine")
	if f.f.HasPointerValue {
		return engine.Dot(f.returnedType()).Call(Id(f.t.Name).Dot(f.t.Name).Dot(title(f.f.Name))).Id(",").Id(f.t.Name).Dot(f.t.Name).Dot(title(f.f.Name)).Op("!=").Lit(0)
	}
	returnedType := f.returnedType()
	if !f.f.HasAnyValue {
		returnedType = title(returnedType)
	}
	return engine.Dot(returnedType).Call(Id(f.t.Name).Dot(f.t.Name).Dot(title(f.f.Name)))
}

func (f fieldGetter) returnSingleType() *Statement {
	if f.f.ValueType().IsBasicType {
		return f.returnBasicType()
	}
	return f.returnNamedType()
}
