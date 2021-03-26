package state

func (se *Engine) assembleGearScore(gearScoreID GearScoreID) (tGearScore, bool) {
	gearScore, hasUpdated := se.Patch.GearScore[gearScoreID]
	if !hasUpdated {
		return tGearScore{}, false
	}

	var treeGearScore tGearScore

	treeGearScore.ID = gearScore.ID
	treeGearScore.OperationKind_ = gearScore.OperationKind_
	treeGearScore.Level = gearScore.Level
	treeGearScore.Score = gearScore.Score
	return treeGearScore, true
}

func (se *Engine) assemblePosition(positionID PositionID) (tPosition, bool) {
	position, hasUpdated := se.Patch.Position[positionID]
	if !hasUpdated {
		return tPosition{}, false
	}

	var treePosition tPosition

	treePosition.ID = position.ID
	treePosition.OperationKind_ = position.OperationKind_
	treePosition.X = position.X
	treePosition.Y = position.Y
	return treePosition, true
}

func (se *Engine) assembleItem(itemID ItemID) (tItem, bool) {
	item, hasUpdated := se.Patch.Item[itemID]
	if !hasUpdated {
		item = se.State.Item[itemID]
	}

	var treeItem tItem

	if treeGearScore, gearScoreHasUpdated := se.assembleGearScore(item.GearScore); gearScoreHasUpdated {
		hasUpdated = true
		treeItem.GearScore = &treeGearScore
	}

	treeItem.ID = item.ID
	treeItem.OperationKind_ = item.OperationKind_

	return treeItem, hasUpdated
}

func (se *Engine) assembleZoneItem(zoneItemID ZoneItemID) (tZoneItem, bool) {
	zoneItem, hasUpdated := se.Patch.ZoneItem[zoneItemID]
	if !hasUpdated {
		zoneItem = se.State.ZoneItem[zoneItemID]
	}

	var treeZoneItem tZoneItem

	if treeItem, itemHasUpdated := se.assembleItem(zoneItem.Item); itemHasUpdated {
		hasUpdated = true
		treeZoneItem.Item = &treeItem
	}
	if treePosition, positionHasUpdated := se.assemblePosition(zoneItem.Position); positionHasUpdated {
		hasUpdated = true
		treeZoneItem.Position = &treePosition
	}

	treeZoneItem.ID = zoneItem.ID
	treeZoneItem.OperationKind_ = zoneItem.OperationKind_
	return treeZoneItem, hasUpdated

}

func (se *Engine) assemblePlayer(playerID PlayerID) (tPlayer, bool) {
	player, hasUpdated := se.Patch.Player[playerID]
	if !hasUpdated {
		player = se.State.Player[playerID]
	}

	var treePlayer tPlayer

	if treeGearScore, gearScoreHasUpdated := se.assembleGearScore(player.GearScore); gearScoreHasUpdated {
		hasUpdated = true
		treePlayer.GearScore = &treeGearScore
	}
	for _, itemID := range deduplicateItemIDs(se.State.Player[player.ID].Items, se.Patch.Player[player.ID].Items) {
		if treeItem, itemHasUpdated := se.assembleItem(itemID); itemHasUpdated {
			hasUpdated = true
			treePlayer.Items = append(treePlayer.Items, treeItem)
		}
	}
	if treePosition, positionHasUpdated := se.assemblePosition(player.Position); positionHasUpdated {
		hasUpdated = true
		treePlayer.Position = &treePosition
	}

	treePlayer.ID = player.ID
	treePlayer.OperationKind_ = player.OperationKind_
	return treePlayer, hasUpdated
}

func (se *Engine) assembleZone(zoneID ZoneID) (tZone, bool) {
	zone, hasUpdated := se.Patch.Zone[zoneID]
	if !hasUpdated {
		zone = se.State.Zone[zoneID]
	}

	var treeZone tZone

	for _, zoneItemID := range deduplicateZoneItemIDs(se.State.Zone[zone.ID].Items, se.Patch.Zone[zone.ID].Items) {
		if treeZoneItem, zoneItemHasUpdated := se.assembleZoneItem(zoneItemID); zoneItemHasUpdated {
			hasUpdated = true
			treeZone.Items = append(treeZone.Items, treeZoneItem)
		}
	}
	for _, playerID := range deduplicatePlayerIDs(se.State.Zone[zone.ID].Players, se.Patch.Zone[zone.ID].Players) {
		if treePlayer, playerHasUpdated := se.assemblePlayer(playerID); playerHasUpdated {
			hasUpdated = true
			treeZone.Players = append(treeZone.Players, treePlayer)
		}
	}

	treeZone.ID = zone.ID
	treeZone.OperationKind_ = zone.OperationKind_
	treeZone.Tags = zone.Tags
	return treeZone, hasUpdated
}

