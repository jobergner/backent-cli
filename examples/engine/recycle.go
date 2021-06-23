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
	for id, zone := range engine.Tree.Zone {
		engine.recycleZone(&zone)
		delete(engine.Tree.Zone, id)
	}
	// for _, zoneItem  := range engine.Tree.ZoneItem {
	// }
}

func (engine *Engine) recycleEquipmentSetReference(equipmentSetReference *EquipmentSetReference) {
	*equipmentSetReference = EquipmentSetReference{}
}

func (engine *Engine) recycleGearScore(gearScore *GearScore) {
	*gearScore = GearScore{}
}

func (engine *Engine) recycleItem(item *Item) {

	if item.BoundTo != nil {
		engine.recyclePlayerReference(item.BoundTo)
		playerReferencePool.Put(item.BoundTo)
	}

	if item.GearScore != nil {
		engine.recycleGearScore(item.GearScore)
		gearScorePool.Put(item.GearScore)
	}

	*item = Item{}
}

func (engine *Engine) recyclePosition(position *Position) {
	*position = Position{}
}

func (engine *Engine) recycleZone(zone *Zone) {
	// if _, notYetRecycled := engine.assembleCache.zone[zone.ID]; !notYetRecycled {
	// 	return
	// }

	// IF NOT NIL
	// FOR recycleInterface
	// clear Interactables map
	// Put Interactables map

	// IF NOT NIL
	// FOR recycleZoneItem
	// clear Items map
	// Put Items map

	if zone.Players != nil {
		for _, player := range zone.Players {
			engine.recyclePlayer(&player)
		}
		for key := range zone.Players {
			delete(zone.Players, key)
		}
		zonePlayersMapPool.Put(zone.Players)
	}
	// IF NOT NIL
	// FOR recyclePlayer
	// clear Players map
	// Put Players map

	// IF NOT NIL
	// clear Tags slice
	// Put Tags slice

	*zone = Zone{}

	// delete(engine.assembleCache.zone, zone.ID)
}

func (engine *Engine) recycleAnyOfPlayer_ZoneItemReference(anyOfPlayer_ZoneItemReference *AnyOfPlayer_ZoneItemReference) {
	*anyOfPlayer_ZoneItemReference = AnyOfPlayer_ZoneItemReference{}
}

func (engine *Engine) recyclePlayerReference(playerReference *PlayerReference) {

	if playerReference.Player != nil {
		engine.recyclePlayer(playerReference.Player)
		playerPool.Put(playerReference.Player)
	}
	// IF NOT NIL
	// recyclePlayer
	// Put Player

	*playerReference = PlayerReference{}
}

func (engine *Engine) recyclePlayer(player *Player) {
	// pools don't notice if the same Pointer is Put in multiple times
	// _, notYetRecycled := engine.assembleCache.player[player.ID]
	// if !notYetRecycled {return}
	// if _, notYetRecycled := engine.assembleCache.player[player.ID]; !notYetRecycled {
	// 	return
	// }

	if player.EquipmentSets != nil {
		for _, equipmentSet := range player.EquipmentSets {
			engine.recycleEquipmentSetReference(&equipmentSet)
		}
		for key := range player.EquipmentSets {
			delete(player.EquipmentSets, key)
		}
		playerEquipmentSetsMapPool.Put(player.EquipmentSets)
	}
	// IF NOT NIL
	// FOR recycleEquipmentSet
	// clear EquipmentSets map
	// Put EquipmentSets map

	if player.GearScore != nil {
		engine.recycleGearScore(player.GearScore)
		gearScorePool.Put(player.GearScore)
	}
	// IF NOT NIL
	// recycleGearScore
	// Put GearScore

	if player.GuildMembers != nil {
		for _, playerReference := range player.GuildMembers {
			engine.recyclePlayerReference(&playerReference)
		}
		for key := range player.GuildMembers {
			delete(player.GuildMembers, key)
		}
		playerGuildMembersMapPool.Put(player.GuildMembers)
	}
	// IF NOT NIL
	// FOR recyclePlayerReference
	// clear GuildMembers map
	// Put PlayerReference map

	if player.Items != nil {
		for _, item := range player.Items {
			engine.recycleItem(&item)
		}
		for key := range player.Items {
			delete(player.Items, key)
		}
		playerItemsMapPool.Put(player.Items)
	}
	// IF NOT NIL
	// FOR recycleItem
	// clear Items map
	// Put Item map

	if player.Position != nil {
		engine.recyclePosition(player.Position)
		positionPool.Put(player.Position)
	}
	// IF NOT NIL
	// recyclePosition
	// Put Position

	if player.Target != nil {
		engine.recycleAnyOfPlayer_ZoneItemReference(player.Target)
		anyOfPlayer_ZoneItemReferencePool.Put(player.Target)
	}
	// IF NOT NIL
	// recycleAnyOfPlayer_ZoneItemReference
	// Put AnyOfPlayer_ZoneItemReference

	if player.TargetedBy != nil {
		for _, anyOfPlayer_ZoneItemReference := range player.TargetedBy {
			engine.recycleAnyOfPlayer_ZoneItemReference(&anyOfPlayer_ZoneItemReference)
		}
		for key := range player.TargetedBy {
			delete(player.TargetedBy, key)
		}
		playerTargetedByMapPool.Put(player.TargetedBy)
	}
	// IF NOT NIL
	// FOR recycleAnyOfPlayer_ZoneItemReference
	// clear AnyOfPlayer_ZoneItemReference map
	// Put AnyOfPlayer_ZoneItemReference map

	*player = Player{}
	// *player = Player{}

	// delete(engine.assembleCache.player, player.ID)
	// delete(engine.assembleCache.player, player.ID)
}

func (engine *Engine) recycleInterface(any interface{}) {

	// IF NOT NIL
	// check every type
	// recycleElement
	// Put Element

}
