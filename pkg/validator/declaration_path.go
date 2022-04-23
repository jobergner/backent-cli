package validator

import (
	"fmt"
	"regexp"
)

type pathClosureKind int

const (
	pathClosureKindUndetermined pathClosureKind = iota
	// path ends due to it being recusrive
	pathClosureKindRecursiveness
	// path ends due to encountering a reference value (*string, []string)
	pathClosureKindReference
	// path ends due to encountering a basic type (string, int)
	pathClosureKindBasicType
)

type dataValueKind int

const (
	valueKindString dataValueKind = iota
	valueKindObject
)

type fieldLevelKind int

const (
	// current level is the object itself
	fieldLevelZero fieldLevelKind = iota
	// current level is the root level of the object
	firstFieldLevel
	secondFieldLevel
)

// TODO: name it to be more like golang.ast package naming
type declaration struct {
	name          string
	dataValueKind dataValueKind
	fieldLevel    fieldLevelKind
}

type declarationPath struct {
	declarations []declaration
	closureKind  pathClosureKind
}

func (path *declarationPath) setClosureKind(closureKind pathClosureKind) {
	path.closureKind = closureKind
}

func (path declarationPath) isRecursive() bool {
	for i, declaration := range path.declarations {
		for j, _declaration := range path.declarations {
			if i == j {
				continue
			}
			if declaration.fieldLevel != firstFieldLevel || _declaration.fieldLevel != firstFieldLevel {
				continue
			}
			if declaration.name == _declaration.name {
				return true
			}
		}
	}
	return false
}

func (path *declarationPath) addDeclaration(
	keyName string,
	dataValueKind dataValueKind,
	fieldLevel fieldLevelKind,
	value interface{},
) {
	path.declarations = append(path.declarations, declaration{keyName, dataValueKind, fieldLevel})
}

// we list the declarations' typeNames with some additional logic
// to be more explicit (paths to nested field names will
// be concatenated eg. "foo.bar")
func (path declarationPath) joinedNames() []string {
	var joinedNames []string

	var wasStructDecl bool
	var parentStructName string

	for _, declaration := range path.declarations {
		if declaration.dataValueKind == valueKindObject {
			wasStructDecl = true
			parentStructName = declaration.name
			continue
		}

		if wasStructDecl && declaration.fieldLevel == secondFieldLevel {
			joinedNames = append(joinedNames, parentStructName+"."+declaration.name)
		} else {
			joinedNames = append(joinedNames, declaration.name)
		}
		wasStructDecl = false
		parentStructName = ""

	}

	// if path ended on struct declaration (it can happen if it's a recursive path)
	if wasStructDecl {
		joinedNames = append(joinedNames, parentStructName)
	}

	return joinedNames
}

func (path declarationPath) copySelf() declarationPath {
	declarationsCopy := make([]declaration, len(path.declarations))
	copy(declarationsCopy, path.declarations)
	pathCopy := path
	pathCopy.declarations = declarationsCopy
	return pathCopy
}

type pathBuilder struct {
	data  map[interface{}]interface{}
	paths []declarationPath
}

func newPathBuilder(data map[interface{}]interface{}) *pathBuilder {
	return &pathBuilder{data: data}
}

func (pb *pathBuilder) addPath(path declarationPath) {
	pb.paths = append(pb.paths, path)
}

// a recursive function to travel through the data
func (pb *pathBuilder) build(path declarationPath, keyName string, value interface{}, fieldLevel fieldLevelKind) {

	switch {
	case isString(value):
		path.addDeclaration(keyName, valueKindString, fieldLevel, value)

		if path.isRecursive() {
			// detected recursiveness implies this is the end of the path
			path.setClosureKind(pathClosureKindRecursiveness)
			pb.addPath(path)
			return
		}

		valueLiteral := fmt.Sprintf("%v", value)

		if isReferenceType(valueLiteral) {
			path.addDeclaration(valueLiteral, valueKindString, fieldLevel, value)
			// a reference type implies this is the end of the path
			path.setClosureKind(pathClosureKindReference)
			pb.addPath(path)
			return
		}

		// "string" -> true, "[2]string" -> true
		if containsOnlyBasicTypes(valueLiteral) {
			path.addDeclaration(valueLiteral, valueKindString, fieldLevelZero, value)
			// a basic type implies this is the end of the path
			path.setClosureKind(pathClosureKindBasicType)
			pb.addPath(path)
			return
		}

		// we extract the type so literals describing simple arrays like "[23]foo" become "foo"
		// this only ever has an effect on arrays because all other types would be either reference types
		// (e.g. "[]foo" or "map[foo]bar") and returned above, or named types like "foo"
		nextTypeLiteral := extractTypes(valueLiteral)[0]
		nextValue := pb.data[nextTypeLiteral]
		pb.build(path, nextTypeLiteral, nextValue, firstFieldLevel)

	case isMap(value):
		path.addDeclaration(keyName, valueKindObject, fieldLevel, value)

		if path.isRecursive() {
			// detected recursiveness implies this is the end of the path
			path.setClosureKind(pathClosureKindRecursiveness)
			pb.addPath(path)
			return
		}

		mapValue := value.(map[interface{}]interface{})

		for _key, _value := range mapValue {
			// the path is copied; this is basically a fork
			pathCopy := path.copySelf()
			_keyName := fmt.Sprintf("%v", _key)
			// we go a level deeper (fieldLevel+1) and handle each key/value
			// pair in the next pb.build() execution
			pb.build(pathCopy, _keyName, _value, fieldLevel+1)
		}

	}
}

func isBasicType(typeString string) bool {
	for _, basicType := range golangBasicTypes {
		if basicType == typeString {
			return true
		}
	}
	return false
}

func containsOnlyBasicTypes(declarationTypeString string) bool {
	extractedTyped := extractTypes(declarationTypeString)
	for _, extractType := range extractedTyped {
		if !isBasicType(extractType) {
			return false
		}
	}
	return true
}

// TODO: this should be revisited at some point
func isReferenceType(declarationTypeString string) bool {
	re := regexp.MustCompile(`\[\]|\*|map\[`)
	return re.MatchString(declarationTypeString)
}

func containsUncomparableValue(typeName string, data map[interface{}]interface{}) bool {
	pathBuilder := newPathBuilder(data)

	pathBuilder.build(declarationPath{}, typeName, data[typeName], firstFieldLevel)

	var isUncomparable bool
	for _, path := range pathBuilder.paths {
		if path.closureKind == pathClosureKindRecursiveness || path.closureKind == pathClosureKindReference {
			isUncomparable = true
		}
	}

	return isUncomparable
}
