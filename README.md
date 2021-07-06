## About

generates state machine

## Example API

```
// modifiers need to check patch if item is already modified
func (sm *StateMachine) {
   newItem := sm.
    NewItem(). // creates new Item and generates ID | mod
    SetValue(3, sm). // sets item value to 3 | mod
    SetWeight(1, sm) // sets item weight to 1 | mod

   player := sm.GetPlayer(playerID) // retrieves player
   player.
    AddItem(newItem.GetID, sm). // adds item id to player | mod
    SetWeight(player.GetWeight()+1, sm) // sets player weight | mod

   player.RemoveArmor(armorID, sm) // removes armor from player | mod
   sm.RemoveArmor(armorID) // removes armor from sm | mod

   player.SetDefense(player.GetDefense() - 1, sm) // mod

   sm.
    GetEnemy(enemyID).
    GetLocation().
    SetCoordX(6, sm).
    setCoordY(8, sm)

   sm.finish()
}
```

### data consistency:

certain changes within an action must be denieable due to being invalid.

Eg. ValueX is manipulated by User2 to be "2". However User1 has already manipulated ValueX to be "1" at an earlier point in time, but has a slower connection than User2.
If the connection of User1 is slow enough, the value would turn out to be "1", since User1's change would be applied later despite having happened earlier than the changes of User2.

However, this invalidaiton of changes must not be applied to all pieces of data which are affected by the change. Imagine a cenario where manipulating ValueY sets a negative (-) User positive (+).
Let's say, again, User1(-) manipulates ValueY to be "1", but has a slow connection, while User2(-) manipuates ValueY to be "2". If we have an action wide invalidation of all changes to any data,
ValueY would be 2, as User1's changes would be invalidated. However User1 would still be negative (1) as the changes to his own state would also be denied.

As result invalidation must happen action specific, or there has to be a type system in place where certain actions have certain overwriting rules.

### thread safety:

Only one action can be processed at a given time to ensure every piece of data is at the most recent state before manipulation or reading occurs.
This might not be the fastest way to process actions, but only this way conflicts can be ruled out.

### concurrent action processing:

dumb idea. pieces of data within one action may depend on each other.

### frames & patch batching:

some processing may be required to be done without an action triggering it. Such as physics when eg. an objetc is thrown/falling/colliding.
therefore the server state should be able to be broadcasted with every x frames.
A patch is complete when all actions in the queue and the frame itself is processed. then the patch will be applied to the state
and broadcasted.

### eventkinds:

a create event kind is not necessary as it is handled just like an update anyway.
it would only make things complicated.

### data races and frames:

processing running concurently will lead to data races. Thus, the action listener will feed actions into a queue,
and with each frame tick, all actions within the queue will be processed then the frame will be processed.

### (environment actor):

with each frame tick actions can be performed and a new patch for these actions can be created.
with patch batching this should be possible.

### generated orchestrator:

- a small script that regisers all \_action files
- with a functions named after the action with parameters
- maintains an "action receiver" file where the new action gets registered in a switch
- a server with socket endpoint
- (sm.finish() dont know if really needed)
  ￼
  1
- create neat CLI with actions like 'register actions' (looks for file with action\_ prefix), 'generate from config'

### validating input:

the original idea was to let the user configure the state inside a go file. However there are a few problems:

- the user would have to keep a file with type declarations somewhere outside or clutter their code base with types
- using go would allow/suggest too much freedom
- gives wrong idea about how the data works

better alternative would be to let the user define their state inside a yaml file with strict rules and validation.
make use of https://github.com/Java-Jonas/yamltostruct, maybe fork or copy/paste usable code.
additional restirctions:

- no named types
- only basic/self defined types and slices of them

should return simple-ast of the input.

### testing

- with decltostring
- some files will just be copy/pasted as they will always be the same

### state conveyor:

- go WASM clinet
- caches the previous revision of each element
- assembles the state in it's original tree-like structure
- 'sends' state to browser
- removes meta fields
- adds 'hasUpdated' field
- updates of elements with multiple children will only include children that actually updated
- upon generation a dist folder is created which includes wasm_exec.js, client.wasm, index.html and an your_code.js file

