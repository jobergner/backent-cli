# backent-cli
backent-cli provides a toolkit to generate a server which enables real-time state broadcasting of entities through websockets.

Create your own API which braodcasts all changes to entities to the connected clients automatically:
```golang
func broadcastChanges(params state.ReceivedParams, engine *state.Engine) {

	player := engine.CreatePlayer()                  // creating the player

	player.SetName(params.Name)                      // setting the player name

	player.AddItem().SetName(params.FirstItemName)   // add an item and set item name

}
```

# Start Experimenting!
Explore backent-cli and its features with the Inspector and toy around until you feel comfortable. 
```bash
# set up directories
mkdir backent_example; cd backent_example;
go mod init backentexample;

# generate the code
backent-cli generate -example -out ./backent/;

# run and use the inspector
go run .;
backent-cli inspect -port 3496;
```
The Inspector is a graphical user interface for you to run locally and inspect your backent-cli generated servers, or in this case an example server that backent-cli will set up for you.
# NICE SCREENSHOT OF INSPECTOR
# Overview
## Example Configuration
This is what a very simple `config.json` can look like. You may define as many types and actions as you like. You can read more about the possible configurations below!
```JSON
{
    "state": {
        "player": {
            "name": "string",
            "items": "[]item"
        },
        "item": {
            "name": "string",
            "value": "float64"
        }
    },
    "actions": {
        "createPlayer": {
            "name": "string",
            "firstItemName": "string"
        }
    }
}
```
## Generate Server:
This is how you would generate the backent package as soon as you've defined your entity `config.json`.
```bash
backent-cli generate -config config.json -out backent/
```
## Run the Server:
```golang
package main

// import the API
import (
    state "yourproject/generatedpackage"
)

// how many processing frames per second the server will run
const fps = 30

func main() {
	err := state.Start(actions, sideEffects, fps, 3496)
	if err != nil {
		panic(err)
	}
}

// define what is being executed on receiving a message
var actions = state.Actions{
	CreatePlayer: func(params state.CreatePlayerParams, engine *state.Engine) {
		player := engine.CreatePlayer()                // creating the player

		player.SetName(params.Name)                    // setting the player name

		player.AddItem().SetName(params.FirstItemName) // add an item and set item name

	}, // state change is automatically broadcasted
}

// define what is being executed on server deploy and after all actions for a processing frame tick are processed
var sideEffects = state.SideEffects{
	OnDeploy:    func(engine *state.Engine) {},
	OnFrameTick: func(engine *state.Engine) {},
}
```
## Send a Message to trigger an Action:
This is how a message the server understands may look like. It interprets the message to trigger actions. In this case the server will trigger the `CreatePlayer` action with the given data passed as parameter.
```JSON
{
    "kind": "createPlayer",
    "content": "{\"name\": \"string\",\"firstItemName\": \"string\"}"
}
```
# The Basics
## Defining the Config:
The config's syntax is inspired by Go's own syntax. If you have knowledge of Go you will intuitively understand what is going on. And if you find yourself struggling and make mistakes, comprehensive error messages will help you correct them. There are however some additional restrictions to which values you can use where. More info on that here.

The config may consist of 3 parts: `state`, `actions` and `responses`.

