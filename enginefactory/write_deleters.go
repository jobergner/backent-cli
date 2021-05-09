package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeDeleters() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {

		tw := typeDeleterWrapper{
			t: configType,
		}

		decls.File.Func().Params(tw.receiverParams()).Id(tw.name()).Params(tw.params()).Block(
			onlyIf(!configType.IsRootType, tw.getElement()),
			onlyIf(!configType.IsRootType, If(tw.hasParent()).Block(
				Return(),
			)),
			tw.deleteElement(),
		)

		t := typeDeleter{
			t: configType,
			f: nil,
		}

		decls.File.Func().Params(t.receiverParams()).Id(t.name()).Params(t.params()).Block(
			t.getElement(),
			t.setOperationKind(),
			t.updateElementInPatch(),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				t.f = &field
				if field.ValueType().IsBasicType {
					return Empty()
				}
				if !field.HasSliceValue {
					return t.deleteElement()
				}
				return For(t.loopConditions().Block(
					t.deleteElementInLoop(),
				))
			}),
		)
	})

	decls.Render(s.buf)
	return s
}

type typeDeleterWrapper struct {
	t ast.ConfigType
}

func (tw typeDeleterWrapper) receiverParams() *Statement {
	return Id("se").Id("*Engine")
}

func (tw typeDeleterWrapper) name() string {
	return "Delete" + title(tw.t.Name)
}

func (tw typeDeleterWrapper) idParam() string {
	return tw.t.Name + "ID"
}

func (tw typeDeleterWrapper) params() *Statement {
	return Id(tw.idParam()).Id(title(tw.t.Name) + "ID")
}

func (tw typeDeleterWrapper) getElement() *Statement {
	return Id(tw.t.Name).Op(":=").Id("se").Dot(title(tw.t.Name)).Call(Id(tw.idParam())).Dot(tw.t.Name)
}

func (tw typeDeleterWrapper) hasParent() *Statement {
	return Id(tw.t.Name).Dot("HasParent_")
}

func (tw typeDeleterWrapper) deleteElement() *Statement {
	return Id("se").Dot("delete" + title(tw.t.Name)).Call(Id(tw.idParam()))
}

type typeDeleter struct {
	t ast.ConfigType
	f *ast.Field
}

func (t typeDeleter) receiverParams() *Statement {
	return Id("se").Id("*Engine")
}

func (t typeDeleter) name() string {
	return "delete" + title(t.t.Name)
}

func (t typeDeleter) idParam() string {
	return t.t.Name + "ID"
}

func (t typeDeleter) params() *Statement {
	return Id(t.idParam()).Id(title(t.t.Name) + "ID")
}

func (t typeDeleter) getElement() *Statement {
	return Id(t.t.Name).Op(":=").Id("se").Dot(title(t.t.Name)).Call(Id(t.idParam())).Dot(t.t.Name)
}

func (t typeDeleter) setOperationKind() *Statement {
	return Id(t.t.Name).Dot("OperationKind_").Op("=").Id("OperationKindDelete")
}

func (t typeDeleter) updateElementInPatch() *Statement {
	return Id("se").Dot("Patch").Dot(title(t.t.Name)).Index(Id(t.t.Name).Dot("ID")).Op("=").Id(t.t.Name)
}

func (t typeDeleter) loopConditions() *Statement {
	return List(Id("_"), Id(t.f.ValueType().Name+"ID")).Op(":=").Range().Id(t.t.Name).Dot(title(t.f.Name))
}

func (t typeDeleter) deleteElementInLoop() *Statement {
	return Id("se").Dot("delete" + title(t.f.ValueType().Name)).Call(Id(t.f.ValueType().Name + "ID"))
}

func (t typeDeleter) deleteElement() *Statement {
	return Id("se").Dot("delete" + title(t.f.ValueType().Name)).Call(Id(t.t.Name).Dot(title(t.f.ValueType().Name)))
}