### local state

should be optional
Is just empty if not defined
only affected by 'local actions'
actions may be local and remote at the same time

### local-only-state and remote-state mapping:

some actions should/must not rely on the server accepting, processing, and broadcasting it.
(apart from UI) like a character walking, rotating, aiming. In this case the action should
be processed lokally, changes be applied to a local state, which only handles local data.
However a second action should still be send to the server, as the movement of a character
needs to be broadcasted to everyone.
The position of the the character should therefore be read from the local state.
What if the actions of another player manipulate the position of the character?
In this case the remote position state of the character would be changed.
it's up to the dev to decide which state is used, local or remote. maybe the dev should implement a flag
on the server side state to signal the frontend to use the remote state instead.

### deletion -> update in queue

what if a queued action including a delete is followed by a queue action with an update on the same element??
Setters, removers and adders will return receiver when operation kind is delete.

### Marshaling:

Fields of elements should not be accessible directly, only through the API.
A Wrapper for each element is needed which only exposes the API but not the fields themselves.

### Later Usage:

the package gets generated
is being updated with every 'register action'
User can decide initially whether he wants to use 'go install' to make it global.
use https://github.com/mailru/easyjson to generate json marshal func.

### Declared elements as receiver:

Having your basic type declarations in a package takes away the ability to to write methods for them.
The user would have to reassign and cast the type into a self defined type. but wont fix atm this is a price im willing to pay.

### Unintuitive API:

```
following issue:
...
player := sm.CreatePlayer(sm) // "attribute" default is 0
player.SetAttribute(1, sm) // sets "attribute" to 1
fmt.Println(player.GetAttribute()) // prints 0
```

this is very unintuitive as you'd expect the third line to print 1. However, player was returned as value, not a pointer, so it's not.
to work around this all methods ignore their receivers values, and rather get the element out of the statemachine via it's getter.
this would result in the following api

```
player := sm.CreatePlayer(sm)
player.SetAttribute(1, sm)
fmt.Println(player.GetAttribute(sm)) // prints 1, but requires the stateMachine to be passed
```

the `GetAttribute` method would look like this

```
func (_p Player) GetAttribute(sm *StateMachine) int {
   p := sm.GetPlayer(_p.player.ID)
   return p.Attribute
}
```

### Tree

its purpose is to assemble the data of the incoming patch in a tree.
the tree is assembled from the patch and fills in the missing parents of elements with the elements in the stateMachine's state.
the interesting thing about the tree is that it really only holds updated data and their parents,
and will omit children of elements that haven't updated (with the use of pointers)

### preventing the developer from mistakes

when returning slices in getter a new slice is created in order to
prevent the user from manipulating the slice directly and therefore disturbing
altering slices within the stateMachine's State or Patch

### making the server aware of user defined actions

- each "register actions" call re-generates the entire thing
- reads a actions config, which will be validated by validator
- creates action files based on defined actions (if file already exist log a warning)
- assumes statefunction import by reading go.mod file `"module <name>"` and appends it to parameters of action
- a new server.Start() method will be generated which expects the user defined actions as parameters

### conflicting variable names

when code is generated based on user input there is always a chance that the generated code does not compile, or even worse, silently contains errors. for example

```
<userDefinedName> := 3
foo := "hello"
```

If the `<userDefinedName>` happened to be `foo`, the code would not compile (obviously)
This risk can be reduced (and in some cases eliminated) by encrypting local variables if they are named exactly (!) after a user input

```
<userDefinedName>XYZ := 3
foo := "hello"
```

Even if `<userDefinedName>` is `foo`, the generated local variable would be named `fooXYZ` and therefore not cause an error.
It is still possible for errors and conflicts to occur, for example:

```
object.<userDefinedName1>UUID = "123"
object.<userDefinedName2>ID = "234"
```

if <userDefinedName1> is "foo" and <userDefinedName2> is "fooUU" the generated code would look like this

```
object.fooUUID = "123"
object.fooUUID = "234"
```

