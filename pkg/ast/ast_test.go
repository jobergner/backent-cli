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
		"moveInEvent": map[interface{}]interface{}{
			"__event__":   "true",
			"destination": "*address",
		},
		"person": map[interface{}]interface{}{
			"name":       "string",
			"age":        "int",
			"friends":    "[]*person",
			"secondHome": "*house",
			"foo":        "anyOf<address,person>",
			"bar":        "*anyOf<address,person>",
			"baz":        "[]anyOf<address,person>",
			"ban":        "[]*anyOf<address,person>",
			"residentOf": "*city",
			"action":     "[]moveInEvent",
		},
		"city": map[interface{}]interface{}{
			"name": "string",
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

	responseData := map[interface{}]interface{}{
		"removeResidents": map[interface{}]interface{}{
			"newResidentsCount": "int",
		},
		"changeAddress": map[interface{}]interface{}{
			"isValidAddress": "bool",
		},
	}

	t.Run("should build the structure of AST", func(t *testing.T) {
		actual := buildASTStructure(stateData, actionsData, responseData)

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
					Response: map[string]Field{
						"newResidentsCount": {
							Name:            "newResidentsCount",
							ValueString:     "int",
							HasSliceValue:   false,
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
					Response: map[string]Field{
						"isValidAddress": {
							Name:            "isValidAddress",
							ValueString:     "bool",
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
				"moveInEvent": {
					Name:    "moveInEvent",
					IsEvent: true,
					Fields: map[string]Field{
						"destination": {
							Name:            "destination",
							ValueString:     "*address",
							HasSliceValue:   false,
							HasPointerValue: true,
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
						"action": {
							Name:            "action",
							ValueString:     "[]moveInEvent",
							HasSliceValue:   true,
							HasPointerValue: false,
							ValueTypes:      make(map[string]*ConfigType),
						},
						"foo": {
							Name:            "foo",
							ValueString:     "anyOf<address,person>",
							HasSliceValue:   false,
							HasPointerValue: false,
							HasAnyValue:     true,
							ValueTypes:      make(map[string]*ConfigType),
						},
						"bar": {
							Name:            "bar",
							ValueString:     "*anyOf<address,person>",
							HasSliceValue:   false,
							HasPointerValue: true,
							HasAnyValue:     true,
							ValueTypes:      make(map[string]*ConfigType),
						},
						"baz": {
							Name:            "baz",
							ValueString:     "[]anyOf<address,person>",
							HasSliceValue:   true,
							HasPointerValue: false,
							HasAnyValue:     true,
							ValueTypes:      make(map[string]*ConfigType),
						},
						"ban": {
							Name:            "ban",
							ValueString:     "[]*anyOf<address,person>",
							HasSliceValue:   true,
							HasPointerValue: true,
							HasAnyValue:     true,
							ValueTypes:      make(map[string]*ConfigType),
						},
						"residentOf": {
							Name:            "residentOf",
							ValueString:     "*city",
							HasSliceValue:   false,
							HasPointerValue: true,
							HasAnyValue:     false,
							ValueTypes:      make(map[string]*ConfigType),
						},
					},
				},
				"city": {
					Name: "city",
					Fields: map[string]Field{
						"name": {
							Name:            "name",
							ValueString:     "string",
							HasSliceValue:   false,
							HasPointerValue: false,
							ValueTypes:      make(map[string]*ConfigType),
						},
					},
				},
			},
		}

		assert.Equal(t, expected, actual)

	})

	t.Run("should fill in references of AST", func(t *testing.T) {

		actual := buildASTStructure(stateData, actionsData, responseData)
		actual.fillInReferences().fillInParentalInfo()

		houseType := actual.Types["house"]
		livingSpaceField := houseType.Fields["livingSpace"]
		assert.Equal(t, livingSpaceField.Parent.Name, "house")
		assert.Equal(t, livingSpaceField.ValueTypes["int"].Name, "int")
		assert.Equal(t, livingSpaceField.ValueTypes["int"].IsBasicType, true)
		assert.Equal(t, livingSpaceField.ValueTypeName, "int")
		residentsField := houseType.Fields["residents"]
		assert.Equal(t, residentsField.Parent.Name, "house")
		assert.Equal(t, residentsField.ValueTypes["person"].Name, "person")
		assert.Equal(t, residentsField.ValueTypes["person"].IsBasicType, false)
		assert.Equal(t, residentsField.ValueTypeName, "person")
		addressField := houseType.Fields["address"]
		assert.Equal(t, addressField.Parent.Name, "house")
		assert.Equal(t, addressField.ValueTypes["address"].Name, "address")
		assert.Equal(t, addressField.ValueTypes["address"].IsBasicType, false)
		assert.Equal(t, addressField.ValueTypeName, "address")

		addressType := actual.Types["address"]
		streetField := addressType.Fields["street"]
		assert.Equal(t, streetField.Parent.Name, "address")
		assert.Equal(t, streetField.ValueTypes["string"].Name, "string")
		assert.Equal(t, streetField.ValueTypes["string"].IsBasicType, true)
		assert.Equal(t, streetField.ValueTypeName, "string")
		houseNumberField := addressType.Fields["houseNumber"]
		assert.Equal(t, houseNumberField.Parent.Name, "address")
		assert.Equal(t, houseNumberField.ValueTypes["int"].Name, "int")
		assert.Equal(t, houseNumberField.ValueTypes["int"].IsBasicType, true)
		assert.Equal(t, houseNumberField.ValueTypeName, "int")
		cityField := addressType.Fields["city"]
		assert.Equal(t, cityField.Parent.Name, "address")
		assert.Equal(t, cityField.ValueTypes["string"].Name, "string")
		assert.Equal(t, cityField.ValueTypes["string"].IsBasicType, true)
		assert.Equal(t, cityField.ValueTypeName, "string")

		moveInEventType := actual.Types["moveInEvent"]
		destinationField := moveInEventType.Fields["destination"]
		assert.Equal(t, destinationField.Parent.Name, "moveInEvent")
		assert.Equal(t, destinationField.ValueTypes["address"].Name, "address")
		assert.Equal(t, destinationField.ValueTypes["address"].IsBasicType, false)
		assert.Equal(t, destinationField.ValueTypeName, "moveInEventDestinationRef")

		personType := actual.Types["person"]
		NameField := personType.Fields["name"]
		assert.Equal(t, NameField.Parent.Name, "person")
		assert.Equal(t, NameField.ValueTypes["string"].Name, "string")
		assert.Equal(t, NameField.ValueTypes["string"].IsBasicType, true)
		assert.Equal(t, NameField.ValueTypeName, "string")
		ageField := personType.Fields["age"]
		assert.Equal(t, ageField.Parent.Name, "person")
		assert.Equal(t, ageField.ValueTypes["int"].Name, "int")
		assert.Equal(t, ageField.ValueTypes["int"].IsBasicType, true)
		assert.Equal(t, ageField.ValueTypeName, "int")
		friendsField := personType.Fields["friends"]
		assert.Equal(t, friendsField.Parent.Name, "person")
		assert.Equal(t, friendsField.ValueTypes["person"].Name, "person")
		assert.Equal(t, friendsField.ValueTypes["person"].IsBasicType, false)
		assert.Equal(t, friendsField.ValueTypeName, "personFriendRef")
		secondHomeField := personType.Fields["secondHome"]
		assert.Equal(t, secondHomeField.Parent.Name, "person")
		assert.Equal(t, secondHomeField.ValueTypes["house"].Name, "house")
		assert.Equal(t, secondHomeField.ValueTypes["house"].IsBasicType, false)
		assert.Equal(t, secondHomeField.ValueTypeName, "personSecondHomeRef")
		fooField := personType.Fields["foo"]
		assert.Equal(t, fooField.Parent.Name, "person")
		assert.Equal(t, fooField.ValueTypes["address"].Name, "address")
		assert.Equal(t, fooField.ValueTypes["person"].Name, "person")
		assert.Equal(t, fooField.ValueTypes["address"].IsBasicType, false)
		assert.Equal(t, fooField.ValueTypes["person"].IsBasicType, false)
		assert.Equal(t, fooField.ValueTypeName, "anyOfAddress_Person")
		barField := personType.Fields["bar"]
		assert.Equal(t, barField.Parent.Name, "person")
		assert.Equal(t, barField.ValueTypes["address"].Name, "address")
		assert.Equal(t, barField.ValueTypes["person"].Name, "person")
		assert.Equal(t, barField.ValueTypes["address"].IsBasicType, false)
		assert.Equal(t, barField.ValueTypes["person"].IsBasicType, false)
		assert.Equal(t, barField.ValueTypeName, "personBarRef")
		bazField := personType.Fields["baz"]
		assert.Equal(t, bazField.Parent.Name, "person")
		assert.Equal(t, bazField.ValueTypes["address"].Name, "address")
		assert.Equal(t, bazField.ValueTypes["person"].Name, "person")
		assert.Equal(t, bazField.ValueTypes["address"].IsBasicType, false)
		assert.Equal(t, bazField.ValueTypes["person"].IsBasicType, false)
		assert.Equal(t, bazField.ValueTypeName, "anyOfAddress_Person")
		banField := personType.Fields["ban"]
		assert.Equal(t, banField.Parent.Name, "person")
		assert.Equal(t, banField.ValueTypes["address"].Name, "address")
		assert.Equal(t, banField.ValueTypes["person"].Name, "person")
		assert.Equal(t, banField.ValueTypes["address"].IsBasicType, false)
		assert.Equal(t, banField.ValueTypes["person"].IsBasicType, false)
		assert.Equal(t, banField.ValueTypeName, "personBanRef")
		residentOfField := personType.Fields["residentOf"]
		assert.Equal(t, residentOfField.Parent.Name, "person")
		assert.Equal(t, residentOfField.ValueTypes["city"].Name, "city")
		assert.Equal(t, residentOfField.ValueTypes["city"].IsBasicType, false)
		assert.Equal(t, residentOfField.ValueTypeName, "personResidentOfRef")
		actionField := personType.Fields["action"]
		assert.Equal(t, actionField.Parent.Name, "person")
		assert.Equal(t, actionField.ValueTypes["moveInEvent"].Name, "moveInEvent")
		assert.Equal(t, actionField.ValueTypes["moveInEvent"].IsBasicType, false)
		assert.Equal(t, actionField.ValueTypeName, "moveInEvent")

		cityType := actual.Types["city"]
		nameField := cityType.Fields["name"]
		assert.Equal(t, nameField.Parent.Name, "city")
		assert.Equal(t, nameField.ValueTypes["string"].Name, "string")
		assert.Equal(t, nameField.ValueTypes["string"].IsBasicType, true)
		assert.Equal(t, nameField.ValueTypeName, "string")

		removeResidentsAction := actual.Actions["removeResidents"]
		residentsParam := removeResidentsAction.Params["residents"]
		assert.Equal(t, residentsParam.ValueTypes["person"].Name, "person")
		assert.Equal(t, residentsParam.ValueTypes["person"].IsBasicType, false)
		newResidentsCountResponseValue := removeResidentsAction.Response["newResidentsCount"]
		assert.Equal(t, newResidentsCountResponseValue.ValueTypes["int"].Name, "int")
		assert.Equal(t, newResidentsCountResponseValue.ValueTypes["int"].IsBasicType, true)

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
		isValidAddressResponseValue := changeAddressAction.Response["isValidAddress"]
		assert.Equal(t, isValidAddressResponseValue.ValueTypes["bool"].Name, "bool")
		assert.Equal(t, isValidAddressResponseValue.ValueTypes["bool"].IsBasicType, true)

		assert.ElementsMatch(t, personType.ReferencedBy, []*Field{&banField, &barField, &friendsField})
		assert.ElementsMatch(t, addressType.ReferencedBy, []*Field{&banField, &barField, &destinationField})
		assert.ElementsMatch(t, houseType.ReferencedBy, []*Field{&secondHomeField})
		assert.ElementsMatch(t, cityType.ReferencedBy, []*Field{&residentOfField})
	})

	t.Run("should fill in parentalInfo", func(t *testing.T) {

		actual := buildASTStructure(stateData, actionsData, responseData)
		actual.fillInReferences().fillInParentalInfo()

		assert.True(t, actual.Types["house"].IsRootType)
		assert.False(t, actual.Types["house"].IsLeafType)
		assert.False(t, actual.Types["person"].IsRootType)
		assert.False(t, actual.Types["person"].IsLeafType)
		assert.False(t, actual.Types["address"].IsRootType)
		assert.True(t, actual.Types["address"].IsLeafType)
		assert.True(t, actual.Types["city"].IsRootType)
		assert.True(t, actual.Types["city"].IsLeafType)
	})
}

// func namesInFieldSlice(fields []*Field) []string {

// }
