package enginefactory

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"

	"bar-cli/ast"
	. "bar-cli/factoryutils"
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
	config := ast.Parse(stateConfigData, map[interface{}]interface{}{})
	s := newStateFactory(config).
		writePackageName(). // to be able to format the code without errors
		writeImports().
		writeAdders().
		writeAny().
		writeAssembleTree().
		writeAssembleTreeElement().
		writeAssembleTreeReference().
		writeCreators().
		writeDeleters().
		writeGetters().
		writeDeduplicate().
		writeAllIDsMethod().
		writeMergeIDs().
		writePathTrack().
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
		writeWalkElement()

	err := s.format()
	if err != nil {
		// unexpected error
		panic(err)
	}

	buf.WriteString(TrimPackageName(s.buf.String()))
}

func (s *EngineFactory) writePackageName() *EngineFactory {
	s.buf.WriteString("package main\n")
	return s
}

func (s *EngineFactory) writeImports() *EngineFactory {
	s.buf.WriteString("import \"strconv\"\n")
	return s
}

func newStateFactory(config *ast.AST) *EngineFactory {
	return &EngineFactory{
		config: config,
		buf:    &bytes.Buffer{},
	}
}

func (s *EngineFactory) format() error {
	config, err := parser.ParseFile(token.NewFileSet(), "", s.buf.String(), parser.AllErrors)
	if err != nil {
		return err
	}

	s.buf.Reset()
	err = format.Node(s.buf, token.NewFileSet(), config)
	return err
}