and `object.fooUUID` would be "234" instead of "123" so there is still a need for caution.
The reason I only encrypt the occurances of user defined input when it is exactly (e.g. NOT <userDefinedName>ID etc.) a user defined string within the code is because
i want the exported methods not to be encrypted (obviously), but also not their parameters (e.g. `func DoThis(user_a7h6v8ID UserID)` would be confusing for the user).
a workaround would be immediately reassign the parameters at the beginning of every exported method:

```
func DoThis(userID UserID) {
   user_a7h6v8ID := userID
   ...
}
```

Maybe a workaround for this could be wrapper methods that only wrap

```
func (i item) WrapDoThing(userID UserID) {
   i.actuallyDoThing(userID)
}
func (i item) actuallyDoThing(user_a8n39vkID UserID) {
   // code goes here
}
```

I will have to revisit this issue later.

### meta fields

structs will have meta fields like ID, OperationKind and HasParent. This would mean that the user can not define fields with the names "ID", "OperationKind" and "HasParent.
To give the user a tiny bit more freedom I will suffix these fields with "\_" ("OperationKind" -> "OperationKind\_"). I can't do this with ID however, because ID is also a getter method
and therefore unique. It will remain "ID".

### code generation: templates vs jennifer

| Templates                                 | Jennifer                                                                 |
| ----------------------------------------- | ------------------------------------------------------------------------ |
| writing code is easy and quick            | writing code is a bit tedious                                            |
| can generate anything                     | only able to generate go code                                            |
| logic statements with own syntax          | is go                                                                    |
| no formatting, type-/spellchecking        | some helper methods to reduce the risk of misspelling (jen.Range() etc.) |
| logic makes templates hardly readable     | easily deal with difficult logic statements                              |
| thrown errors are not very comprehensible | throws comprehensible errors based on go/format                          |

I believe templates are a great way of generating code, as long as there is not too much logic envolved.
As soon as a lot of confitionals and loops are used it feels like you are just writing code in a underdeveloped language.
When the output is not as expected you might as well guess where the issue is within your template, as there is no way of debugging it.
Writing code generation with jennifer was tedious but the written code is very maintainable compared to templates.

### actions convenience

the user should be able to pass entire objects as parameters instead of destructuring objects into many parameters in order to conveniently have all the data available on the server

### 'getting' non-existing element

'getting' a non existing element may lead to the creation of an element with the ID 0 when a 'setter' is used on it. To avoid this 'getters' will return an
element with OperationKindDelete when asked to return a non-existing element, so all manipulating operations know not to put it in the patch.

### reference objects

the user may want to just reference an object within an object e.g.

```
type player struct {
   target *player // targeted combat (e.g. WoW)
}
```

this should be possible. However some things need to be considered:

- for each element kind a wrapper is created (`player` => `playerRef`) which wraps all getters, setters, removers etc. but also adds the `IsSet`, `Unset` and `Set(Player)` method. (these meta methods need to be reserved in validator)
- references will be (as expected) indicated by a star (\*player) within the config, and represented by `ElementReference{elementID, parentID, parentKind}` in the data. this way no pointers will be used
- fields with references will need setter methods (with the entire object as parameter), an `IsSet` (false if ID == 0) and an `Unset` method
- when marshalling the tree into JSON recursiveness will not be a problem as the item will be represented as `{elementID, parentID, parentKind}` within the JSON
- deleting elements would now require to check the state for any references of that element and call `Unset` on it.
- adding/removing from a slice of references must not create/delete the element
- the getter method would not exist on the `{elementID, parentID, parentKind}` element, but on the parent element
- only self defined types can be used wiht references (not string, int etc.)
- for the client's convenience a get from reference method would be needed to get the in the tree referenced element out of the state
- can not be used on basic types
  What if references were already implemented:
- a reference is just an id to the element
- when a referenced element gets updated, the assembling step automatically includes it -> no need for recursive, expensive searching
- Updates and Deletes are already taken care of
- slices are also dealt with in assembling step
- only extra logic is:
  - setters/unsetters for non slice references
  - adders/removers should not create/delete the element
  - deleters need to search all elements for references of deleted element and set them to 0/remove them
  - (need to figure out when and how to stop recursiveness in tree (logic that is required anyway))
  - how to you know when a reference got deleted when there is no OperationKindDelete on it? you just dont
