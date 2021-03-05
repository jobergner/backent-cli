package statemachine

// this is kind of a mess atm
// the implementation completely missed the idea
// this code is supposed to only run in the wasm client
// its purpose is to assemble the data of the incoming patch in a tree
// so it would be called on the stateMachine in the client
// sm.Patch = incomingPatch
// dataTree := sm.assembleTree()
// sm.UpdateState()
// ...
// the tree is assembled from the patch and fills in the missing parents of elements
// with the elements it has in it's state.
// the implementation would look like this
// func (sm *StateMachine) assemble() Tree {
// the inteesting thing of the tree is that it really only holds updated data and their parents,
// and will omit children of elements that haven't updated (with the use of pointers)
//
// the implementation should also not be too hard to write.
// I just need to loop through each elementKind and of State AND Patch
// and check for len(parentage) == 0
// from there I build ALL children, BUT have 2 returned values:
// 1. the element
// 2. a boolean value whether the actually was an item taken from the Patch, and not only from the State
// (alternative is the item always gets returned as pointer, and is just nil in case)
// if the boolean is false, the built element get's dircarded

func (sm *StateMachine) assembleGearScore(gearScoreID GearScoreID) (_gearScore, bool) {
	gearScore, hasUpdated := sm.Patch.GearScore[gearScoreID]
	if !hasUpdated {
		return _gearScore{}, false
	}

	var treeGearScore _gearScore

	treeGearScore.ID = gearScore.ID
	treeGearScore.OperationKind = gearScore.OperationKind
	treeGearScore.Level = gearScore.Level
	treeGearScore.Score = gearScore.Score
	return treeGearScore, true
}

func (sm *StateMachine) assemblePosition(positionID PositionID) (_position, bool) {
	position, hasUpdated := sm.Patch.Position[positionID]
	if !hasUpdated {
		return _position{}, false
	}

	var treePosition _position

	treePosition.ID = position.ID
	treePosition.OperationKind = position.OperationKind
	treePosition.X = position.X
	treePosition.Y = position.Y
	return treePosition, true
}

func (sm *StateMachine) assembleItem(itemID ItemID) (_item, bool) {
	item, hasUpdated := sm.Patch.Item[itemID]
	if !hasUpdated {
		item = sm.State.Item[itemID]
	}

	var treeItem _item

	if treeGearScore, gearScoreHasUpdated := sm.assembleGearScore(item.GearScore); gearScoreHasUpdated {
		hasUpdated = true
		treeItem.GearScore = &treeGearScore
	}

	treeItem.ID = item.ID
	treeItem.OperationKind = item.OperationKind

	return treeItem, hasUpdated
}

func (sm *StateMachine) assembleZoneItem(zoneItemID ZoneItemID) (_zoneItem, bool) {
	zoneItem, hasUpdated := sm.Patch.ZoneItem[zoneItemID]
	if !hasUpdated {
		zoneItem = sm.State.ZoneItem[zoneItemID]
	}

	var treeZoneItem _zoneItem

	if treeItem, itemHasUpdated := sm.assembleItem(zoneItem.Item); itemHasUpdated {
		hasUpdated = true
		treeZoneItem.Item = &treeItem
	}
	if treePosition, positionHasUpdated := sm.assemblePosition(zoneItem.Position); positionHasUpdated {
		hasUpdated = true
		treeZoneItem.Position = &treePosition
	}

	treeZoneItem.ID = zoneItem.ID
	treeZoneItem.OperationKind = zoneItem.OperationKind
	return treeZoneItem, hasUpdated

}

func (sm *StateMachine) assemblePlayer(playerID PlayerID) (_player, bool) {
	player, hasUpdated := sm.Patch.Player[playerID]
	if !hasUpdated {
		player = sm.State.Player[playerID]
	}

	var treePlayer _player

	if treeGearScore, gearScoreHasUpdated := sm.assembleGearScore(player.GearScore); gearScoreHasUpdated {
		hasUpdated = true
		treePlayer.GearScore = &treeGearScore
	}
	for _, itemID := range deduplicateItemIDs(sm.State.Player[player.ID].Items, sm.Patch.Player[player.ID].Items) {
		if treeItem, itemHasUpdated := sm.assembleItem(itemID); itemHasUpdated {
			hasUpdated = true
			treePlayer.Items = append(treePlayer.Items, treeItem)
		}
	}
	if treePosition, positionHasUpdated := sm.assemblePosition(player.Position); positionHasUpdated {
		hasUpdated = true
		treePlayer.Position = &treePosition
	}

	treePlayer.ID = player.ID
	treePlayer.OperationKind = player.OperationKind
	return treePlayer, hasUpdated
}

