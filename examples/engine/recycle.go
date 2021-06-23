package state

func (engine *Engine) recycleTree() {
	// for _, equipmentSet  := range engine.Tree.EquipmentSet {
	// }
	// for _, gearScore  := range engine.Tree.GearScore {
	// }
	// for _, item  := range engine.Tree.Item {
	// }
	for id, player := range engine.Tree.Player {
		engine.recyclePlayer(&player)
		// no Put since not a pointer
		delete(engine.Tree.Player, id)
	}
	// for _, position  := range engine.Tree.Position {
	// }
	// for _, zone  := range engine.Tree.Zone {
	// }
	// for _, zoneItem  := range engine.Tree.ZoneItem {
	// }
}

func (engine *Engine) recyclePlayer(player *Player) {
	// pools don't notice if the same Pointer is Put in multiple times
	// _, notYetRecycled := engine.assembleCache.player[player.ID]
	// if !notYetRecycled {return}

	// IF NOT NIL
	// FOR recycleEquipmentSet
	// clear EquipmentSets map
	// Put EquipmentSets map
	// set EquipmentSets == nil

	// IF NOT NIL
	// recycleGearScore
	// Put GearScore
	// set GearScore == nil

	// IF NOT NIL
	// FOR recyclePlayerReference
	// Put PlayerReference map
	// set GuildMembers == nil

	// IF NOT NIL
	// FOR recycleItem
	// Put Item map
	// set Items == nil

	// IF NOT NIL
	// recyclePosition
	// Put Position
	// set Position == nil

	// IF NOT NIL
	// recycleAnyOfPlayer_ZoneItemReference
	// Put AnyOfPlayer_ZoneItemReference
	// set AnyOfPlayer_ZoneItemReference == nil

	// IF NOT NIL
	// FOR recycleAnyOfPlayer_ZoneItemReference
	// Put AnyOfPlayer_ZoneItemReference map
	// set TargetedBy == nil

	// *player = Player{}

	// delete(engine.assembleCache.player, player.ID)
}

func (engine *Engine) recyclePlayerReference(playerReference *PlayerReference) {

	// IF NOT NIL
	// recyclePlayer
	// Put Player
	// set Player == nil

	// *playerReference = PlayerReference{}
}

func (engine *Engine) recycleInterface(any interface{}) {

	// IF NOT NIL
	// check every type
	// recycleElement
	// Put Element

}
