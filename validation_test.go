package validator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findErrorIn(val error, slice []error) (int, bool) {
	for i, item := range slice {
		if item.Error() == val.Error() {
			return i, true
		}
	}
	return -1, false
}

func removeErrorFromSlice(slice []error, index int) []error {
	return append(slice[:index], slice[index+1:]...)
}

func matchErrors(actualErrors, expectedErrors []error) (leftoverErrors, redundantErrors []error) {
	// redefine redunantErrors so it never returns as nil (which happens when there are no redunant errors)
	// so comparing error slices becomes more conventient
	redundantErrors = make([]error, 0)
	leftoverErrors = make([]error, len(expectedErrors))
	copy(leftoverErrors, expectedErrors)

	for _, actualError := range actualErrors {
		leftoverErrorIndex, isFound := findErrorIn(actualError, leftoverErrors)
		if isFound {
			leftoverErrors = removeErrorFromSlice(leftoverErrors, leftoverErrorIndex)
		} else {
			redundantErrors = append(redundantErrors, actualError)
		}
	}

	return
}

func TestMatchErrors(t *testing.T) {
	t.Run("should return no errors when all errors matched", func(t *testing.T) {
		actualErrors := []error{errors.New("abc"), errors.New("def"), errors.New("ghi")}
		expectedErrors := []error{errors.New("abc"), errors.New("def"), errors.New("ghi")}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Equal(t, 0, len(missingErrors))
		assert.Equal(t, 0, len(redundantErrors))
	})

	t.Run("should return missing errors when some are missing", func(t *testing.T) {
		actualErrors := []error{errors.New("abc")}
		expectedErrors := []error{errors.New("abc"), errors.New("def"), errors.New("ghi")}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		expectedMissingErrors := []error{errors.New("def"), errors.New("ghi")}
		assert.Equal(t, expectedMissingErrors, missingErrors)
		expectedRedundantErrors := []error{}
		assert.Equal(t, expectedRedundantErrors, redundantErrors)
	})

	t.Run("should return redundant errors when some are missing", func(t *testing.T) {
		actualErrors := []error{errors.New("abc"), errors.New("def"), errors.New("ghi")}
		expectedErrors := []error{errors.New("abc")}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		expectedMissingErrors := []error{}
		assert.Equal(t, expectedMissingErrors, missingErrors)
		expectedRedundantErrors := []error{errors.New("def"), errors.New("ghi")}
		assert.Equal(t, expectedRedundantErrors, redundantErrors)
	})

	t.Run("output should not be index order", func(t *testing.T) {
		actualErrors := []error{errors.New("abc"), errors.New("def"), errors.New("ghi")}
		expectedErrors := []error{errors.New("def"), errors.New("ghi"), errors.New("abc")}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		expectedMissingErrors := []error{}
		assert.Equal(t, expectedMissingErrors, missingErrors)
		expectedRedundantErrors := []error{}
		assert.Equal(t, expectedRedundantErrors, redundantErrors)
	})
}

