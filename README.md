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