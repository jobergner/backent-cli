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

func TestValidateEventHandling(t *testing.T) {
	t.Run("should procude no errors even with meta field included", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"fooNonEvent": map[interface{}]interface{}{
				"bar": "string",
			},
			"fooEvent": map[interface{}]interface{}{
				"__event__": "true",
				"bar":       "string",
			},
			"barEvent": map[interface{}]interface{}{
				"__event__": "true",
				"bar":       "string",
			},
			"baz": map[interface{}]interface{}{
				"bun": "anyOf<barEvent,fooEvent>",
			},
		}

		actualErrors := ValidateStateConfig(data)
		expectedErrors := []error{
			newValidationErrorInvalidEventUsage("anyOf<barEvent,fooEvent>"),
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

		errs := ValidateStateConfig(data)

		assert.Empty(t, errs)
	})

	t.Run("should procude expected errors", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{},
			"bar": map[interface{}]interface{}{},
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

		actualErrors := ValidateStateConfig(data)
		expectedErrors := []error{
			newValidationErrorUnavailableFieldName("iD"),
			newValidationErrorUnavailableFieldName("hasParent"),
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
	t.Run("validates structure if actions (no panicking)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "int64",
			},
		}
		actionsConfigData := map[interface{}]interface{}{
			"fooAction": map[interface{}]interface{}{
				"foo": "fooID",
			},
			"bar": "int64",
		}

		actualErrors := ValidateActionsConfig(data, actionsConfigData)
		expectedErrors := []error{
			newValidationErrorNonObjectType("bar"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
	t.Run("should procude expected errors", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "int64",
			},
			"wii": map[interface{}]interface{}{
				"wuu": "int64",
			},
			"baz": map[interface{}]interface{}{
				"ban":  "int64",
				"bam":  "foo",
				"bunt": "[]foo",
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
	t.Run("returns invalid anyOf definition error", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"fan": "string",
			},
			"bar": map[interface{}]interface{}{
				"lar": "anyOf<foo>",
			},
			"baz": map[interface{}]interface{}{
				"ban": "anyOf<bar,bar>",
			},
		}

		actualErrors := ValidateStateConfig(data)
		expectedErrors := []error{
			newValidationErrorInvalidAnyOfDefinition("anyOf<foo>"),
			newValidationErrorInvalidAnyOfDefinition("anyOf<bar,bar>"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
	t.Run("on invalid anyOf definitions, it just considers field as ill-defined", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"baz": map[interface{}]interface{}{
				"ban": "anyOf<>",
			},
		}

		actualErrors := ValidateStateConfig(data)
		expectedErrors := []error{
			newValidationErrorInvalidAnyOfDefinition("anyOf<>"),
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
				"lar": "anyOf<baz,foo>",
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
				"lar": "anyOf<baz,foo>",
			},
			"baz": map[interface{}]interface{}{
				"ban": "anyOf<fam,foo>",
				"ran": "[]anyOf<bar,foo>",
				"kan": "*anyOf<bar,foo>",
				"man": "[]*anyOf<bar,foo>",
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

func TestValidateResponsesConfig(t *testing.T) {
	t.Run("validates responses just like actions but returns error due to response of unknown name", func(t *testing.T) {
		stateConfigData := map[interface{}]interface{}{
			"baz": map[interface{}]interface{}{
				"ban": "string",
			},
		}
		actionsConfigData := map[interface{}]interface{}{
			"dooFoo": map[interface{}]interface{}{
				"bar": "int",
			},
		}
		responsesConfigData := map[interface{}]interface{}{
			"dooFoo": map[interface{}]interface{}{
				"ban": "bazID",
				"bau": "*baz",
			},
		}

		actualErrors := ValidateResponsesConfig(stateConfigData, actionsConfigData, responsesConfigData)
		expectedErrors := []error{
			newValidationErrorDirectTypeUsage("dooFoo", "*baz"),
			newValidationErrorIllegalPointerParameter("dooFoo", "bau"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}

// func TestFoo(t *testing.T) {
// 	var StateConfig = map[interface{}]interface{}{
// 		"player": map[interface{}]interface{}{
// 			"items":         "[]item",
// 			"equipmentSets": "[]*equipmentSet",
// 			"gearScore":     "gearScore",
// 			"position":      "position",
// 			"guildMembers":  "[]*player",
// 			"target":        "*anyOf<player,zoneItem>",
// 			"targetedBy":    "[]*anyOf<player,zoneItem>",
// 		},
// 		"zone": map[interface{}]interface{}{
// 			"items":         "[]zoneItem",
// 			"players":       "[]player",
// 			"tags":          "[]string",
// 			"interactables": "[]anyOf<item,player,zoneItem>",
// 		},
// 		"zoneItem": map[interface{}]interface{}{
// 			"position": "position",
// 			"item":     "item",
// 		},
// 		"position": map[interface{}]interface{}{
// 			"x": "float64",
// 			"y": "float64",
// 		},
// 		"item": map[interface{}]interface{}{
// 			"name":      "string",
// 			"gearScore": "gearScore",
// 			"boundTo":   "*player",
// 			"origin":    "anyOf<player,position>",
// 		},
// 		"gearScore": map[interface{}]interface{}{
// 			"level": "int",
// 			"score": "int",
// 		},
// 		"equipmentSet": map[interface{}]interface{}{
// 			"name":      "string",
// 			"equipment": "[]*item",
// 		},
// 	}

// 	t.Run("returns only pre-validation errors if existent (1/2)", func(t *testing.T) {

// 		actualErrors := ValidateStateConfig(StateConfig)
// 		expectedErrors := []error{}

// 		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

// 		assert.Empty(t, missingErrors)
// 		assert.Empty(t, redundantErrors)
// 	})
// 	var ActionsConfig = map[interface{}]interface{}{
// 		"movePlayer": map[interface{}]interface{}{
// 			"player":  "playerID",
// 			"changeX": "float64",
// 			"changeY": "float64",
// 		},
// 		"addItemToPlayer": map[interface{}]interface{}{
// 			"item":    "itemID",
// 			"newName": "string",
// 		},
// 		"spawnZoneItems": map[interface{}]interface{}{
// 			"items": "[]itemID",
// 		},
// 	}
// 	t.Run("returns only pre-validation errors if existent (1/2)", func(t *testing.T) {

// 		actualErrors := ValidateActionsConfig(StateConfig, ActionsConfig)
// 		expectedErrors := []error{}

// 		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

// 		assert.Empty(t, missingErrors)
// 		assert.Empty(t, redundantErrors)
// 	})
// }
