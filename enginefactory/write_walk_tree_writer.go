package enginefactory

import (
	"bar-cli/ast"

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
	return "walk" + title(w.t.Name)
}

func (w walkElementWriter) idParam() string {
	return w.t.Name + "ID"
}

func (w walkElementWriter) params() *Statement {
	return List(Id(w.t.Name+"ID").Id(title(w.t.Name)+"ID"), Id("p").Id("path"))
}

func (w walkElementWriter) dataElementName() string {
	return w.t.Name + "Data"
}

func (w walkElementWriter) getElementFromPatch() *Statement {
	return List(Id(w.dataElementName()), Id("hasUpdated")).Op(":=").Id("engine").Dot("Patch").Dot(title(w.t.Name)).Index(Id(w.idParam()))
}

func (w walkElementWriter) getElementFromState() *Statement {
	return Id(w.dataElementName()).Op("=").Id("engine").Dot("State").Dot(title(w.t.Name)).Index(Id(w.idParam()))
}

func (w walkElementWriter) declarePathVar() *Statement {
	return Var().Id(w.f.Name + "Path").Id("path")
}

func (w walkElementWriter) anyContainerName() string {
	return w.f.Name + "Container"
}

func (w walkElementWriter) usedChildIDIdentifier() *Statement {
	if w.f.HasAnyValue {
		return Id(w.anyContainerName()).Dot(title(w.v.Name))
	}
	if w.f.HasSliceValue {
		return Id(w.v.Name + "ID")
	}
	return Id(w.dataElementName()).Dot(title(w.f.Name))
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
	return Id("engine").Dot("walk"+title(w.v.Name)).Call(w.usedChildIDIdentifier(), Id(w.f.Name+"Path"))
}

func (w walkElementWriter) childrenLoopConditions() *Statement {
	return List(Id("i"), Id(w.v.Name+"ID")).Op(":=").Range().Id("merge"+title(w.v.Name)+"IDs").Call(
		Id("engine").Dot("State").Dot(title(w.t.Name)).Index(Id(w.t.Name+"Data").Dot("ID")).Dot(title(w.f.Name)),
		Id("engine").Dot("Patch").Dot(title(w.t.Name)).Index(Id(w.t.Name+"Data").Dot("ID")).Dot(title(w.f.Name)),
	)
}

func (w walkElementWriter) anyChildLoopConditions() *Statement {
	return List(Id("i"), Id("anyID")).Op(":=").Range().Id(w.dataElementName()).Dot(title(w.f.Name))
}

func (w walkElementWriter) declareAnyContainer() *Statement {
	idName := Id("anyID")
	if !w.f.HasSliceValue {
		idName = Id(w.dataElementName()).Dot(title(w.f.Name))
	}
	return Id(w.anyContainerName()).Op(":=").Id("engine").Dot(w.f.ValueTypeName).Call(idName).Dot(w.f.ValueTypeName)
}

func (w walkElementWriter) updatePath() *Statement {
	return Id("engine").Dot("PathTrack").Dot(w.t.Name).Index(Id(w.idParam())).Op("=").Id("p")
}