func TestGeneralValidation(t *testing.T) {
	t.Run("should procude no errors", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": "string",
			"baf": "[2]foo",
			"bal": "map[foo]bar",
			"bum": "*int",
			"baz": map[interface{}]interface{}{
				"ban":  "int32",
				"bam":  "bar",
				"bunt": "[]baf",
				"bap":  "map[bar]foo",
				"bal":  "***bar",
				"slap": "**[]**baf",
				"arg":  "*baz",
				"barg": "[3]foo",
			},
		}

		errs := generalValidation(data)

		assert.Empty(t, errs)
	})

	t.Run("should procude expected structural errors", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": "",
			"baf": []string{"foo"},
			"bal": [1]string{"foo"},
			"baz": map[interface{}]interface{}{
				"ban": map[interface{}]interface{}{
					"lan": "int32",
				},
				"ran": []string{"foo"},
				"kan": [1]string{"foo"},
			},
		}

		actualErrors := generalValidation(data)
		expectedErrors := []error{
			newValidationErrorIllegalValue("bar", "root"),
			newValidationErrorIllegalValue("baf", "root"),
			newValidationErrorIllegalValue("bal", "root"),
			newValidationErrorIllegalValue("ban", "baz"),
			newValidationErrorIllegalValue("ran", "baz"),
			newValidationErrorIllegalValue("kan", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should procude expected syntactical errors", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"a":    "[]string",
			"b":    "map[int]string]",
			"fo o": "int[]",
			"bar":  "[]map[int]string",
			"bu":   "bar",
			"baz": map[interface{}]interface{}{
				"ba$n": "[]in[t32",
				"fan":  "[]",
				"c":    "map[int][]string",
				"d":    "map[int]string*",
				"e":    "[*]int32",
				"f":    "int*",
				"g":    "[]in t32",
				"h":    " ",
				"i":    "[]in@t32",
				"j":    "fo o",
			},
		}

		actualErrors := generalValidation(data)
		expectedErrors := []error{
			newValidationErrorInvalidValueString("map[int]string]", "b", "root"),
			newValidationErrorInvalidValueString("int[]", "fo o", "root"),
			newValidationErrorInvalidValueString("[]in[t32", "ba$n", "baz"),
			newValidationErrorInvalidValueString("[]", "fan", "baz"),
			newValidationErrorInvalidValueString("map[int]string*", "d", "baz"),
			newValidationErrorInvalidValueString("[*]int32", "e", "baz"),
			newValidationErrorInvalidValueString("int*", "f", "baz"),
			newValidationErrorInvalidValueString("[]in t32", "g", "baz"),
			newValidationErrorInvalidValueString(" ", "h", "baz"),
			newValidationErrorInvalidValueString("[]in@t32", "i", "baz"),
			newValidationErrorInvalidValueString("fo o", "j", "baz"),
			newValidationErrorIllegalTypeName("fo o", "root"),
			newValidationErrorIllegalTypeName("ba$n", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should procude logical errors", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo":  "int",
			"bunt": "map[int]string",
			"bar": map[interface{}]interface{}{
				"foo": "bam",
			},
			"baz": map[interface{}]interface{}{
				"ban": "bar",
				"bor": "[23]kan",
			},
			"bam": map[interface{}]interface{}{
				"baf": "baz",
				"ban": "map[[2]foo]int",
				"buf": "map[[]foo]int",
				"bul": "map[bunt]int",
			},
		}

		actualErrors := generalValidation(data)
		expectedErrors := []error{
			newValidationErrorRecursiveTypeUsage([]string{"bam.baf", "baz.ban", "bar.foo", "bam"}),
			newValidationErrorRecursiveTypeUsage([]string{"baz.ban", "bar.foo", "bam.baf", "baz"}),
			newValidationErrorRecursiveTypeUsage([]string{"bar.foo", "bam.baf", "baz.ban", "bar"}),
			newValidationErrorInvalidMapKey("[]foo", "map[[]foo]int"),
			newValidationErrorInvalidMapKey("bunt", "map[bunt]int"),
			newValidationErrorTypeNotFound("kan", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}

func TestValidateStateConfig(t *testing.T) {
	t.Run("should procude no errors", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{},
			"baz": map[interface{}]interface{}{
				"ban":  "int32",
				"bam":  "bar",
				"bunt": "[]int",
			},
		}

		errs := validateStateConfig(data)

		assert.Empty(t, errs)
	})

	t.Run("should procude expected errors", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": "float64",
			"baz": map[interface{}]interface{}{
				"ban":       "int32",
				"bam":       "bar",
				"bunt":      "[]foo",
				"bap":       "map[bar]foo",
				"bal":       "***bar",
				"slap":      "**[]**bar",
				"arg":       "*foo",
				"unt":       "*[]int",
				"rnt":       "[]*int",
				"barg":      "[3]foo",
				"iD":        "string",
				"hasParent": "bool",
			},
		}

		actualErrors := validateStateConfig(data)
		expectedErrors := []error{
			newValidationErrorUnavailableFieldName("iD"),
			newValidationErrorUnavailableFieldName("hasParent"),
			newValidationErrorNonObjectType("foo"),
			newValidationErrorNonObjectType("bar"),
			newValidationErrorIncompatibleValue("map[bar]foo", "bap", "baz"),
			newValidationErrorIncompatibleValue("***bar", "bal", "baz"),
			newValidationErrorIncompatibleValue("**[]**bar", "slap", "baz"),
			newValidationErrorIncompatibleValue("[3]foo", "barg", "baz"),
			newValidationErrorIncompatibleValue("*[]int", "unt", "baz"),
			newValidationErrorIncompatibleValue("[]*int", "rnt", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}

func TestValidateActionsConfig(t *testing.T) {
	t.Run("should procude expected errors", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": "float64",
			"baz": map[interface{}]interface{}{
				"ban":  "int32",
				"bam":  "bar",
				"bunt": "[]foo",
				"bap":  "map[bar]foo",
				"bal":  "***bar",
				"slap": "**[]**bar",
				"arg":  "*foo",
				"barg": "[3]foo",
			},
		}
		actionsConfigData := map[interface{}]interface{}{
			"fooAction": map[interface{}]interface{}{
				"foo": "int32",
			},
			"barAction": map[interface{}]interface{}{
				"bug":  "fooAction",
				"bam":  "baz",
				"bum":  "*baz",
				"lum":  "bazID",
				"foot": "string",
				"feet": "string",
			},
			"BazAction": map[interface{}]interface{}{
				"baz": "int32",
			},
		}

		actualErrors := ValidateActionsConfig(data, actionsConfigData)
		expectedErrors := []error{
			newValidationErrorTypeNotFound("fooAction", "barAction"),
			newValidationErrorIllegalCapitalization("BazAction", literalKindType),
			newValidationErrorDirectTypeUsage("barAction", "baz"),
			newValidationErrorIllegalPointerParameter("barAction", "bum"),
			newValidationErrorDirectTypeUsage("barAction", "fooAction"),
			newValidationErrorDirectTypeUsage("barAction", "*baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}

func TestValidateStateConfigWithAnyOfTypes(t *testing.T) {
	t.Run("returns only pre-validation errors if existent (1/2)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": nil,
			"baz": map[interface{}]interface{}{
				"ban": "",
			},
		}

		actualErrors := ValidateStateConfig(data)
		expectedErrors := []error{
			newValidationErrorIllegalValue("bar", "root"),
			newValidationErrorIllegalValue("ban", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
	t.Run("returns only pre-validation errors if existent (2/2)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"baz": map[interface{}]interface{}{
				"ban": "int",
			},
		}

		actualErrors := ValidateStateConfig(data)
		expectedErrors := []error{
			newValidationErrorNonObjectType("foo"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
	t.Run("on invalid anyOf definitions, it just considers field as ill-defined", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"baz": map[interface{}]interface{}{
				"ban": "anyof<>",
			},
		}

		actualErrors := ValidateStateConfig(data)
		expectedErrors := []error{
			newValidationErrorInvalidValueString("anyof<>", "ban", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
	t.Run("returns errors of different combinations when multiple exist", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"fan": "string",
			},
			"bar": map[interface{}]interface{}{
				"lar": "anyOf<foo,baz>",
			},
			"baz": map[interface{}]interface{}{
				"ban": "anyOf<bar,foo>",
			},
		}

		actualErrors := ValidateStateConfig(data)
		expectedErrors := []error{
			newValidationErrorRecursiveTypeUsage([]string{"bar.lar", "baz.ban", "bar"}),
			newValidationErrorRecursiveTypeUsage([]string{"baz.ban", "bar.lar", "baz"}),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
	t.Run("does not cause any erros on valid definition", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"fan": "string",
			},
			"bar": map[interface{}]interface{}{
				"lar": "anyOf<foo,baz>",
			},
			"baz": map[interface{}]interface{}{
				"ban": "anyOf<foo,fam>",
			},
			"fam": map[interface{}]interface{}{
				"lam": "int",
			},
		}

		actualErrors := ValidateStateConfig(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
