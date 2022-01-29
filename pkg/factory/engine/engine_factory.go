package engine

import (
	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"
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

type EngineFactory struct {
	config *ast.AST
	file   *jen.File
}

// WriteEngine writes source code for a given StateConfig
func WriteEngine(file *jen.File, stateConfigData map[interface{}]interface{}) {

	config := ast.Parse(
		stateConfigData,
		map[interface{}]interface{}{},
		map[interface{}]interface{}{},
	)

	newStateFactory(file, config).
		writeAdders().
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
		writeIDs().
		writeState().
		writeElements().
		writeOperationKind().
		writeEngine().
		writeGenerateID().
		writeUpdateState().
		writeReferencedDataStatus().
		writeElementKinds().
		writeTree().
		writeTreeElements().
		writePools()
}

func newStateFactory(file *jen.File, config *ast.AST) *EngineFactory {
	return &EngineFactory{
		config: config,
		file:   file,
	}
}
