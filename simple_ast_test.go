package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleAST(t *testing.T) {
	nilSimpleStruct := func() *simpleStructDecl { return nil }
	t.Run("should build a rudimentary simpleAST from data", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"house": map[interface{}]interface{}{
				"residents":   "[]person",
				"livingSpace": "int",
				"address":     "address",
			},
			"address": map[interface{}]interface{}{
				"street":      "string",
				"houseNumber": "int",
				"city":        "string",
			},
			"person": map[interface{}]interface{}{
				"name": "string",
				"age":  "int",
			},
		}

		actual := buildRudimentarySimpleAST(data)

		expected := simpleAST{
			decls: map[string]simpleStructDecl{
				"house": {
					name: "house",
					fields: map[string]simpleFieldDecl{
						"residents": {
							name:          "residents",
							valueString:   "[]person",
							hasSliceValue: true,
						},
						"livingSpace": {
							name:          "livingSpace",
							valueString:   "int",
							hasSliceValue: false,
						},
						"address": {
							name:          "address",
							valueString:   "address",
							hasSliceValue: false,
						},
					},
				},
				"address": {
					name: "address",
					fields: map[string]simpleFieldDecl{
						"street": {
							name:          "street",
							valueString:   "string",
							hasSliceValue: false,
						},
						"houseNumber": {
							name:          "houseNumber",
							valueString:   "int",
							hasSliceValue: false,
						},
						"city": {
							name:          "city",
							valueString:   "string",
							hasSliceValue: false,
						},
					},
				},
				"person": {
					name: "person",
					fields: map[string]simpleFieldDecl{
						"name": {
							name:          "name",
							valueString:   "string",
							hasSliceValue: false,
						},
						"age": {
							name:          "age",
							valueString:   "int",
							hasSliceValue: false,
						},
					},
				},
			},
		}

		assert.Equal(t, expected, actual)

	})

	t.Run("should fill in references of rudimentary simpleAST", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"house": map[interface{}]interface{}{
				"residents":   "[]person",
				"livingSpace": "int",
				"address":     "address",
			},
			"address": map[interface{}]interface{}{
				"street":      "string",
				"houseNumber": "int",
				"city":        "string",
			},
			"person": map[interface{}]interface{}{
				"name": "string",
				"age":  "int",
			},
		}

		actual := buildRudimentarySimpleAST(data)
		actual.fillInReferences()

		livingSpaceField := actual.decls["house"].fields["livingSpace"]
		assert.Equal(t, livingSpaceField.parent.name, "house")
		assert.Equal(t, livingSpaceField.valueType, nilSimpleStruct())
		residentsField := actual.decls["house"].fields["residents"]
		assert.Equal(t, residentsField.parent.name, "house")
		assert.Equal(t, residentsField.valueType.name, "person")
		addressField := actual.decls["house"].fields["address"]
		assert.Equal(t, addressField.parent.name, "house")
		assert.Equal(t, addressField.valueType.name, "address")

		streetField := actual.decls["address"].fields["street"]
		assert.Equal(t, streetField.parent.name, "address")
		assert.Equal(t, streetField.valueType, nilSimpleStruct())
		houseNumberField := actual.decls["address"].fields["houseNumber"]
		assert.Equal(t, houseNumberField.parent.name, "address")
		assert.Equal(t, houseNumberField.valueType, nilSimpleStruct())
		cityField := actual.decls["address"].fields["city"]
		assert.Equal(t, cityField.parent.name, "address")
		assert.Equal(t, cityField.valueType, nilSimpleStruct())

		nameField := actual.decls["person"].fields["name"]
		assert.Equal(t, nameField.parent.name, "person")
		assert.Equal(t, nameField.valueType, nilSimpleStruct())
		ageField := actual.decls["person"].fields["age"]
		assert.Equal(t, ageField.parent.name, "person")
		assert.Equal(t, ageField.valueType, nilSimpleStruct())
	})
}