### state:
The state consists of types which you can consider the equivalent to Go's structs: Structures with field names and values describing the types. As it is with go, when defining a type, you can use it as a field's value:
```JSON
{
  "address": {
    "streetName": "string",
    "houseNumber": "int"
  },
  "house": {
    "address": "address"
  }
}
```
More about defining state types can be found here.
### actions:
Actions is how the server and client communicate. Here you can define which client data you want the server to be aware of in order to react with the assigned behaviour. Defining actions is very similar to defining the state except for some additional limitations. Only Go's basic types, and type references in the form of IDs can be used as values.
```JSON
{
  "buildNewHouse": {
    "streetName": "string",
    "houseNumber": "int"
  }
}
```
The generator creates all necessary structures for you to simply tell the server how to react to a received action. The `params` struct will contain the defined data. Make use of your personal API to manipulate the `engine`'s state. The engine will keep track of all the changes you have made and tell the server to broadcast the changes to all connected clients automatically. 
```golang
// ...
var actions = state.Actions{
	BuildNewHouse: func(params state.BuildNewHouseParams, engine *state.Engine) {
		house := engine.CreateHouse()
		address := house.Address()
		address.SetStreetName(params.StreetName)
		address.SetHouseNumber(params.HouseNumber)
	},
}
// ...
```
### responses:
Sometimes you may want to send back data to the client who sent the action. This can be done with responses. Defining responses works exactly the same way as defining actions. Assigning a response to an actions works by giving the response the same name as the action:
```JSON
{
  "actions": {
    "buildNewHouse": {
      "streetName": "string",
      "houseNumber": "int"
    }
  }
}
{
  "responses": {
    "buildNewHouse": {
      "houseID": "houseID",
    }
  }
}
```
Now we can tell the server to return data to the client.
```golang
// ...
var actions = state.Actions{
	BuildNewHouse: func(params state.BuildNewHouseParams, engine *state.Engine) state.BuildNewHouseResponse {
		house := engine.CreateHouse()
		address := house.Address()
		address.SetStreetName(params.StreetName)
		address.SetHouseNumber(params.HouseNumber)
		// `ID()` is a getter method to acces the ID of an entity
		return state.BuildNewHouseResponse{houseID: house.ID()} // <- return data to client
	},
}
// ...
```

## State Structure and Updates:
Updates are assembled in a tree-like structure, containing only entities that have updated or who's children have updated. In the action section we have learned how to create a new entity of the `house` type. Creating an entity automatically creates all it's children with default values, even if they are not modified. It is just what you'd expect from Go. So the tree update of just the `engine.CreateHouse()` call alone woud look like this:
```JSON
{
    "house": {
        "1": {
            "address": {
                "id": 2,
                "streetName": "",
                "houseNumber": 0,
                "operationKind": "UPDATE"
            },
            "operationKind": "UPDATE"
        }
    }
}
```
This is the data every connected client would receive as result of a `engine.CreateHouse()` call. You can see how each created entity has a `operationKind:"UPDATE"` value. This tells the client that this entity is new or has updated since the last received update.

Imagine you defined a second action with the name `changeHouseNumber` which behaves like this:
```golang
var actions = state.Actions{
  // ...
	ChangeHouseNumber: func(params state.ChangeHouseNumberParams, engine *state.Engine) {
      house := engine.House(params.HouseID)
      house.Address().SetHouseNumber(params.NewHouseNumber)
	},
}
```
Triggering this action with a message would result in the following tree update:
```JSON
{
    "house": {
        "1": {
            "address": {
                "id": 2,
                "streetName": "",
                "houseNumber": 1,
                "operationKind": "UPDATE"
            },
            "operationKind": "UNCHANGED"
        }
    }
}
```
As the the `house` entity itself has not updated, but only it's child `address`, it maintains the `operationKind:"UNCHANGED"` value. This way the client can tell that the `house` entity has remained the same since the last update.