- maybe references must appear in tree as a reference. the element is then retrieved by a getter method which returns a fully rendered json
  Problems:
- no OperationKindDelete when the reference is deleted
- no OperationKindUpdate (slice is empty) when reference is added
- when a reference is added/set the parent element receives an OperationKindUpdate, but the field of the reference may still appear as nil
- when a reference is removed/unset there the parent element does receive an OperationKindUpdate, but the field of the reference will appear empty
- without true OperationKinds there can't be optional fields like references
  there needs to be some kind of object which maintains the operation state of the reference in order to achieve optional fields!
  right now there is no need for recursive tree climbing, as the tree assembling goes top down an inlcudes all everything until it hits a reference with OperationKindDelete anyway
  when an element which is mentioned in a reference gets deleted, immediately all references need to be looked at and deleted
  elements with non-slice references are empty (id == 0) when element is created because otherwise the functionality would be even more different from the slice references
  Problem: the parent object and field needs to be mentioned in order to remove the referenceID from a slice when referenced element gets deleted
  god damn iT!!!!!!
  When a reference is deleted, the field containing the reference ID is set to 0. but then there is no way for the assembler to include the reference with the ReferenceKindDelete
  there needs to be a way to delay setting the id=0 and straight away removing the reference id from a reference slice --> no, there needs to be a diff func
  ids arent straight away set to zero/reference ids removed from slice, but getters wont return them anymore (not returning references with OPKDelete) --> also no
  id=0/ref removing happens after tree assembling -->> nope
  -- need a diff function for ALL slices (even normal slices dont even work atm) and single references to use in assembling
  -- because then i can set id=0 right away/remove ids from slice
  getting a single reference should prob have 2 return values, (ref,ok). is !ok when ref is not set

### the any type

`anyOf<player,enemy>`

- in the data the value will be represented as `{ElementKind, ElementID}` called `anyOfEnemyPlayer` (alphabetical order). by creation of ast duplicates need to be taken care of (when the `anyOf<player,enemy>` is used in multiple fields)
- for each compilation of `anyOf<...>` a reference type will also be generated (e.g. `anyOfEnemyPlayerRef`). It will have the same setters methods as a reference type except for main setter `SetPlayer(Player)` and `SetEnemy(Enemy)`. However, unlike in usual references, the `anyOfEnemyPlayerRef` struct will not impletmend all all getters, setters, removers etc., but include the getter methods for each element kind
- by creation of the value the item mentioned first will be used (e.g. `player` when `anyOf<player,enemy>`), except if its a reference type
- the item used can be set by a `SetPlayer()`, or `SetEnemy()` method, which call create for respective element. The previously existing element would be deleted. If the any type is a pointer, the setters will require the entire object as parameter.
- getter methods are not callable on the parent element of the field, but on the `{elementKind, elementID}` element itself
- getter methods will always be callable, but if the value is set to `player` and the `Enemy` method is called the behaviour will be the same as when getting an object which does not exist
- when used as reference (`*anyOf<player,enemy>`) the element also gains the `IsSet` and `Unset` method.
- can not be used on basic types

---

