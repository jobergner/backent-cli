package getstartedfactory

import (
	. "bar-cli/factoryutils"
	"bytes"

	"bar-cli/ast"
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

func WriteGetStarted(stateConfigData, actionsConfigData map[interface{}]interface{}) string {
	config := ast.Parse(stateConfigData, actionsConfigData)
	g := newGetStartedFactory(config).
		writePackageName().
		writeImport().
		writeMainFunc()

	// err := Format(g.buf)
	// if err != nil {
	// 	// unexpected error
	// 	panic(err)
	// }

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
)
	`)
	return g
}

func (g *GetStartedFactory) writeMainFunc() *GetStartedFactory {
	decls := NewDeclSet()

	decls.File.Func().Id("main").Params().Block(
		Id("state").Dot("Start").Call(
			ForEachActionInAST(g.config, func(action ast.Action) *Statement {
				return Line().Func().Params(Id("state").Dot(Title(action.Name)+"Params"), Id("*state").Dot("Engine")).Block().Id(",")
			}).
				Comment("onDeploy").Line().
				Func().Params(Id("*state").Dot("Engine")).Block().Id(",").Line().
				Comment("onFrameTick").Line().
				Func().Params(Id("*state").Dot("Engine")).Block().Id(",").Line(),
		),
	)

	decls.Render(g.buf)
	return g
}
