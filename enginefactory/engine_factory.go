package enginefactory

import (
	"bytes"

	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/Java-Jonas/bar-cli/factoryutils"
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
	buf    *bytes.Buffer
}

// WriteEngine writes source code for a given StateConfig
func WriteEngine(buf *bytes.Buffer, stateConfigData map[interface{}]interface{}) {
	config := ast.Parse(stateConfigData, map[interface{}]interface{}{}, map[interface{}]interface{}{})
	s := newStateFactory(config).
		writePackageName(). // to be able to format the code without errors
		writeAdders().
		writeAny().
		writeAnyRefs().
		writeAssembleTree().
		writeAssembleTreeElement().
		writeAssembleTreeReference().
		writeCreators().
		writeDeleters().
		writeGetters().
		writeDeduplicate().
		writeAllIDsMethod().
		writeMergeIDs().
		writeIdentifiers().
		writePathSegments().
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
		writeRecursionCheck().
		writeAssembleCache().
		writePools()

	err := Format(s.buf)
	if err != nil {
		// unexpected error
		panic(err)
	}

	buf.WriteString(TrimPackageName(s.buf.String()))
}

func (s *EngineFactory) writePackageName() *EngineFactory {
	s.buf.WriteString("package state\n")
	return s
}

func newStateFactory(config *ast.AST) *EngineFactory {
	return &EngineFactory{
		config: config,
		buf:    &bytes.Buffer{},
	}
}
