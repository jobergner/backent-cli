package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateConfigAST(t *testing.T) {
	stateData := map[interface{}]interface{}{
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

	actionsData := map[interface{}]interface{}{
		"removeResidents": map[interface{}]interface{}{
			"residents": "[]person",
		},
		"changeAddress": map[interface{}]interface{}{
			"newStreet":      "string",
			"newHouseNumber": "int",
			"newCity":        "string",
		},
		"renamePerson": map[interface{}]interface{}{
			"newName": "string",
		},
	}

	t.Run("should build the structure of AST", func(t *testing.T) {
		actual := buildASTStructure(stateData, actionsData)

		expected := &AST{
			Actions: map[string]Action{
				"removeResidents": {
					Name: "removeResidents",
					Params: map[string]Field{
						"residents": {
							Name:          "residents",
							ValueString:   "[]person",
							HasSliceValue: true,
						},
					},
				},
				"changeAddress": {
					Name: "changeAddress",
					Params: map[string]Field{
						"newStreet": {
							Name:          "newStreet",
							ValueString:   "string",
							HasSliceValue: false,
						},
						"newHouseNumber": {
							Name:          "newHouseNumber",
							ValueString:   "int",
							HasSliceValue: false,
						},
						"newCity": {
							Name:          "newCity",
							ValueString:   "string",
							HasSliceValue: false,
						},
					},
				},
				"renamePerson": {
					Name: "renamePerson",
					Params: map[string]Field{
						"newName": {
							Name:          "newName",
							ValueString:   "string",
							HasSliceValue: false,
						},
					},
				},
			},
			Types: map[string]ConfigType{
				"house": {
					Name: "house",
					Fields: map[string]Field{
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
					Fields: map[string]Field{
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
					Fields: map[string]Field{
						"name": {
							Name:          "name",
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

	t.Run("should fill in references of AST", func(t *testing.T) {

		actual := buildASTStructure(stateData, actionsData)
		actual.fillInReferences().fillInParentalInfo()

		houseType := actual.Types["house"]
		livingSpaceField := houseType.Fields["livingSpace"]
		assert.Equal(t, livingSpaceField.ValueType.Name, "int")
		assert.Equal(t, livingSpaceField.ValueType.IsBasicType, true)
		residentsField := houseType.Fields["residents"]
		assert.Equal(t, residentsField.ValueType.Name, "person")
		assert.Equal(t, residentsField.ValueType.IsBasicType, false)
		addressField := houseType.Fields["address"]
		assert.Equal(t, addressField.ValueType.Name, "address")
		assert.Equal(t, addressField.ValueType.IsBasicType, false)

		addressType := actual.Types["address"]
		streetField := addressType.Fields["street"]
		assert.Equal(t, streetField.ValueType.Name, "string")
		assert.Equal(t, streetField.ValueType.IsBasicType, true)
		houseNumberField := addressType.Fields["houseNumber"]
		assert.Equal(t, houseNumberField.ValueType.Name, "int")
		assert.Equal(t, houseNumberField.ValueType.IsBasicType, true)
		cityField := addressType.Fields["city"]
		assert.Equal(t, cityField.ValueType.Name, "string")
		assert.Equal(t, cityField.ValueType.IsBasicType, true)

		personType := actual.Types["person"]
		NameField := personType.Fields["name"]
		assert.Equal(t, NameField.ValueType.Name, "string")
		assert.Equal(t, NameField.ValueType.IsBasicType, true)
		ageField := personType.Fields["age"]
		assert.Equal(t, ageField.ValueType.Name, "int")
		assert.Equal(t, ageField.ValueType.IsBasicType, true)

		removeResidentsAction := actual.Actions["removeResidents"]
		residentsParam := removeResidentsAction.Params["residents"]
		assert.Equal(t, residentsParam.ValueType.Name, "person")
		assert.Equal(t, residentsParam.ValueType.IsBasicType, false)

		changeAddressAction := actual.Actions["changeAddress"]
		newStreetParam := changeAddressAction.Params["newStreet"]
		assert.Equal(t, newStreetParam.ValueType.Name, "string")
		assert.Equal(t, newStreetParam.ValueType.IsBasicType, true)
		newHouseNumberParam := changeAddressAction.Params["newHouseNumber"]
		assert.Equal(t, newHouseNumberParam.ValueType.Name, "int")
		assert.Equal(t, newHouseNumberParam.ValueType.IsBasicType, true)
		newCityParam := changeAddressAction.Params["newCity"]
		assert.Equal(t, newCityParam.ValueType.Name, "string")
		assert.Equal(t, newCityParam.ValueType.IsBasicType, true)
	})

	t.Run("should fill in parentalInfo", func(t *testing.T) {

		actual := buildASTStructure(stateData, actionsData)
		actual.fillInReferences().fillInParentalInfo()

		assert.True(t, actual.Types["house"].IsRootType)
		assert.False(t, actual.Types["house"].IsLeafType)
		assert.False(t, actual.Types["person"].IsRootType)
		assert.True(t, actual.Types["person"].IsLeafType)
		assert.False(t, actual.Types["address"].IsRootType)
		assert.True(t, actual.Types["address"].IsLeafType)
	})
}
