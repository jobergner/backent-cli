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
	f.RangeValueTypes(func(configType *ast.ConfigType) {
		name += Title(configType.Name)
	})
	return name
}

type EngineFactory struct {
	config *ast.AST
	buf    *bytes.Buffer
}

// WriteEngineFrom writes source code for a given State-/ActionsConfig
func WriteEngineFrom(stateConfigData map[interface{}]interface{}) []byte {
	config := ast.Parse(stateConfigData, map[interface{}]interface{}{})
	s := newStateFactory(config).
		writePackageName().
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

	return s.writtenSourceCode()
}

func (s *EngineFactory) writePackageName() *EngineFactory {
	s.buf.WriteString("package state\n")
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

func (s *EngineFactory) writtenSourceCode() []byte {
	return s.buf.Bytes()
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
