package getstartedfactory

import (
	"bytes"

	. "github.com/Java-Jonas/bar-cli/factoryutils"

	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/dave/jennifer/jen"
)

type GetStartedFactory struct {
	config *ast.AST
	buf    *bytes.Buffer
}

func newGetStartedFactory(config *ast.AST) *GetStartedFactory {
	return &GetStartedFactory{
		config: config,
		buf:    &bytes.Buffer{},
	}
}

func WriteGetStarted(stateConfigData, actionsConfigData, responsesConfigData map[interface{}]interface{}) string {
	config := ast.Parse(stateConfigData, actionsConfigData, responsesConfigData)
	g := newGetStartedFactory(config).
		writePackageName().
		writeImport().
		writeMainFunc()

	return g.buf.String()
}

func (g *GetStartedFactory) writePackageName() *GetStartedFactory {
	g.buf.WriteString("package main\n")
	return g
}

// TODO fix
func (g *GetStartedFactory) writeImport() *GetStartedFactory {
	g.buf.WriteString(`
import (
	"foo/tmp"
)`)
	return g
}

func (g *GetStartedFactory) writeMainFunc() *GetStartedFactory {
	decls := NewDeclSet()

	decls.File.Const().Id("fps").Op("=").Lit(30)

	decls.File.Var().Id("sideEffects").Op("=").Id("state").Dot("SideEffects").Values(Dict{
		Id("OnDeploy"):    Func().Params(Id("engine").Id("*state.Engine")).Block(),
		Id("OnFrameTick"): Func().Params(Id("engine").Id("*state.Engine")).Block(),
	}).Line()

	decls.File.Var().Id("actions").Op("=").Id("state").Dot("Actions").Values(
		Line().Add(
			ForEachActionInAST(g.config, func(action ast.Action) *Statement {
				if action.Response == nil {
					return Id(Title(action.Name)).Op(":").Func().Params(Id("params").Id("state").Dot(Title(action.Name)+"Params"), Id("engine").Id("*state.Engine")).Block().Id(",")
				}
				responseName := Id("state").Dot(Title(action.Name) + "Response")
				return Id(Title(action.Name)).Op(":").Func().Params(Id("params").Id("state").Dot(Title(action.Name)+"Params"), Id("engine").Id("*state.Engine")).Add(responseName).Block(
					Return(responseName).Values(),
				).Id(",")
			}),
		),
	)

	decls.File.Func().Id("main").Params().Block(
		Id("err").Op(":=").Id("state").Dot("Start").Call(Id("actions"), Id("sideEffects"), Id("fps"), Lit(8080)),
		If(Id("err").Op("!=").Nil()).Block(
			Panic(Id("err")),
		),
	)

	decls.Render(g.buf)
	return g
}
