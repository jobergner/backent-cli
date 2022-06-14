package server

import (
	"bytes"

	"github.com/dave/jennifer/jen"

	"github.com/jobergner/backent-cli/pkg/ast"
	"github.com/jobergner/backent-cli/pkg/factory/utils"
)

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

func (f *Factory) Write() string {
	f.writeTriggerAction().
		writeController()

	buf := bytes.NewBuffer(nil)
	f.file.Render(buf)

	return utils.TrimPackageClause(buf.String())
}
