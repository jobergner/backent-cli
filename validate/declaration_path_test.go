package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldLevels(t *testing.T) {
	t.Run("should build declarations with expected fieldLevels", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "string",
			},
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "", data, fieldLevelZero)

		assert.Equal(t, 1, len(pb.paths))
		assert.Equal(t, []string{"foo.bar", "string"}, pb.paths[0].joinedNames())
		assert.Contains(t, pb.paths[0].declarations, declaration{"foo", valueKindObject, firstFieldLevel})
		assert.Contains(t, pb.paths[0].declarations, declaration{"bar", valueKindString, secondFieldLevel})
		assert.Contains(t, pb.paths[0].declarations, declaration{"string", valueKindString, fieldLevelZero})
	})
}

func TestClosureKind(t *testing.T) {
	t.Run("should have closureKind closureKindBasicType", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "bar",
			"bar": "string",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "foo", data["foo"], firstFieldLevel)
		assert.Equal(t, pb.paths[0].closureKind, pathClosureKindBasicType)
	})

	t.Run("should have closureKind closureKindReference (1/3)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "bar",
			"bar": "*string",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "foo", data["foo"], firstFieldLevel)
		assert.Equal(t, pb.paths[0].closureKind, pathClosureKindReference)
	})

	t.Run("should have closureKind closureKindReference (2/3)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "bar",
			"bar": "[]string",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "foo", data["foo"], firstFieldLevel)
		assert.Equal(t, pb.paths[0].closureKind, pathClosureKindReference)
	})

	t.Run("should have closureKind closureKindReference (3/3)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "bar",
			"bar": "map[string]int",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "foo", data["foo"], firstFieldLevel)
		assert.Equal(t, pb.paths[0].closureKind, pathClosureKindReference)
	})

	t.Run("should have closureKind closureKindRecursiveness", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "bar",
			"bar": "foo",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "foo", data["foo"], firstFieldLevel)
		assert.Equal(t, pb.paths[0].closureKind, pathClosureKindRecursiveness)
	})
}

