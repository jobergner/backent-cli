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
- a small script that regisers all _action files
- with a functions named after the action with parameters
- maintains an "action receiver" file where the new action gets registered in a switch
- a server with socket endpoint
- (sm.finish() dont know if really needed)
- create neat CLI with actions like 'register actions' (looks for file with action_ prefix), 'generate from config'

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

### TODO
- finish state machine
- finish state factory tests
- write state factory
- sketch orcestrator
- write orchestrator
- write action receiver generator
- create cli for 'generate from config' and 'register actions'
- write server
- write state conveyor
