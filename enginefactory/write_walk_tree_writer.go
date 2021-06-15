package enginefactory

import (
	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/Java-Jonas/bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

type walkElementWriter struct {
	t ast.ConfigType
	f *ast.Field
	v *ast.ConfigType
}

func (w walkElementWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (w walkElementWriter) name() string {
	return "walk" + Title(w.t.Name)
}

func (w walkElementWriter) idParam() string {
	return w.t.Name + "ID"
}

func (w walkElementWriter) params() *Statement {
	return List(Id(w.t.Name+"ID").Id(Title(w.t.Name)+"ID"), Id("p").Id("path"))
}

func (w walkElementWriter) dataElementName() string {
	return w.t.Name + "Data"
}

func (w walkElementWriter) getElementFromPatch() *Statement {
	return List(Id(w.dataElementName()), Id("hasUpdated")).Op(":=").Id("engine").Dot("Patch").Dot(Title(w.t.Name)).Index(Id(w.idParam()))
}

func (w walkElementWriter) getElementFromState() *Statement {
	return Id(w.dataElementName()).Op("=").Id("engine").Dot("State").Dot(Title(w.t.Name)).Index(Id(w.idParam()))
}

func (w walkElementWriter) declarePathVar() *Statement {
	return Var().Id(w.f.Name + "Path").Id("path")
}

func (w walkElementWriter) anyContainerName() string {
	return w.f.Name + "Container"
}

func (w walkElementWriter) usedChildIDIdentifier() *Statement {
	if w.f.HasAnyValue {
		return Id(w.anyContainerName()).Dot(Title(w.v.Name))
	}
	if w.f.HasSliceValue {
		return Id(w.v.Name + "ID")
	}
	return Id(w.dataElementName()).Dot(Title(w.f.Name))
}

func (w walkElementWriter) getChildPath() *Statement {
	return List(Id("existingPath"), Id("pathExists")).Op(":=").Id("engine").Dot("PathTrack").Dot(w.v.Name).Index(w.usedChildIDIdentifier())
}

func (w walkElementWriter) pathNeedsUpdate() *Statement {
	if w.f.HasSliceValue {
		return Id("!pathExists").Op("||").Id("!existingPath").Dot("equals").Call(Id("p"))
	}
	return Id("!pathExists")
}

func (w walkElementWriter) setChildPathNew() *Statement {
	statement := Id(w.f.Name + "Path").Op("=").Id("p").Dot(w.f.Name).Call()
	if !w.f.HasSliceValue {
		return statement
	}
	return statement.Dot("index").Call(Id("i"))
}

func (w walkElementWriter) setChildPathExisting() *Statement {
	return Id(w.f.Name + "Path").Op("=").Id("existingPath")
}

func (w walkElementWriter) walkChild() *Statement {
	return Id("engine").Dot("walk"+Title(w.v.Name)).Call(w.usedChildIDIdentifier(), Id(w.f.Name+"Path"))
}

func (w walkElementWriter) childrenLoopConditions() *Statement {
	return List(Id("i"), Id(w.v.Name+"ID")).Op(":=").Range().Id("merge"+Title(w.v.Name)+"IDs").Call(
		Id("engine").Dot("State").Dot(Title(w.t.Name)).Index(Id(w.t.Name+"Data").Dot("ID")).Dot(Title(w.f.Name)),
		Id("engine").Dot("Patch").Dot(Title(w.t.Name)).Index(Id(w.t.Name+"Data").Dot("ID")).Dot(Title(w.f.Name)),
	)
}

func (w walkElementWriter) anyChildLoopConditions() *Statement {
	return List(Id("i"), Id("anyID")).Op(":=").Range().Id(w.dataElementName()).Dot(Title(w.f.Name))
}

func (w walkElementWriter) declareAnyContainer() *Statement {
	idName := Id("anyID")
	if !w.f.HasSliceValue {
		idName = Id(w.dataElementName()).Dot(Title(w.f.Name))
	}
	return Id(w.anyContainerName()).Op(":=").Id("engine").Dot(w.f.ValueTypeName).Call(idName).Dot(w.f.ValueTypeName)
}

func (w walkElementWriter) updatePath() *Statement {
	return Id("engine").Dot("PathTrack").Dot(w.t.Name).Index(Id(w.idParam())).Op("=").Id("p")
}

type walkTreeWriter struct {
	t *ast.ConfigType
}

func (w walkTreeWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (w walkTreeWriter) dataElement() string {
	return w.t.Name + "Data"
}

func (w walkTreeWriter) checkWalked() *Statement {
	return Id("walkedCheck").Dot(w.t.Name).Index(Id(w.t.Name + "Data").Dot("ID")).Op("=").True()
}

func (w walkTreeWriter) walkElement() *Statement {
	return Id("engine").Dot("walk"+Title(w.t.Name)).Call(Id(w.t.Name+"Data").Dot("ID"), Id("newPath").Call(Id(w.t.Name+"Identifier"), Int().Call(Id("id"))))
}

func (w walkTreeWriter) patchLoopConditions() *Statement {
	return List(Id("id"), Id(w.t.Name+"Data")).Op(":=").Range().Id("engine").Dot("Patch").Dot(Title(w.t.Name))
}

func (w walkTreeWriter) elementDoesNotHaveParent() *Statement {
	return Id("!" + w.t.Name + "Data").Dot("HasParent")
}

func (w walkTreeWriter) stateLoopConditions() *Statement {
	return List(Id("id"), Id(w.t.Name+"Data")).Op(":=").Range().Id("engine").Dot("State").Dot(Title(w.t.Name))
}

func (w walkTreeWriter) hasNotBeenWalked() (*Statement, *Statement) {
	return List(Id("_"), Id("ok")).Op(":=").Id("walkedCheck").Dot(w.t.Name).Index(Id(w.t.Name + "Data").Dot("ID")), Id("!ok")
}

func (w walkTreeWriter) clearPathTrack() *Statement {
	return For(Id("key").Op(":=").Range().Id("engine").Dot("PathTrack").Dot(w.t.Name)).Block(
		Delete(Id("engine").Dot("PathTrack").Dot(w.t.Name), Id("key")),
	)
}
