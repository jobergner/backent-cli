package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

type assemblePlannerWriter struct {
	t ast.ConfigType
	f *ast.Field
	v *ast.ConfigType
}

func (a assemblePlannerWriter) eachRefInState(source string) *Statement {
	return List(Id("_"), Id(a.f.ValueTypeName)).Op(":=").Range().Id(source).Dot(Title(a.f.ValueTypeName))
}

func (a assemblePlannerWriter) pathAlreadyIncluded() *Statement {
	return List(Id("_"), Id("ok")).Op(":=").Id("ap").Dot("updatedReferencePaths").Index(Id("ComplexID").Call(Id(a.f.ValueTypeName).Dot("ID")))
}

func (a assemblePlannerWriter) checkedElementID() *Statement {
	return Id(a.f.ValueTypeName).Dot("ID").Dot("ChildID")
}

func (a assemblePlannerWriter) includedElementsContainReferencedElement() *Statement {
	return List(Id("_"), Id("ok")).Op(":=").Id("ap").Dot("includedElements").Index(Int().Call(a.checkedElementID()))
}

func (a assemblePlannerWriter) putPathInUpdatedReferencePaths() *Statement {
	return Id("ap").Dot("updatedReferencePaths").Index(Id("ComplexID").Call(Id(a.f.ValueTypeName).Dot("ID"))).Op("=").Id(a.f.ValueTypeName).Dot("Path")
}