func (se *Engine) assembleTree() Tree {
	tree := newTree()
	for _, gearScore := range se.Patch.GearScore {
		if !gearScore.HasParent_ {
			treeGearScore, hasUpdated := se.assembleGearScore(gearScore.ID)
			if hasUpdated {
				tree.GearScore[gearScore.ID] = treeGearScore
			}
		}
	}
	for _, item := range se.Patch.Item {
		if !item.HasParent_ {
			treeItem, hasUpdated := se.assembleItem(item.ID)
			if hasUpdated {
				tree.Item[item.ID] = treeItem
			}
		}
	}
	for _, player := range se.Patch.Player {
		if !player.HasParent_ {
			treePlayer, hasUpdated := se.assemblePlayer(player.ID)
			if hasUpdated {
				tree.Player[player.ID] = treePlayer
			}
		}
	}
	for _, position := range se.Patch.Position {
		if !position.HasParent_ {
			treePosition, hasUpdated := se.assemblePosition(position.ID)
			if hasUpdated {
				tree.Position[position.ID] = treePosition
			}
		}
	}
	for _, zone := range se.Patch.Zone {
		treeZone, hasUpdated := se.assembleZone(zone.ID)
		if hasUpdated {
			tree.Zone[zone.ID] = treeZone
		}
	}
	for _, zoneItem := range se.Patch.ZoneItem {
		if !zoneItem.HasParent_ {
			treeZoneItem, hasUpdated := se.assembleZoneItem(zoneItem.ID)
			if hasUpdated {
				tree.ZoneItem[zoneItem.ID] = treeZoneItem
			}
		}
	}

	for _, gearScore := range se.State.GearScore {
		if !gearScore.HasParent_ {
			if _, ok := tree.GearScore[gearScore.ID]; !ok {
				treeGearScore, hasUpdated := se.assembleGearScore(gearScore.ID)
				if hasUpdated {
					tree.GearScore[gearScore.ID] = treeGearScore
				}
			}
		}
	}
	for _, item := range se.State.Item {
		if !item.HasParent_ {
			if _, ok := tree.Item[item.ID]; !ok {
				treeItem, hasUpdated := se.assembleItem(item.ID)
				if hasUpdated {
					tree.Item[item.ID] = treeItem
				}
			}
		}
	}
	for _, player := range se.State.Player {
		if !player.HasParent_ {
			if _, ok := tree.Player[player.ID]; !ok {
				treePlayer, hasUpdated := se.assemblePlayer(player.ID)
				if hasUpdated {
					tree.Player[player.ID] = treePlayer
				}
			}
		}
	}
	for _, position := range se.State.Position {
		if !position.HasParent_ {
			if _, ok := tree.Position[position.ID]; !ok {
				treePosition, hasUpdated := se.assemblePosition(position.ID)
				if hasUpdated {
					tree.Position[position.ID] = treePosition
				}
			}
		}
	}
	for _, zone := range se.State.Zone {
		if _, ok := tree.Zone[zone.ID]; !ok {
			treeZone, hasUpdated := se.assembleZone(zone.ID)
			if hasUpdated {
				tree.Zone[zone.ID] = treeZone
			}
		}
	}
	for _, zoneItem := range se.State.ZoneItem {
		if !zoneItem.HasParent_ {
			if _, ok := tree.ZoneItem[zoneItem.ID]; !ok {
				treeZoneItem, hasUpdated := se.assembleZoneItem(zoneItem.ID)
				if hasUpdated {
					tree.ZoneItem[zoneItem.ID] = treeZoneItem
				}
			}
		}
	}

	return tree
}
