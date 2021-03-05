package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleAST(t *testing.T) {
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
			"Name": "string",
			"age":  "int",
		},
	}

	t.Run("should build a rudimentary simpleAST from data", func(t *testing.T) {
		actual := buildRudimentarySimpleAST(data)

		expected := simpleAST{
			Decls: map[string]simpleTypeDecl{
				"house": {
					Name: "house",
					Fields: map[string]simpleFieldDecl{
						"residents": {
							Name:          "residents",
							ValueString:   "[]person",
							HasSliceValue: true,
						},
						"livingSpace": {
							Name:          "livingSpace",
							ValueString:   "int",
							HasSliceValue: false,
						},
						"address": {
							Name:          "address",
							ValueString:   "address",
							HasSliceValue: false,
						},
					},
				},
				"address": {
					Name: "address",
					Fields: map[string]simpleFieldDecl{
						"street": {
							Name:          "street",
							ValueString:   "string",
							HasSliceValue: false,
						},
						"houseNumber": {
							Name:          "houseNumber",
							ValueString:   "int",
							HasSliceValue: false,
						},
						"city": {
							Name:          "city",
							ValueString:   "string",
							HasSliceValue: false,
						},
					},
				},
				"person": {
					Name: "person",
					Fields: map[string]simpleFieldDecl{
						"Name": {
							Name:          "Name",
							ValueString:   "string",
							HasSliceValue: false,
						},
						"age": {
							Name:          "age",
							ValueString:   "int",
							HasSliceValue: false,
						},
					},
				},
			},
		}

		assert.Equal(t, expected, actual)

	})

	t.Run("should fill in references of rudimentary simpleAST", func(t *testing.T) {

		actual := buildRudimentarySimpleAST(data)
		actual.fillInReferences().fillInParentalInfo()

		livingSpaceField := actual.Decls["house"].Fields["livingSpace"]
		assert.Equal(t, livingSpaceField.Parent.Name, "house")
		assert.Equal(t, livingSpaceField.ValueType.Name, "int")
		assert.Equal(t, livingSpaceField.ValueType.IsBasicType, true)
		residentsField := actual.Decls["house"].Fields["residents"]
		assert.Equal(t, residentsField.Parent.Name, "house")
		assert.Equal(t, residentsField.ValueType.Name, "person")
		assert.Equal(t, residentsField.ValueType.IsBasicType, false)
		addressField := actual.Decls["house"].Fields["address"]
		assert.Equal(t, addressField.Parent.Name, "house")
		assert.Equal(t, addressField.ValueType.Name, "address")
		assert.Equal(t, addressField.ValueType.IsBasicType, false)

		streetField := actual.Decls["address"].Fields["street"]
		assert.Equal(t, streetField.Parent.Name, "address")
		assert.Equal(t, streetField.ValueType.Name, "string")
		assert.Equal(t, streetField.ValueType.IsBasicType, true)
		houseNumberField := actual.Decls["address"].Fields["houseNumber"]
		assert.Equal(t, houseNumberField.Parent.Name, "address")
		assert.Equal(t, houseNumberField.ValueType.Name, "int")
		assert.Equal(t, houseNumberField.ValueType.IsBasicType, true)
		cityField := actual.Decls["address"].Fields["city"]
		assert.Equal(t, cityField.Parent.Name, "address")
		assert.Equal(t, cityField.ValueType.Name, "string")
		assert.Equal(t, cityField.ValueType.IsBasicType, true)

		NameField := actual.Decls["person"].Fields["Name"]
		assert.Equal(t, NameField.Parent.Name, "person")
		assert.Equal(t, NameField.ValueType.Name, "string")
		assert.Equal(t, NameField.ValueType.IsBasicType, true)
		ageField := actual.Decls["person"].Fields["age"]
		assert.Equal(t, ageField.Parent.Name, "person")
		assert.Equal(t, ageField.ValueType.Name, "int")
		assert.Equal(t, ageField.ValueType.IsBasicType, true)
	})

	t.Run("should fill in parentalInfo", func(t *testing.T) {

		actual := buildRudimentarySimpleAST(data)
		actual.fillInReferences().fillInParentalInfo()

		assert.True(t, actual.Decls["house"].IsRootType)
		assert.False(t, actual.Decls["house"].IsLeafType)
		assert.False(t, actual.Decls["person"].IsRootType)
		assert.True(t, actual.Decls["person"].IsLeafType)
		assert.False(t, actual.Decls["address"].IsRootType)
		assert.True(t, actual.Decls["address"].IsLeafType)
	})
}
