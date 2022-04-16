package state

import (
	"bytes"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/ast"
	"github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
)

func anyNameByField(f ast.Field) string {
	name := "anyOf"
	firstIteration := true
	f.RangeValueTypes(func(configType *ast.ConfigType) {
		if firstIteration {
			name += Title(configType.Name)
		} else {
			name += "_" + Title(configType.Name)
		}
		firstIteration = false
	})
	return name
}

type Factory struct {
	config *ast.AST
	file   *jen.File
}

// Write writes source code for a given StateConfig
func (f *Factory) Write() string {
	f.writeAdders().
		writeAny().
		writeAnyRefs().
		writeAssemblePlanner().
		writeAssemblePlannerClear().
		writeAssemblePlannerPlan().
		writeAssemblePlannerFill().
		writeAssembleBranch().
		writeAssembleTree().
		writeCreators().
		writeDeleters().
		writeGetters().
		writeDeduplicate().
		writeAllIDsMethod().
		writeIdentifiers().
		writePath().
		writeReference().
		writeDereference().
		writeRemovers().
		writeSetters().
		writeComplexID().
		writeIDs().
		writeState().
		writeMetaData().
		writeElements().
		writeOperationKind().
		writeEngine().
		writeGenerateID().
		writeUpdateState().
		writeImportPatch().
		writeReferencedDataStatus().
		writeElementKinds().
		writeTree().
		writeTreeElements().
		writePools()

	buf := bytes.NewBuffer(nil)
	f.file.Render(buf)

	return utils.TrimPackageClause(buf.String())
}

func NewFactory(config *ast.AST) *Factory {
	return &Factory{
		config: config,
		file:   jen.NewFile(utils.PackageName),
	}
}

func defaultValueForBasicType(typeLiteral string) interface{} {
	switch typeLiteral {
	case "bool":
		return bool(false)
	case "string":
		return string("")
	case "int8":
		return int8(0)
	case "byte":
		return byte(0)
	case "int16":
		return int16(0)
	case "uint16":
		return uint16(0)
	case "rune":
		return rune(0)
	case "uint32":
		return uint32(0)
	case "int64":
		return 0
	case "uint64":
		return uint64(0)
	case "int":
		return int(0)
	case "uint":
		return uint(0)
	case "uintptr":
		return uintptr(0)
	case "float32":
		return float32(0)
	case "float64":
		return float64(0)
	case "complex64":
		return complex64(0)
	case "complex128":
		return complex128(0)
	}

	return 0
}
