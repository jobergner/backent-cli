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
			"foo":        "anyOf<address,person>",
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
							ValueTypes:      make(map[string]*ConfigType),
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
							ValueTypes:      make(map[string]*ConfigType),
						},
						"newHouseNumber": {
							Name:            "newHouseNumber",
							ValueString:     "int",
							HasSliceValue:   false,
							HasPointerValue: false,
							ValueTypes:      make(map[string]*ConfigType),
						},
						"newCity": {
							Name:            "newCity",
							ValueString:     "string",
							HasSliceValue:   false,
							HasPointerValue: false,
							ValueTypes:      make(map[string]*ConfigType),
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
							ValueTypes:      make(map[string]*ConfigType),
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
							ValueTypes:      make(map[string]*ConfigType),
						},
						"livingSpace": {
							Name:            "livingSpace",
							ValueString:     "int",
							HasSliceValue:   false,
							HasPointerValue: false,
							ValueTypes:      make(map[string]*ConfigType),
						},
						"address": {
							Name:            "address",
							ValueString:     "address",
							HasSliceValue:   false,
							HasPointerValue: false,
							ValueTypes:      make(map[string]*ConfigType),
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
							ValueTypes:      make(map[string]*ConfigType),
						},
						"houseNumber": {
							Name:            "houseNumber",
							ValueString:     "int",
							HasSliceValue:   false,
							HasPointerValue: false,
							ValueTypes:      make(map[string]*ConfigType),
						},
						"city": {
							Name:            "city",
							ValueString:     "string",
							HasSliceValue:   false,
							HasPointerValue: false,
							ValueTypes:      make(map[string]*ConfigType),
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
							ValueTypes:      make(map[string]*ConfigType),
						},
						"age": {
							Name:            "age",
							ValueString:     "int",
							HasSliceValue:   false,
							HasPointerValue: false,
							ValueTypes:      make(map[string]*ConfigType),
						},
						"friends": {
							Name:            "friends",
							ValueString:     "[]*person",
							HasSliceValue:   true,
							HasPointerValue: true,
							ValueTypes:      make(map[string]*ConfigType),
						},
						"secondHome": {
							Name:            "secondHome",
							ValueString:     "*house",
							HasSliceValue:   false,
							HasPointerValue: true,
							ValueTypes:      make(map[string]*ConfigType),
						},
						"foo": {
							Name:            "foo",
							ValueString:     "anyOf<address,person>",
							HasSliceValue:   false,
							HasPointerValue: false,
							ValueTypes:      make(map[string]*ConfigType),
							HasAnyValue:     true,
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
		assert.Equal(t, livingSpaceField.ValueTypes["int"].Name, "int")
		assert.Equal(t, livingSpaceField.ValueTypes["int"].IsBasicType, true)
		residentsField := houseType.Fields["residents"]
		assert.Equal(t, residentsField.ValueTypes["person"].Name, "person")
		assert.Equal(t, residentsField.ValueTypes["person"].IsBasicType, false)
		addressField := houseType.Fields["address"]
		assert.Equal(t, addressField.ValueTypes["address"].Name, "address")
		assert.Equal(t, addressField.ValueTypes["address"].IsBasicType, false)

		addressType := actual.Types["address"]
		streetField := addressType.Fields["street"]
		assert.Equal(t, streetField.ValueTypes["string"].Name, "string")
		assert.Equal(t, streetField.ValueTypes["string"].IsBasicType, true)
		houseNumberField := addressType.Fields["houseNumber"]
		assert.Equal(t, houseNumberField.ValueTypes["int"].Name, "int")
		assert.Equal(t, houseNumberField.ValueTypes["int"].IsBasicType, true)
		cityField := addressType.Fields["city"]
		assert.Equal(t, cityField.ValueTypes["string"].Name, "string")
		assert.Equal(t, cityField.ValueTypes["string"].IsBasicType, true)

		personType := actual.Types["person"]
		NameField := personType.Fields["name"]
		assert.Equal(t, NameField.ValueTypes["string"].Name, "string")
		assert.Equal(t, NameField.ValueTypes["string"].IsBasicType, true)
		ageField := personType.Fields["age"]
		assert.Equal(t, ageField.ValueTypes["int"].Name, "int")
		assert.Equal(t, ageField.ValueTypes["int"].IsBasicType, true)
		friendsField := personType.Fields["friends"]
		assert.Equal(t, friendsField.ValueTypes["person"].Name, "person")
		assert.Equal(t, friendsField.ValueTypes["person"].IsBasicType, false)
		secondHomeField := personType.Fields["secondHome"]
		assert.Equal(t, secondHomeField.ValueTypes["house"].Name, "house")
		assert.Equal(t, secondHomeField.ValueTypes["house"].IsBasicType, false)

		removeResidentsAction := actual.Actions["removeResidents"]
		residentsParam := removeResidentsAction.Params["residents"]
		assert.Equal(t, residentsParam.ValueTypes["person"].Name, "person")
		assert.Equal(t, residentsParam.ValueTypes["person"].IsBasicType, false)

		changeAddressAction := actual.Actions["changeAddress"]
		newStreetParam := changeAddressAction.Params["newStreet"]
		assert.Equal(t, newStreetParam.ValueTypes["string"].Name, "string")
		assert.Equal(t, newStreetParam.ValueTypes["string"].IsBasicType, true)
		newHouseNumberParam := changeAddressAction.Params["newHouseNumber"]
		assert.Equal(t, newHouseNumberParam.ValueTypes["int"].Name, "int")
		assert.Equal(t, newHouseNumberParam.ValueTypes["int"].IsBasicType, true)
		newCityParam := changeAddressAction.Params["newCity"]
		assert.Equal(t, newCityParam.ValueTypes["string"].Name, "string")
		assert.Equal(t, newCityParam.ValueTypes["string"].IsBasicType, true)

		assert.Equal(t, personType.ReferencedBy, []*Field{&friendsField})
		assert.Equal(t, houseType.ReferencedBy, []*Field{&secondHomeField})
	})

	t.Run("should fill in parentalInfo", func(t *testing.T) {

		actual := buildASTStructure(stateData, actionsData)
		actual.fillInReferences().fillInParentalInfo()

		assert.True(t, actual.Types["house"].IsRootType)
		assert.False(t, actual.Types["house"].IsLeafType)
		assert.False(t, actual.Types["person"].IsRootType)
		assert.False(t, actual.Types["person"].IsLeafType)
		assert.False(t, actual.Types["address"].IsRootType)
		assert.True(t, actual.Types["address"].IsLeafType)
	})
}
