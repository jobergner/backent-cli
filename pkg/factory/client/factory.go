package client

import (
	"bytes"

	"github.com/dave/jennifer/jen"

	"github.com/jobergner/backent-cli/pkg/ast"
	"github.com/jobergner/backent-cli/pkg/factory/configs"
	"github.com/jobergner/backent-cli/pkg/factory/utils"
)

func newSimpleASTExample() *ast.AST {
	simpleAST := ast.Parse(configs.StateConfig, configs.ActionsConfig, configs.ResponsesConfig)
	return simpleAST
}

type Factory struct {
	config *ast.AST
	file   *jen.File
}

func NewFactory(config *ast.AST) *Factory {
	return &Factory{
		config: config,
		file:   jen.NewFile(utils.PackageName),
	}
}

// Write writes source code for a given ActionsConfig
func (f *Factory) Write() string {
	f.writeActions().writeController()

	buf := bytes.NewBuffer(nil)
	f.file.Render(buf)

	return utils.TrimPackageClause(buf.String())
}
