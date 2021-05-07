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
			"name":       "string",
			"age":        "int",
			"friends":    "[]*person",
			"secondHome": "*house",
			// "foo":        "anyOf<house,person>",
			// "bar":        "*anyOf<house,person>",
			// "baz":        "[]anyOf<house,person>",
			// "ban":        "[]*anyOf<house,person>",
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
							Name:            "residents",
							ValueString:     "[]person",
							HasSliceValue:   true,
							HasPointerValue: false,
						},
					},
				},
				"changeAddress": {
					Name: "changeAddress",
					Params: map[string]Field{
						"newStreet": {
							Name:            "newStreet",
							ValueString:     "string",
							HasSliceValue:   false,
							HasPointerValue: false,
						},
						"newHouseNumber": {
							Name:            "newHouseNumber",
							ValueString:     "int",
							HasSliceValue:   false,
							HasPointerValue: false,
						},
						"newCity": {
							Name:            "newCity",
							ValueString:     "string",
							HasSliceValue:   false,
							HasPointerValue: false,
						},
					},
				},
				"renamePerson": {
					Name: "renamePerson",
					Params: map[string]Field{
						"newName": {
							Name:            "newName",
							ValueString:     "string",
							HasSliceValue:   false,
							HasPointerValue: false,
						},
					},
				},
			},
			Types: map[string]ConfigType{
				"house": {
					Name: "house",
					Fields: map[string]Field{
						"residents": {
							Name:            "residents",
							ValueString:     "[]person",
							HasSliceValue:   true,
							HasPointerValue: false,
						},
						"livingSpace": {
							Name:            "livingSpace",
							ValueString:     "int",
							HasSliceValue:   false,
							HasPointerValue: false,
						},
						"address": {
							Name:            "address",
							ValueString:     "address",
							HasSliceValue:   false,
							HasPointerValue: false,
						},
					},
				},
				"address": {
					Name: "address",
					Fields: map[string]Field{
						"street": {
							Name:            "street",
							ValueString:     "string",
							HasSliceValue:   false,
							HasPointerValue: false,
						},
						"houseNumber": {
							Name:            "houseNumber",
							ValueString:     "int",
							HasSliceValue:   false,
							HasPointerValue: false,
						},
						"city": {
							Name:            "city",
							ValueString:     "string",
							HasSliceValue:   false,
							HasPointerValue: false,
						},
					},
				},
				"person": {
					Name: "person",
					Fields: map[string]Field{
						"name": {
							Name:            "name",
							ValueString:     "string",
							HasSliceValue:   false,
							HasPointerValue: false,
						},
						"age": {
							Name:            "age",
							ValueString:     "int",
							HasSliceValue:   false,
							HasPointerValue: false,
						},
						"friends": {
							Name:            "friends",
							ValueString:     "[]*person",
							HasSliceValue:   true,
							HasPointerValue: true,
						},
						"secondHome": {
							Name:            "secondHome",
							ValueString:     "*house",
							HasSliceValue:   false,
							HasPointerValue: true,
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
		assert.Equal(t, livingSpaceField.ValueTypes[0].Name, "int")
		assert.Equal(t, livingSpaceField.ValueTypes[0].IsBasicType, true)
		residentsField := houseType.Fields["residents"]
		assert.Equal(t, residentsField.ValueTypes[0].Name, "person")
		assert.Equal(t, residentsField.ValueTypes[0].IsBasicType, false)
		addressField := houseType.Fields["address"]
		assert.Equal(t, addressField.ValueTypes[0].Name, "address")
		assert.Equal(t, addressField.ValueTypes[0].IsBasicType, false)

		addressType := actual.Types["address"]
		streetField := addressType.Fields["street"]
		assert.Equal(t, streetField.ValueTypes[0].Name, "string")
		assert.Equal(t, streetField.ValueTypes[0].IsBasicType, true)
		houseNumberField := addressType.Fields["houseNumber"]
		assert.Equal(t, houseNumberField.ValueTypes[0].Name, "int")
		assert.Equal(t, houseNumberField.ValueTypes[0].IsBasicType, true)
		cityField := addressType.Fields["city"]
		assert.Equal(t, cityField.ValueTypes[0].Name, "string")
		assert.Equal(t, cityField.ValueTypes[0].IsBasicType, true)

		personType := actual.Types["person"]
		NameField := personType.Fields["name"]
		assert.Equal(t, NameField.ValueTypes[0].Name, "string")
		assert.Equal(t, NameField.ValueTypes[0].IsBasicType, true)
		ageField := personType.Fields["age"]
		assert.Equal(t, ageField.ValueTypes[0].Name, "int")
		assert.Equal(t, ageField.ValueTypes[0].IsBasicType, true)
		friendsField := personType.Fields["friends"]
		assert.Equal(t, friendsField.ValueTypes[0].Name, "person")
		assert.Equal(t, friendsField.ValueTypes[0].IsBasicType, false)
		secondHomeField := personType.Fields["secondHome"]
		assert.Equal(t, secondHomeField.ValueTypes[0].Name, "house")
		assert.Equal(t, secondHomeField.ValueTypes[0].IsBasicType, false)

		removeResidentsAction := actual.Actions["removeResidents"]
		residentsParam := removeResidentsAction.Params["residents"]
		assert.Equal(t, residentsParam.ValueTypes[0].Name, "person")
		assert.Equal(t, residentsParam.ValueTypes[0].IsBasicType, false)

		changeAddressAction := actual.Actions["changeAddress"]
		newStreetParam := changeAddressAction.Params["newStreet"]
		assert.Equal(t, newStreetParam.ValueTypes[0].Name, "string")
		assert.Equal(t, newStreetParam.ValueTypes[0].IsBasicType, true)
		newHouseNumberParam := changeAddressAction.Params["newHouseNumber"]
		assert.Equal(t, newHouseNumberParam.ValueTypes[0].Name, "int")
		assert.Equal(t, newHouseNumberParam.ValueTypes[0].IsBasicType, true)
		newCityParam := changeAddressAction.Params["newCity"]
		assert.Equal(t, newCityParam.ValueTypes[0].Name, "string")
		assert.Equal(t, newCityParam.ValueTypes[0].IsBasicType, true)

		assert.Equal(t, personType.ReferencedBy, []*Field{&friendsField})
		assert.Equal(t, houseType.ReferencedBy, []*Field{&secondHomeField})
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