- for each combination of elements a new element kind is created (`eg. anyOf<enemy,player>` -> `anyOfEnemyPlayer`)
- in the data it will be represented as `{ElementKind, ElementID}`
- will have `func (any anyOfEnemyPlayer) SetPlayer()` method. this way it can be decided whether it holds and `enemy` or `player` element (does not need to have knowledge of it's parent element)
- will have `func (any anyOfEnemyPlayer) Player()` as getters
- when a `player` element is being deleted, the deletion method will also have to check all `any` types which include `player` and delete them as well
- on deletion of an `any` type, the `any` type itself, and the contained element will be deleted
- the default value will be the first element kind mentioned
- when assembling a tree the `assembleAnyOfEnemyPlayer` method will return an interface
- references should be able to be used as expected with not additional complication
- since adders create an any container, removers need to remove it
- since I have a setDefaultValue bool in creators, i need an optional delete downstream bool in any-deleters
  when removing the any type as part of a reference, i dont want to delete the contained element
  when deleting an element (item) i want to delete the contained element
- deleting a reference if an any type should automatically remove the any type

### reduce complexity in references

right now its a bit of a mess. references definitely made the code base a lot more complex. There are too many different places where reference logic ist handled (updaters, setters, removers, adders, assembling) and the logic works very differently from the rest of the element handling. There will always be extra logic to references, as they are optional values.
changes i want to introduce:

- add OperationKinds to references -> handling becomes similar to other elements
- references do not always need to be communicated, only when changed
- introduce ids to references themselves. as a slice of references may contain the same referenced element twice, and if a reference was to appear without an id, it'd be ambigous which reference is meant.
- updating elements based on whether they contain a reference to an updated element should be it's own step within the updating cycle. atm it happens everytime an element is updated
  -> this concludes that references better be their own type in a state object, like every other element
- updating all elements referencing an updating element is recursive and very complex -> a better solution is required

### thoughts on the client

- golang webassembly is too slow because of syscall/js and too big
- tinygo does not have any useful serialization options and is also slowed down by syscall/js
- assemblyscript isnt worth it
- typescript might be the only viable option
  problem with using typescript is, that I'm not willing to write another AST and generator for it, especially since the relevant parts are already taken care of.
  currently my best option is to assemble the tree server side, make sure all necessary data is included, and use json-path for references within the tree.

### building a JSON path

- when an element is being referenced, the client needs to know where in the tree the element is to retrieve it (except when the reference was created, then the referenced element is included fully)
- since we know the trees structure at compile time, we can generate all possible paths along with it
- there are however dynamic elements to a path:
  - the id of the root element
  - on every field with slice values (index)
- when elements are added/removed, the indeces can change for all elements within the slice
- a path tracker can track all paths. it will persist through multiple UpdateStates
- when a child of an element is assembled, the assembler checks if the path of that child already exists. if not, or the path has updated, a new path is created and tracked

### Using Maps instead of slices internally

- paths can be evaluated upon creation of element as there are no shifting indeces (no need for path walking)
- paths can just be the json path string within each element (exported jsonPath field, internal pathSegments field)
- users can always retrieve data from the tree by using the path (which never changes)
- assembling can happen bottom-up (performance improvement for later)

### server responses

the client may need to be aware of the IDs the server generates, or the paths to elements which are created by the clients actions.
this is why actions need to return responses to the client who sent them.

### `Set<TypeName>` methods on references of any types

any types have setter methods to replace the current contained child element with a different one. Eg the type `anyOfPlayer_ZoneItem` owns the `.SetPlayer()` and `.SetZoneItem()` methods
which both delete the current child, and create a new one. This however causes problems when you are dealing with references of any types, as a setter method would delete the current child.
being able to call the setters on these references makes no sense. This is why any types have ref wrappers, which are created just before `Get()` is called on the reference. These wrappers exclude the
setter methods.

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
- nice tutorial for inspector
- document flags
- inspector to show the sent message as well
- show config in inspector
- port param for start method
- example flag for getStartedFactory
- should engine.Getters also be able to get non-root entities
- build fails because of required modules from github. what do?
- find open port for integratpon test
- setters to return if new value == current value so no change is triggered
- `path` as reserved field name in validator
- documentation
- find a way to create a UI to observe changes (one where you can define and send actions, get the updated state, the state change, and response displayed)
- custom handlers
- data persistence
- the generated code should prefix user defined names (or in some other way alter them to be unique) so they do not conflict with local variables
- release tree func (release slices, maps, and the pointers themselves)
- (this only appeared to be an issue because i didnt consider that the Setters create an entirely new Ref with anyContainer as child. So the ElementKind is always empty and the delete mthod of the child is therefore never triggered. For more clarity I added a deleteCurrentChild parameter to the function) SetTargetPlayer (a reference field) calls the `setPlayer` method, which removes the child element. CRITICAL ERROR!!!

<br>
<br>
<br>
<br>

# backent-cli
backent-cli generates a server which enables real-time state broadcasting of entities through websockets. Simply by using the generated API all changes to entities are braodcasted to the connected clients automatically. The API is being generated as a package based on a config where you can define the entity types.
```
func broadcastChanges(params state.ReceivedParams, engine *state.Engine) {

	player := engine.CreatePlayer()                  // creating the player

	player.SetName(params.Name)                      // setting the player name

	player.AddItem().SetName(params.FirstItemName)   // add an item and set item name

}
```

## Start Experimenting:
Explore backent-cli and its features with the Inspector and toy around until you feel comfortable. The Inspector is a graphical user interface for you to run locally and inspect your backent-cli generated servers, or in this case an example server that backent-cli will set up for you.
# NICE SCREENSHOT OF INSPECTOR
```
mkdir backent_example; cd backent_example;
backent-cli generate -example -out ./backent/;
go run .;
backent-cli inspect -port 3496;
```
# Overview
## Example Configuration
This is what a very simple `config.json` can look like. You may define as many types and actions as you like. You can read more about the possible configurations below!
```
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
```
backent-cli generate -config config.json -out backent/
```
## Run the Server:
```
package main

import "yourproject/state"

// how many processing frames per second the server will run
const fps = 30

func main() {
	err := state.Start(actions, sideEffects, fps)
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
```
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
```
{
  "address: {
    "streetName": "string",
    "houseNumber": "int
  },
  "house": {
    "address": "address"
  }
}
```
More about defining state types can be found here.
### actions:
Actions is how the server and client communicate. Here you can define which client data you want the server to be aware of in order to react with the assigned behaviour. Defining actions is very similar to defining the state except for some additional limitations. Only Go's basic types, and type references in the form of IDs can be used as values.
```
{
  "buildNewHouse": {
    "streetName": "string",
    "houseNumber": "int"
  }
}
```
The generator creates all necessary structures for you to simply tell the server how to react to a received action. The `params` struct will contain the defined data. Make use of your personal API to manipulate the `engine`'s state. The engine will keep track of all the changes you have made and tell the server to broadcast the changes to all connected clients automatically. 
```
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
```
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
```
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
```
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
```
var actions = state.Actions{
  // ...
	ChangeHouseNumber: func(params state.ChangeHouseNumberParams, engine *state.Engine) {
      house := engine.House(params.HouseID)
      house.Address().SetHouseNumber(params.NewHouseNumber)
	},
}
```
Triggering this action with a message would result in the following tree update:
```
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
```
{
  "address: {
    "streetName": "string",
    "houseNumber": "int
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
```
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
```
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
```
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
```
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
```
{
  "address: {
    "streetName": "string",
    "houseNumber": "int
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
```
house := engine.House(id)                  // house
residents := house.Residents()             // []person
address := house.Address()                 // address
streetName := house.Address().StreetName() // string
name := house.Residents()[0].Name()        // string
```
In the case of a referenced value:
```
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
retrieve the values like this:
```
menu := engine.Menu(id)                // menu
dealRef := menu.TodaysDeal()           // reference object of dish
isSet := dealRef.IsSet()               // bool
deal := dealRef.Get()                  // dish
// if you know that todaysDeal exists you may retrieve it directly
deal = menu.TodaysDeal().Get()         // dish
```
More on references and their methods can be read here.

In case of fields with `anyOf` types:
```
{
    "farm": {
        "owner": "string",
        "cutestResident": "anyOf<chicken,cow,pig>
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
```
farm := engine.Farm(id)                   // farm
cutestResident := farm.CutestResident()   // chicken|cow|pig
animalKind := cutestResident.Kind()       // "chicken"|"cow"|"pig"
cow := cutestResident.Cow()               // cow
// if you know of which kind the animal is you can retrieve it directly
cow := farm.CutestResident().Cow()        // cow
```
More on `anyOf` types and their methods can be read here.

If you try to retrieve a value where there is none, all manipulations applied to this entity will have no effect.
This can happen during the following curcumstances:
```
{
    "foo": {
        "bam": "*bar",
        "bal": "anyOf<bar,baz>"
    }
}
``` 
```
foo := engine.Foo(123)      // foo with id 123 may not exist
bam := foo.Bam()            // bam may not be set
bal := foo.Bal().Bar()      // bal may be of type baz and not bar
```