func TestPathBuilder(t *testing.T) {
	t.Run("should build one path of flat types and exit on loop", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "bar",
			"ban": "foo",
			"bar": "ban",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "ban", data["ban"], firstFieldLevel)

		assert.Equal(t, 1, len(pb.paths))
		assert.Equal(t, []string{"ban", "foo", "bar", "ban"}, pb.paths[0].joinedNames())
	})

	t.Run("should build one path of flat types that exits on a basic type", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "bar",
			"bar": "string",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "foo", data["foo"], firstFieldLevel)

		assert.Equal(t, 1, len(pb.paths))
		assert.Equal(t, []string{"foo", "bar", "string"}, pb.paths[0].joinedNames())
	})

	t.Run("should build one path of a flat type directly", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": "string",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "bar", data["bar"], firstFieldLevel)

		assert.Equal(t, 1, len(pb.paths))
		assert.Equal(t, []string{"bar", "string"}, pb.paths[0].joinedNames())
	})

	t.Run("should build one path of nested types", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "baz",
			},
			"baz": "bar",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "baz", data["baz"], firstFieldLevel)

		assert.Equal(t, 1, len(pb.paths))
		assert.Equal(t, []string{"baz", "bar.foo", "baz"}, pb.paths[0].joinedNames())
	})

	t.Run("should build two paths of nested types", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "baz",
				"bam": "string",
			},
			"baz": "bar",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "baz", data["baz"], firstFieldLevel)

		assert.Equal(t, 2, len(pb.paths))
		joinedNamess := [][]string{
			pb.paths[0].joinedNames(),
			pb.paths[1].joinedNames(),
		}

		assert.Contains(t, joinedNamess, []string{"baz", "bar.foo", "baz"})
		assert.Contains(t, joinedNamess, []string{"baz", "bar.bam", "string"})
	})

	t.Run("should build two paths of nested types including arrays", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "[23]buf",
				"bam": "string",
			},
			"baz": "[2]bar",
			"buf": "string",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "baz", data["baz"], firstFieldLevel)

		assert.Equal(t, 2, len(pb.paths))
		joinedNamess := [][]string{
			pb.paths[0].joinedNames(),
			pb.paths[1].joinedNames(),
		}

		assert.Contains(t, joinedNamess, []string{"baz", "bar.foo", "buf", "string"})
		assert.Contains(t, joinedNamess, []string{"baz", "bar.bam", "string"})
	})

	t.Run("should path ending with array of basic type", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"baz": "[2]string",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "baz", data["baz"], firstFieldLevel)

		assert.Equal(t, 1, len(pb.paths))
		joinedNamess := [][]string{
			pb.paths[0].joinedNames(),
		}

		assert.Contains(t, joinedNamess, []string{"baz", "[2]string"})
	})

	t.Run("should build a path with a itself-referring struct", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "bar",
			},
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "bar", data["bar"], firstFieldLevel)

		assert.Equal(t, 1, len(pb.paths))
		assert.Equal(t, []string{"bar.foo", "bar"}, pb.paths[0].joinedNames())
	})

	t.Run("should build multiple paths of nested types", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "baz",
				"bam": "string",
				"bal": "bar",
				"fof": "bas",
			},
			"bas": map[interface{}]interface{}{
				"ban":  "string",
				"bunt": "bant",
			},
			"baz":  "bar",
			"bant": "int",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "baz", data["baz"], firstFieldLevel)

		assert.Equal(t, 5, len(pb.paths))
		joinedNamess := [][]string{
			pb.paths[0].joinedNames(),
			pb.paths[1].joinedNames(),
			pb.paths[2].joinedNames(),
			pb.paths[3].joinedNames(),
			pb.paths[4].joinedNames(),
		}

		assert.Contains(t, joinedNamess, []string{"baz", "bar.foo", "baz"})
		assert.Contains(t, joinedNamess, []string{"baz", "bar.bam", "string"})
		assert.Contains(t, joinedNamess, []string{"baz", "bar.bal", "bar"})
		assert.Contains(t, joinedNamess, []string{"baz", "bar.fof", "bas.ban", "string"})
		assert.Contains(t, joinedNamess, []string{"baz", "bar.fof", "bas.bunt", "bant", "int"})
	})

	t.Run("should build paths from data root", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "baz",
				"bam": "string",
			},
			"baz": "bar",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "", data, fieldLevelZero)

		assert.Equal(t, 4, len(pb.paths))
		joinedNamess := [][]string{
			pb.paths[0].joinedNames(),
			pb.paths[1].joinedNames(),
			pb.paths[2].joinedNames(),
			pb.paths[3].joinedNames(),
		}

		assert.Contains(t, joinedNamess, []string{"baz", "bar.foo", "baz"})
		assert.Contains(t, joinedNamess, []string{"baz", "bar.bam", "string"})
		assert.Contains(t, joinedNamess, []string{"bar.bam", "string"})
		assert.Contains(t, joinedNamess, []string{"bar.foo", "baz", "bar"})
	})

	t.Run("should build recursive path from self referencing type", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "foo",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "", data, fieldLevelZero)

		assert.Equal(t, 1, len(pb.paths))
		joinedNamess := [][]string{
			pb.paths[0].joinedNames(),
		}

		assert.Contains(t, joinedNamess, []string{"foo", "foo"})
	})

	t.Run("should build 2 recursive paths on self referencing objects", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "baz",
			},
			"baz": map[interface{}]interface{}{
				"ban": "bar",
			},
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "", data, fieldLevelZero)

		assert.Equal(t, 2, len(pb.paths))
		joinedNamess := [][]string{
			pb.paths[0].joinedNames(),
			pb.paths[1].joinedNames(),
		}

		assert.Contains(t, joinedNamess, []string{"bar.foo", "baz.ban", "bar"})
		assert.Contains(t, joinedNamess, []string{"baz.ban", "bar.foo", "baz"})
	})

	t.Run("should stop paths on reference type used", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "*baz",
				"fan": "[]baz",
				"faz": "map[int]baz",
			},
			"baz": "int",
			"ban": "bar",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "ban", data["ban"], firstFieldLevel)

		assert.Equal(t, 3, len(pb.paths))
		joinedNamess := [][]string{
			pb.paths[0].joinedNames(),
			pb.paths[1].joinedNames(),
			pb.paths[2].joinedNames(),
		}

		assert.Contains(t, joinedNamess, []string{"ban", "bar.foo", "*baz"})
		assert.Contains(t, joinedNamess, []string{"ban", "bar.fan", "[]baz"})
		assert.Contains(t, joinedNamess, []string{"ban", "bar.faz", "map[int]baz"})
	})

	t.Run("should stop paths on reference type used even withtin array", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "[23][]buf",
			},
			"baz": "[2]bar",
			"buf": "string",
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "baz", data["baz"], firstFieldLevel)

		assert.Equal(t, 1, len(pb.paths))
		joinedNamess := [][]string{
			pb.paths[0].joinedNames(),
		}

		assert.Contains(t, joinedNamess, []string{"baz", "bar.foo", "[23][]buf"})
	})

}