### How Slices Work:
Slices behave like you'd expect slices to work. However, to make all element paths within a tree structure immutable, slices are marshalled as maps. This way we can use the element's ID instead of it's index which could shift during entity modification. In a scenario where your config looks like this:
```JSON
{
  "address": {
    "streetName": "string",
    "houseNumber": "int"
  },
  "house": {
    "address": "address",
    "residents": "[]person"
  },
  "person": {
    "name": "string"
  }
}
```
and you trigger an action which looks like this:
```golang
// ...
var actions = state.Actions{
	AddResidentToHouse: func(params state.AddResidentToHouseParams, engine *state.Engine) {
		house := engine.House(params.HouseID)
		house.AddResident()
	},
}
// ...
```
this would be the emitted update:
```JSON
{
    "house": {
        "1": {
            "residents": {
                "2": {
                  "id": 2,
                  "name": "",
                  "operationKind": "UPDATE"
                }
            },
            "operationKind": "UPDATE"
        }
    }
}
```
(note how the `house` has `operationKind:"UPDATE"` as it's `residents` field got modified)

As you can see even though `residents` is defined as slice, and a getter call of `house.Residents()` would retrieve a slice of `person`, the field is marshalled as if it was a map. This way this particular `person` will always have the same path within the tree throughout it`s entire lifecycle.

# Advanced Types:
## Type References:
Sometimes you want an entity to have a certain value, but not necessarily own that value, as the value is an entity that exists on itself, and not as a child of another entity. This can be done by using references. An example that would make it's usefullness clear would be this one:
```JSON
{
    "menu": {
        "dishes": "[]*dish",
        "glutenFree": "[]*dish",
        "vegetarian": "[]*dish",
        "todaysDeal": "*dish"
    },
    "dish": {
        "name": "string",
        "ingredients": "[]string"
    }
}
```
Read here on how to use the API to handle references.

## The `anyOf` Type:
The `anyOf` type is a Quality of Life feature which lets you define fields to contain more than one type. This brings great flexibility with no additional overhead:
```JSON
{
    "farm": {
        "owner": "string",
        "animals": "[]anyOf<chicken,cow,pig>"
    },
    "chicken": {
        "eggsPerDay": "int"
    },
    "cow": {
        "weight": "float64"
    },
    "pig": {
        "name": "string"
    }
}
```
Read here on how to use the API to handle `anyOf` types.

# API Reference
## getters
The value of every field can be retrieved by calling the name of the field. Given the following config:
```JSON
{
  "address": {
    "streetName": "string",
    "houseNumber": "int"
  },
  "house": {
    "address": "address",
    "residents": "[]person"
  },
  "person": {
    "name": "string"
  }
}
```
the values can be retrieved like this:
```golang
house := engine.House(id)                  // house
residents := house.Residents()             // []person
address := house.Address()                 // address
streetName := house.Address().StreetName() // string
name := house.Residents()[0].Name()        // string
```
In the case of a referenced value:
```JSON
{
    "menu": {
        "todaysDeal": "*dish"
    },
    "dish": {
        "name": "string",
        "ingredients": "[]string"
    }
}
```
retrieve the values like this:
```golang
menu := engine.Menu(id)                // menu
dealRef, isSet := menu.TodaysDeal()    // reference object of dish, bool whether it's set
isSet = dealRef.IsSet()                // also bool whether it's set
deal := dealRef.Get()                  // dish
```
More on references and their methods can be read here.

In case of fields with `anyOf` types:
```JSON
{
    "farm": {
        "owner": "string",
        "cutestResident": "anyOf<chicken,cow,pig>"
    },
    "chicken": {
        "eggsPerDay": "int"
    },
    "cow": {
        "weight": "float64"
    },
    "pig": {
        "name": "string"
    }
}
```
retrieve the values like this:
```golang
farm := engine.Farm(id)                   // farm
cutestResident := farm.CutestResident()   // chicken|cow|pig
animalKind := cutestResident.Kind()       // "Chicken"|"Cow"|"Pig"
cow := cutestResident.Cow()               // cow
// if you know of which kind the animal is you can retrieve it directly
cow := farm.CutestResident().Cow()        // cow
```
More on `anyOf` types and their methods can be read here.

If you try to retrieve a value where there is none, all manipulations applied to this entity will have no effect.
This can happen during the following curcumstances:
```JSON
{
    "foo": {
        "bam": "*bar",
        "bal": "anyOf<bar,baz>"
    }
}
``` 
```golang
foo := engine.Foo(123)           // foo with id 123 may not exist
bam := foo.Bam()                 // bam may not be set
balRef, isSet := foo.Bal()
bar := balRef.Bar()              // bal may be of type baz and not bar
```

## creators
The `engine` has creator methods for all entities. They are as straightforward as it gets:
```JSON
{
    "chicken": {
        "eggsPerDay": "int"
    },
    "cow": {
        "weight": "float64"
    },
    "pig": {
        "name": "string"
    }
}
```
```golang
chicken := engine.CreateChicken()
cow := engine.CreateCow()
pig := engine.CreatePig()
```
## deleters
You can delete created elements just as easily as you created them:
```JSON
{
    "chicken": {
        "eggsPerDay": "int"
    },
    "cow": {
        "weight": "float64"
    },
    "pig": {
        "name": "string"
    }
}
```
```golang
engine.DeleteChicken(chickenID)
engine.DeleteCow(cowID)
engine.DeletePig(pigID)
```

## setters
Every field with a value of one of Go's basic types has a setter method to set the value. The method will always be called `Set<FieldName>`. The following config:
```JSON
{
    "chicken": {
        "eggsPerDay": "int"
    },
    "cow": {
        "weight": "float64"
    },
    "pig": {
        "name": "string"
    }
}
```
will come with these setters:
```golang
engine.Chicken(chickenID).SetEggsPerDay(3)
engine.Cow(cowID).SetWeight(56.4)
engine.Pig(pigID).SetName("Gunter")
```
Referenced values will come with additional setters when they are not part of a slice:
```JSON
{
    "chicken": {
        "bestFriend": "*chicken",
    }
}
```
```golang
chicken := engine.Chicken(id)
friendlyChicken := engine.CreateChicken()
_, isSet := chicken.BestFriend()               // false

chicken.SetBestFriend(friendlyChicken.ID())    // sets the reference
friendlyChickenRef, _ := chicken.BestFriend()  // the friendly chicken reference
isSet := friendlyChickenRef.IsSet()            // true
friendlyChicken := friendlyChickenRef.Get()    // the friendly chicken

chicken.BestFriend().Unset()                   // unsets friendlyChicken as chicken's best friend
isSet = friendlyChickenRef.IsSet()             // false
```
Fields with `anyOf` values also have additional setters, when they are not references:
```JSON
{
    "farm": {
        "owner": "string",
        "cutestResident": "anyOf<chicken,cow,pig>"
    }
}
```
```golang
farm := engine.Farm(id)
cutestResidentKind := farm.CutestResident().Kind()   // default is always the first type of the list ("Chicken")
farm.CutestResident().SetCow()
cutestResidentKind = farm.CutestResident().Kind()    // "Cow"
```
## adders
Adders are methods to add elements to fields with slice values. These are the different variants of slices that exist:
```JSON
{
    "person": {
        "ensurances": "[]ensurance",
        "friends": "[]*person",
        "nickNames": "[]string"
    }
}
```
The adders look like this:
```golang
person := engine.Person(id)
newEnsurance := person.AddEnsurance()  // returns the newly created ensurance

newFriend := engine.CreatePerson()
person.AddFriend(newFriend.ID())       // new friend added, no return

person.AddNickNames("peter", "pete")   // since nickNames is a slice of a basic type AddNickNames is a variadic function
```

## removers
Just like you can add to fields with slice values, you can also remove elements:
```JSON
{
    "person": {
        "ensurances": "[]ensurance",
        "friends": "[]*person",
        "nickNames": "[]string"
    }
}
```
Removers will always return the entity they are called on for convenient method chaining:
```golang
person := engine.Person(id)
_ = person.RemoveEnsurance(ensuranceID)   // returns person, just like all removers

person = person.RemoveFriend(personID)

person.RemoveNickNames("peter", "pete")
```

## meta fields
every entity comes with meta fields that you can access freely. Currently the only meta fields are `Path()` and `ID()`:
```JSON
{
  "address": {
    "streetName": "string",
    "houseNumber": "int"
  },
  "house": {
    "address": "address",
  },
}
```
```golang
address := engine.CreateHouse().Address()

houseID := house.ID()                   // 1
housePath := house.Path()               // "$.house.1"

addressID := address.ID()               // 2
addressPath := address.Path()           // "$.house.1.address"
```

# Defining the Config
## Restrictions and their Validation Error Messages
### structural:
| Error           | Text                                                             | Meaning                                                         |
| --------------- | ---------------------------------------------------------------- | --------------------------------------------------------------- |
| ErrIllegalValue | value assigned to key "{KeyName}" in "{ParentObject}" is invalid | An invalid value was defined (nil, "", List, Object in Object). |
<br/> 

### syntactical:
| Error                 | Text                                                                         | Meaning                                                                               |
| --------------------- | ---------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| ErrIllegalTypeName    | illegal type name "{KeyName}" in "{ParentObject}"                            | A type was named without adhering to syntax limitations (e.g. "fo$o", "func", "<-+"). |
| ErrInvalidValueString | value "{ValueString}" assigned to "{KeyName}" in "{ParantObject}" is invalid | An invalid value was assigned to a key                                                |
<br/> 

### logical:
| Error                 | Text                                                          | Meaning                                                              |
| --------------------- | ------------------------------------------------------------- | -------------------------------------------------------------------- |
| ErrTypeNotFound       | type with name "{TypeName}" in "{ParentObject}" was not found | A type was referenced as value but not defined anywhere in the data. |
| ErrRecursiveTypeUsage | illegal recursive type detected for "{RecurringKeyNames}"     | A recursive type was defined.                                        |
| ErrInvalidMapKey      | "{MapKey}" in "{ValueString}" is not a valid map key          | An uncomparable type was chosen as map key.                          |
| ErrUnknownMethod      | type "{TypeName}" has no method "{Literal}".                  | An unknown method was attempted to be referenced.                    |
<br/> 

### thematical:
Despite the fact that each of these errors would find a place in one of the above mentioned categories, they are listed separately from them since they are specific to the use case, and not related to the validation of actual go declarations at all.
| Error                        | Text                                                                                          | Meaning                                                                                                                          |
| ---------------------------- | --------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| ErrIncompatibleValue         | value "{ValueString}" assigned to "{KeyName}" in "{ParentObject}" is incompatible.            | The assigned value can't be used, as only golang's basic types, self defined types, and slices and pointers of them can be used. |
| ErrNonObjectType             | type "{TypeName}" is not an object type.                                                      | The defined type is not an object.                                                                                               |
| ErrIllegalCapitalization     | {type/field name} "{literal}" starts with a capital letter.                                   | A type or field name starts with a capital letter, which is not allowed.                                                         |
| ErrConflictingSingular       | "{KeyName1}" and "{KeyName2}" share the same singular form "{Singular}".                      | Due to the way state will be used two field names cannot have the same singular form.                                            |
| ErrUnavailableFieldName      | "{KeyName}" not an available name.                                                            | Due to internal usage of this FieldName it is unavailable.                                                                       |
| ErrDirectTypeUsage           | the type "{TypeName}" was used directly in "{ActionName}" instead of it's ID ("{TypeName}ID") | Only IDs of types are available in actions                                                                                       |
| ErrIllegalPointerParameter   | the parameter "{FieldName}" in "{ActionName}" contains a pointer value                        | Pointers can not be used as parameter as it would not make any sense                                                             |
| ErrTypeAndActionWithSameName | type and action "{Name}" have the same name                                                   | Types and Actions with the same name would cause conflicts in the generated code                                                 |
| ErrInvalidAnyOfDefinition    | "{valueString}" is not a valid `anyOf` definition                                             | anyOf definitions can not have single or duplicate types and must be in alphabetical order                                       |
| ErrResponeToUnknownAction    | there is no action defined for response "{ResponseName}"                                      | a response can only be defined with the same name as the action it belongs to                                                    |


# For Developers

## Getting started:
```
# install necessary dependencies
bash bootstrap.sh;

# generate necessary files (you may ignore output if it's not an error)
go generate;

# run tests
go test ./...
```

### test cases check list
- create element -> Create()
- delete element -> Delete()
- manipulate element -> SetName()
- add element -> AddItem()
- remove element -> RemoveItem()
- set reference -> SetBoundTo()
- unset reference -> BoundTo().Unset()
- add reference -> AddGuildMember()
- remove reference -> RemoveGuildMember()
- set anyOf reference -> SetTargetPlayer()
- unset anyOf reference -> Target().Unset()
- switch anyOf type in slice -> Interactable.SetItem()
- switch anyOf type -> SetOriginPosition()
- add anyOf reference -> AddTargetedByPlayer()
- remove anyOf reference -> RemoveTargetedByPlayer()
- creation of reference -> a new reference is created
- deletion of reference -> a reference is deleted
- replacement of reference -> an existing reference is replaced with a different one
- includes element if reference of reference got updated
- manages cyclical references
- manages self referencing elements

### TODO
- describe how operationkindDelete in tree
- nice tutorial for inspector
- document flags
- documentation

- build fails because of required modules from github. what do? (cant reproduce)
- find open port for integratpon test
- setters to return if new value == current value so no change is triggered
- custom handlers
- the generated code should prefix user defined names (or in some other way alter them to be unique) so they do not conflict with local variables
- data persistence
- release tree func (release slices, maps, and the pointers themselves)
- (this only appeared to be an issue because i didnt consider that the Setters create an entirely new Ref with anyContainer as child. So the ElementKind is always empty and the delete mthod of the child is therefore never triggered. For more clarity I added a deleteCurrentChild parameter to the function) SetTargetPlayer (a reference field) calls the `setPlayer` method, which removes the child element. CRITICAL ERROR!!!


more to come