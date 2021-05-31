package enginefactory

import (
	"bar-cli/ast"

	. "github.com/dave/jennifer/jen"
)

type pathTrackWriter struct {
	t ast.ConfigType
}

func (p pathTrackWriter) a() *Statement {
	return Id("")
}