func TestDeclarationPath(t *testing.T) {
	t.Run("should build declaration path", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "baz",
			},
			"baz": map[interface{}]interface{}{
				"ban": "bar",
			},
		}

		pb := pathBuilder{
			paths: []declarationPath{},
			data:  data,
		}

		pb.build(declarationPath{}, "", data, fieldLevelZero)

		assert.Equal(t, 2, len(pb.paths))
		joinedNamess := [][]string{
			pb.paths[0].joinedNames(),
			pb.paths[1].joinedNames(),
		}

		assert.Contains(t, joinedNamess, []string{"bar.foo", "baz.ban", "bar"})
		assert.Contains(t, joinedNamess, []string{"baz.ban", "bar.foo", "baz"})
	})
}

func TestContainsOnlyBasicTypes(t *testing.T) {
	t.Run("should detect basic types", func(t *testing.T) {
		assert.Equal(t, containsOnlyBasicTypes("string"), true)
		assert.Equal(t, containsOnlyBasicTypes("*string"), true)
		assert.Equal(t, containsOnlyBasicTypes("*foo"), false)
		assert.Equal(t, containsOnlyBasicTypes("int"), true)
		assert.Equal(t, containsOnlyBasicTypes("*int"), true)
		assert.Equal(t, containsOnlyBasicTypes("[]string"), true)
		assert.Equal(t, containsOnlyBasicTypes("map[int]string"), true)
		assert.Equal(t, containsOnlyBasicTypes("map[foo]string"), false)
		assert.Equal(t, containsOnlyBasicTypes("map[*foo]string"), false)
		assert.Equal(t, containsOnlyBasicTypes("map[*string]string"), true)
	})
}

func TestIsReferenceType(t *testing.T) {
	t.Run("should detect reference types", func(t *testing.T) {
		assert.Equal(t, isReferenceType("string"), false)
		assert.Equal(t, isReferenceType("*string"), true)
		assert.Equal(t, isReferenceType("*foo"), true)
		assert.Equal(t, isReferenceType("int"), false)
		assert.Equal(t, isReferenceType("*int"), true)
		assert.Equal(t, isReferenceType("[]string"), true)
		assert.Equal(t, isReferenceType("[23]string"), false)
		assert.Equal(t, isReferenceType("*[23]string"), true)
		assert.Equal(t, isReferenceType("[23]*string"), true)
		assert.Equal(t, isReferenceType("[23][]int"), true)
		assert.Equal(t, isReferenceType("map[int]string"), true)
		assert.Equal(t, isReferenceType("map[foo]string"), true)
		assert.Equal(t, isReferenceType("map[*foo]string"), true)
		assert.Equal(t, isReferenceType("map[*string]string"), true)
		assert.Equal(t, isReferenceType("map[int][23]string"), true)
		assert.Equal(t, isReferenceType("map[[23]int]string"), true)
	})
}

func TestContainsUncomparableValue(t *testing.T) {
	t.Run("should not find reference value (1/2)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "string",
			"bar": "foo",
		}

		assert.Equal(t, containsUncomparableValue("bar", data), false)
	})
	t.Run("should not find reference value (2/2)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "bar",
			"bar": map[interface{}]interface{}{
				"ban": "string",
				"baz": "int",
			},
		}

		assert.Equal(t, containsUncomparableValue("foo", data), false)
	})
	t.Run("should find reference value (1/4)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "[]string",
			"bar": "foo",
		}

		assert.Equal(t, containsUncomparableValue("bar", data), true)
	})
	t.Run("should find reference value (2/4)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "map[int]string",
			"bar": "foo",
		}

		assert.Equal(t, containsUncomparableValue("bar", data), true)
	})
	t.Run("should find reference value (3/4)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "*string",
			"bar": "foo",
		}

		assert.Equal(t, containsUncomparableValue("bar", data), true)
	})
	t.Run("should find reference value (4/4)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "bar",
			"bar": map[interface{}]interface{}{
				"ban": "[]string",
				"baz": "int",
			},
		}

		assert.Equal(t, containsUncomparableValue("foo", data), true)
	})
}