func (sm *StateMachine) assembleZone(zoneID ZoneID) (_zone, bool) {
	zone, hasUpdated := sm.Patch.Zone[zoneID]
	if !hasUpdated {
		zone = sm.State.Zone[zoneID]
	}

	var treeZone _zone

	for _, zoneItemID := range deduplicateZoneItemIDs(sm.State.Zone[zone.ID].Items, sm.Patch.Zone[zone.ID].Items) {
		if treeZoneItem, zoneItemHasUpdated := sm.assembleZoneItem(zoneItemID); zoneItemHasUpdated {
			hasUpdated = true
			treeZone.Items = append(treeZone.Items, treeZoneItem)
		}
	}
	for _, playerID := range deduplicatePlayerIDs(sm.State.Zone[zone.ID].Players, sm.Patch.Zone[zone.ID].Players) {
		if treePlayer, playerHasUpdated := sm.assemblePlayer(playerID); playerHasUpdated {
			hasUpdated = true
			treeZone.Players = append(treeZone.Players, treePlayer)
		}
	}

	treeZone.ID = zone.ID
	treeZone.OperationKind = zone.OperationKind
	return treeZone, hasUpdated
}

func (sm *StateMachine) assembleTree() Tree {
	tree := newTree()
	for _, gearScore := range sm.Patch.GearScore {
		if len(gearScore.Parentage) == 0 {
			treeGearScore, hasUpdated := sm.assembleGearScore(gearScore.ID)
			if hasUpdated {
				tree.GearScore[gearScore.ID] = treeGearScore
			}
		}
	}
	for _, item := range sm.Patch.Item {
		if len(item.Parentage) == 0 {
			treeItem, hasUpdated := sm.assembleItem(item.ID)
			if hasUpdated {
				tree.Item[item.ID] = treeItem
			}
		}
	}
	for _, player := range sm.Patch.Player {
		if len(player.Parentage) == 0 {
			treePlayer, hasUpdated := sm.assemblePlayer(player.ID)
			if hasUpdated {
				tree.Player[player.ID] = treePlayer
			}
		}
	}
	for _, position := range sm.Patch.Position {
		if len(position.Parentage) == 0 {
			treePosition, hasUpdated := sm.assemblePosition(position.ID)
			if hasUpdated {
				tree.Position[position.ID] = treePosition
			}
		}
	}
	for _, zone := range sm.Patch.Zone {
		treeZone, hasUpdated := sm.assembleZone(zone.ID)
		if hasUpdated {
			tree.Zone[zone.ID] = treeZone
		}
	}
	for _, zoneItem := range sm.Patch.ZoneItem {
		if len(zoneItem.Parentage) == 0 {
			treeZoneItem, hasUpdated := sm.assembleZoneItem(zoneItem.ID)
			if hasUpdated {
				tree.ZoneItem[zoneItem.ID] = treeZoneItem
			}
		}
	}

	for _, gearScore := range sm.State.GearScore {
		if len(gearScore.Parentage) == 0 {
			if _, ok := tree.GearScore[gearScore.ID]; !ok {
				treeGearScore, hasUpdated := sm.assembleGearScore(gearScore.ID)
				if hasUpdated {
					tree.GearScore[gearScore.ID] = treeGearScore
				}
			}
		}
	}
	for _, item := range sm.State.Item {
		if len(item.Parentage) == 0 {
			if _, ok := tree.Item[item.ID]; !ok {
				treeItem, hasUpdated := sm.assembleItem(item.ID)
				if hasUpdated {
					tree.Item[item.ID] = treeItem
				}
			}
		}
	}
	for _, player := range sm.State.Player {
		if len(player.Parentage) == 0 {
			if _, ok := tree.Player[player.ID]; !ok {
				treePlayer, hasUpdated := sm.assemblePlayer(player.ID)
				if hasUpdated {
					tree.Player[player.ID] = treePlayer
				}
			}
		}
	}
	for _, position := range sm.State.Position {
		if len(position.Parentage) == 0 {
			if _, ok := tree.Position[position.ID]; !ok {
				treePosition, hasUpdated := sm.assemblePosition(position.ID)
				if hasUpdated {
					tree.Position[position.ID] = treePosition
				}
			}
		}
	}
	for _, zone := range sm.State.Zone {
		if _, ok := tree.Zone[zone.ID]; !ok {
			treeZone, hasUpdated := sm.assembleZone(zone.ID)
			if hasUpdated {
				tree.Zone[zone.ID] = treeZone
			}
		}
	}
	for _, zoneItem := range sm.State.ZoneItem {
		if len(zoneItem.Parentage) == 0 {
			if _, ok := tree.ZoneItem[zoneItem.ID]; !ok {
				treeZoneItem, hasUpdated := sm.assembleZoneItem(zoneItem.ID)
				if hasUpdated {
					tree.ZoneItem[zoneItem.ID] = treeZoneItem
				}
			}
		}
	}

	return tree
}
