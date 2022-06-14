package ast

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gertd/go-pluralize"
)

var pluralizeClient *pluralize.Client = pluralize.NewClient()

func caseInsensitiveSort(keys []string) func(i, j int) bool {
	return func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	}
}

func typeNameSort(configTypes []*ConfigType) func(i, j int) bool {
	return func(i, j int) bool {
		return configTypes[i].Name < configTypes[j].Name
	}
}

func valueTypeNameSort(fields []*Field) func(i, j int) bool {
	return func(i, j int) bool {
		return strings.ToLower(fields[i].Parent.Name+title(pluralizeClient.Singular(fields[i].Name))) < strings.ToLower(fields[j].Parent.Name+title(pluralizeClient.Singular(fields[j].Name)))
	}
}

// TODO: all this needs explanation

// "[]string" -> true
// "string" -> false
func isSliceValue(valueString string) bool {
	re := regexp.MustCompile(`\[\]`)
	return re.MatchString(valueString)
}

func isPointerValue(valueString string) bool {
	re := regexp.MustCompile(`\*`)
	return re.MatchString(valueString)
}

func isAnyValue(valueString string) bool {
	re := regexp.MustCompile(`anyOf<`)
	return re.MatchString(valueString)
}

func extractAnyTypes(valueString string) []string {
	re := regexp.MustCompile(`<.*>`)
	s := re.FindString(valueString)
	typesRe := regexp.MustCompile(`[A-Za-z]+`)
	types := typesRe.FindAllString(s, -1)
	return types
}

// "[]float64" -> float64
// "float64" -> float64
func extractValueType(valueString string) string {
	re := regexp.MustCompile(`[A-Za-z]+[0-9]*`)
	return re.FindString(valueString)
}

func getSring(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func title(name string) string {
	return strings.Title(name)
}

func anyNameByField(f Field) string {
	name := "anyOf"
	f.RangeValueTypes(func(configType *ConfigType) {
		name += title(configType.Name)
	})
	return name
}
